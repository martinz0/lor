package mq

import "errors"

var ErrTemporary = errors.New("temporary error")

type MQ interface {
	Put(tube string, data interface{}) error
	Get(tube string, f func(data interface{}) error)
}
