package signal

import (
	"testing"
)

func TestSig(t *testing.T) {
	Serve(func() { println(1) }, func() { println(2) })
}
