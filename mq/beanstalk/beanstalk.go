package beanstalk

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/kr/beanstalk"
	"github.com/martinz0/lor/mq"
)

type Conn struct {
	*beanstalk.Conn
	addr string
}

func (c *Conn) Close() error {
	println("close " + c.addr)
	return c.Conn.Close()
}

type Beanstalk struct {
	mu    sync.Mutex
	conns map[string][]*Conn

	addrs []string
	rr    int
}

func Dial(addrs ...string) *Beanstalk {
	conns := make(map[string][]*Conn)
	for _, addr := range addrs {
		conns[addr] = make([]*Conn, 0)
	}
	return &Beanstalk{
		conns: conns,
		addrs: addrs,
	}
}

func (b *Beanstalk) Do(f func(*Conn) error) error {
	conn, err := b.borrowConn(b.addrs[b.rr])
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			e, ok := err.(beanstalk.ConnError)
			if ok {
				if e.Err == io.EOF {
					conn.Close()
					return
				}
				ee, ok := e.Err.(net.Error)
				if ok && !ee.Temporary() {
					conn.Close()
					return
				}
			}
		}
		b.returnConn(conn)
	}()
	err = f(conn)
	return err
}

func (b *Beanstalk) borrowConn(addr string) (*Conn, error) {
	conn := b.fromConns(addr)
	if conn != nil {
		return conn, nil
	}
	return b.conn(addr)
}

func (b *Beanstalk) conn(addr string) (*Conn, error) {
	conn, err := beanstalk.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Conn{
		Conn: conn,
		addr: addr,
	}, nil
}

func (b *Beanstalk) fromConns(addr string) *Conn {
	b.mu.Lock()
	defer b.mu.Unlock()
	conns, _ := b.conns[addr]
	if len(conns) == 0 {
		return nil
	}
	conn := conns[0]
	copy(conns, conns[1:])
	b.conns[addr] = conns[:len(conns)-1]
	return conn
}

func (b *Beanstalk) returnConn(conn *Conn) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.conns[conn.addr] = append(b.conns[conn.addr], conn)
}

func (b *Beanstalk) Put(tube string, data interface{}) error {
	return b.Do(func(c *Conn) error {
		t := &beanstalk.Tube{
			Conn: c.Conn,
			Name: tube,
		}
		bytes, err := toBytes(data)
		if err != nil {
			return err
		}
		_, err = t.Put(bytes, 0, 0, 10*time.Second)
		return err
	})
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
	return b.Do(func(c *Conn) error {
		tubeSet := beanstalk.NewTubeSet(c.Conn, tube...)
		return c.get(f, tubeSet)
	})
}

func (b *Beanstalk) ForEach(f func(data []byte) error, num int, tube ...string) {
	p := func() {
		for {
			log.Println(b.Do(func(c *Conn) error {
				tubeSet := beanstalk.NewTubeSet(c.Conn, tube...)
				for {
					if err := c.get(f, tubeSet); err != nil {
						return err
					}
				}
			}))
			time.Sleep(1e9)
		}
	}
	for i := 0; i < num-1; i++ {
		go p()
	}
	p()
}

func (c *Conn) get(f func(data []byte) error, tubeSet *beanstalk.TubeSet) error {
	jobId, body, err := tubeSet.Reserve(time.Minute)
	if err != nil {
		return err
	}
	if err = f(body); err != nil {
		if err == mq.ErrTemporary {
			return c.Release(jobId, 0, time.Second)
		}
		return err
	}
	return c.Delete(jobId)
}

func errorReporter(err error) {
	if err != nil {
		log.Println(err)
	}
}
