package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	Debug(1, 2, 3)
	Info(1, 2, 3)
	Warn(1, 2, 3)
	Error(1, 2, 3, Log{})
}
