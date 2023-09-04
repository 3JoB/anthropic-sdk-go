// The pool subpackage is the new hashmap cache pool, which is only for testing now.
//
// It will replace the Context cache pool in the future.
package pool

import (
	"github.com/cornelk/hashmap"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
)

type Interface[T []data.MessageModule | string] interface {
	UseComress(compress_model string) error
	Get(string) (T, bool)
	Set(string, T)
	Del(string) bool
	Len() int
	Range(f func(string, T) bool)
}

// Create a new pool in slicing mode.
//
// Methods not available in this mode: UseCompress()
func NewPoolWithSlice() Interface[[]data.MessageModule] {
	return &slice_pool{
		pool: hashmap.New[string, []data.MessageModule](),
	}
}

// Create a new pool in cached mode.
func NewPoolWithCache() Interface[string] {
	return &cache_pool{
		pool: hashmap.New[string, string](),
	}
}
