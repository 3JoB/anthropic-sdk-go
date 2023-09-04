// The pool subpackage is the new hashmap cache pool, which is only for testing now.
//
// It will replace the Context cache pool in the future.
package pool

import (
	"github.com/cornelk/hashmap"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/internel/compress"
)

type Pool[T []data.MessageModule | string] struct {
	pool *hashmap.Map[string, T]
	c compress.Interface
	cmp bool // Compress status
	slice bool
}

// Create a new pool in slicing mode.
//
// Methods not available in this mode: UseCompress()
func NewPoolWithSlice() (*Pool[[]data.MessageModule]) {
	return &Pool[[]data.MessageModule]{
		pool: hashmap.New[string, []data.MessageModule](),
		slice: true,
	}
}

// Create a new pool in cached mode.
func NewPoolWithCache() *Pool[string] {
	return &Pool[string]{
		pool: &hashmap.Map[string, string]{},
	}
}

// Enable Compress
func (p *Pool[T]) UseComress(compress_model string) error {
	if p.slice {
		// Compression cannot be used in slice mode.
		return ErrSliceNoCmp
	}
	if p.c != nil {
		return ErrDisableSwitchCmp
	}
	switch compress_model {
	case "br":
		p.c = compress.NewBrotli()
	case "zs", "zst":
		p.c = compress.NewZST()
	case "gzip", "pgzip":
		p.c = compress.NewPGZip()
	case "deflate":
		p.c = compress.NewFlate()
	case "snappy":
		p.c = compress.NewSnappy()
	case "zlib":
		p.c = compress.NewZlib()
	default:
		return ErrUnavaCmpAlg
	}
	return nil
}