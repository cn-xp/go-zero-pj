package go_zero_pj

import (
	"net/http"
	"testing"
)

func Benchmark_test(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := http.Get("http://localhost:8888/expand?shorten=f35b2a")
		if err != nil {
			panic(err)
		}
	}

}
