package log

import (
	"log"

	"github.com/martinz0/lor/etc"
	"github.com/martinz0/lor/mq/beanstalk"
)

var mq = beanstalk.Dial(etc.String("beanstalk.addr"))

func send(l *Log) {
	bytes, err := toBytes(l)
	reportError(err)
	if bytes != nil {
		reportError(mq.Put("lor.log", bytes))
	}
}

func reportError(err error) {
	if err != nil {
		log.Println(err)
	}
}
