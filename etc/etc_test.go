package etc

import (
	"log"
	"testing"
)

func TestEtc(t *testing.T) {
	log.Println(Bool("system.dev"))
}

func BenchmarkGet(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Bool("system.dev")
	}
}
