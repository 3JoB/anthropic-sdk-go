package anthropic_test

import (
	"testing"

	"github.com/3JoB/resty-ilo"
)

func BenchmarkResty(b *testing.B) {
	b.ResetTimer()
	client := resty.New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req, err := client.R().Get("https://example.com/")
			if err !=nil {
				panic(err)
			}
			req.RawBody().Close()
		}
	  })
}

/*1.1.2

goos: windows
goarch: amd64
pkg: github.com/3JoB/anthropic-sdk-go
cpu: Intel(R) Xeon(R) CPU E5-2670 0 @ 2.60GHz
BenchmarkResty-32    	       1	1437605100 ns/op	  363360 B/op	    2426 allocs/op
PASS
ok  	github.com/3JoB/anthropic-sdk-go	3.313s
*/

/* 1.1.1

goos: windows
goarch: amd64
pkg: github.com/3JoB/anthropic-sdk-go
cpu: Intel(R) Xeon(R) CPU E5-2670 0 @ 2.60GHz
BenchmarkResty-32    	       1	1377795500 ns/op	  354680 B/op	    2423 allocs/op
PASS
ok  	github.com/3JoB/anthropic-sdk-go	3.117s
*/