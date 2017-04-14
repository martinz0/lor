package etc

import (
	"log"
	"testing"
)

func TestEtc(t *testing.T) {
	ctl := make(chan struct{}, 5)
	for {
		ctl <- struct{}{}
		log.Println(Bool("system.dev"))
		<-ctl
	}
}

func BenchmarkGet(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Bool("system.dev")
	}
}
