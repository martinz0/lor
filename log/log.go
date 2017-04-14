package log

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/martinz0/lor/etc"
)

const (
	DEBUG = "debug"
	INFO  = "info"
	WARN  = "warn"
	ERROR = "error"
)

type Log struct {
	Time  time.Time
	Level string
	Msg   []string
}

func (l *Log) String() string {
	var b bytes.Buffer
	b.WriteString("level=")
	b.WriteString(l.Level)
	b.WriteString("\t\b\b\b\b\b")
	b.WriteString("msg=")
	b.WriteString(strings.Join(l.Msg, ","))
	return b.String()
}

var logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

func l(level string, v ...interface{}) *Log {
	log := &Log{
		Time:  time.Now(),
		Level: level,
		Msg:   make([]string, 0, len(v)),
	}
	for _, m := range v {
		bytes, _ := toBytes(m)
		log.Msg = append(log.Msg, string(bytes))
	}
	logger.Println(log)
	return log
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

func Debug(v ...interface{}) {
	if etc.Dev() {
		l(DEBUG, v...)
	}
}

func Info(v ...interface{}) {
	if etc.Dev() {
		l(INFO, v...)
	}
}

func Warn(v ...interface{}) {
	send(l(WARN, v...))
}

func Error(v ...interface{}) {
	send(l(ERROR, v...))
}
