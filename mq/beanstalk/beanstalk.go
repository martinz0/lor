package beanstalk

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/kr/beanstalk"

	"lor/mq"
)

type Beanstalk struct {
	*beanstalk.Conn
}

func Dial(addr string) (*Beanstalk, error) {
	conn, err := beanstalk.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Beanstalk{conn}, nil
}

func (b *Beanstalk) Put(tube string, data interface{}) error {
	t := &beanstalk.Tube{
		Conn: b.Conn,
		Name: tube,
	}
	bytes, err := toBytes(data)
	if err != nil {
		return err
	}
	_, err = t.Put(bytes, 0, 0, 10*time.Second)
	return err
}

func toBytes(data interface{}) (bytes []byte, err error) {
	switch d := data.(type) {
	case string:
		bytes = []byte(d)
	case []byte:
		bytes = d
	case bool:
		bytes = strconv.AppendBool(bytes, d)
	case int:
		bytes = strconv.AppendInt(bytes, int64(d), 10)
	case int8:
		bytes = strconv.AppendInt(bytes, int64(d), 10)
	case int16:
		bytes = strconv.AppendInt(bytes, int64(d), 10)
	case int32:
		bytes = strconv.AppendInt(bytes, int64(d), 10)
	case int64:
		bytes = strconv.AppendInt(bytes, int64(d), 10)
	case uint:
		bytes = strconv.AppendUint(bytes, uint64(d), 10)
	case uint8:
		bytes = strconv.AppendUint(bytes, uint64(d), 10)
	case uint16:
		bytes = strconv.AppendUint(bytes, uint64(d), 10)
	case uint32:
		bytes = strconv.AppendUint(bytes, uint64(d), 10)
	case uint64:
		bytes = strconv.AppendUint(bytes, uint64(d), 10)
	case float32:
		bytes = strconv.AppendFloat(bytes, float64(d), 'f', -1, 32)
	case float64:
		bytes = strconv.AppendFloat(bytes, d, 'f', -1, 64)
	default:
		bytes, err = json.Marshal(data)
	}
	return
}

func (b *Beanstalk) Get(f func(data []byte) error, tube ...string) error {
	tubeSet := beanstalk.NewTubeSet(b.Conn, tube...)
	return b.get(f, tubeSet)
}

func (b *Beanstalk) ForEach(f func(data []byte) error, tube ...string) {
	tubeSet := beanstalk.NewTubeSet(b.Conn, tube...)
	for {
		errorReporter(b.get(f, tubeSet))
	}
}

func (b *Beanstalk) get(f func(data []byte) error, tubeSet *beanstalk.TubeSet) error {
	jobId, body, err := tubeSet.Reserve(time.Minute)
	if err != nil {
		return err
	}
	if err = f(body); err != nil {
		if err == mq.ErrTemporary {
			return b.Release(jobId, 0, time.Second)
		}
		return err
	}
	return b.Delete(jobId)
}

func errorReporter(err error) {
	if err != nil {
		log.Println(err)
	}
}
