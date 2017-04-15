package signal

import (
	"os"
	"os/signal"
	"syscall"
)

func Serve(fs ...func()) {
	signal.Ignore()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	<-c

	for _, f := range fs {
		f()
	}
}
