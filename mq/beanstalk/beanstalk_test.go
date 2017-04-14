package beanstalk

import (
	"log"
	"reflect"
	"testing"
	"time"
)

var testToBytesCases = []struct {
	data  interface{}
	bytes []byte
}{
	{1, []byte("1")},
	{int8(-1), []byte("-1")},
	{int16(-1), []byte("-1")},
	{int32(-1), []byte("-1")},
	{int64(-1), []byte("-1")},
	{uint8(1), []byte("1")},
	{uint16(1), []byte("1")},
	{uint32(1), []byte("1")},
	{uint64(1), []byte("1")},
	{true, []byte("true")},
	{false, []byte("false")},
	{[]byte("bytes"), []byte("bytes")},
	{float32(1.234), []byte("1.234")},
	{1.2345678, []byte("1.2345678")},
	{[]int{1, 2, 3}, []byte("[1,2,3]")},
}

func TestBeanstalk(t *testing.T) {
	for i, cs := range testToBytesCases {
		bytes, err := toBytes(cs.data)
		if err != nil {
			t.Error(err)
		} else if !reflect.DeepEqual(bytes, cs.bytes) {
			t.Errorf("toBytes[%d] want %v, but have %v", i, cs.bytes, bytes)
		}
	}

	errorFataler := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	client := Dial("localhost:11300")

	b := time.Now()
	for i := 0; i < 1000; i++ {
		for _, cs := range testToBytesCases {
			errorFataler(client.Put("lor.mq.beanstalk.1", cs.data))
		}
	}
	log.Println(time.Since(b))

	b = time.Now()
	i := make(chan struct{}, 15000)
	client.ForEach(func(data []byte) error {
		i <- struct{}{}
		if len(i) >= 15000 {
			log.Println(time.Since(b))
		}
		return nil
	}, 10, "lor.mq.beanstalk.1")
}

func TestLog(t *testing.T) {
	client := Dial("localhost:11300")
	client.ForEach(func(data []byte) error {
		log.Println(string(data))
		return nil
	}, 1, "lor.log")
}
