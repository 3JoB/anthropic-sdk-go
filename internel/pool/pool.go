// The pool subpackage is the new hashmap cache pool, which is only for testing now.
//
// It will replace the Context cache pool in the future.
package pool

import (
	"github.com/cornelk/hashmap"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/internel/compress"
)

type Pool[v []data.MessageModule | string] struct {
	pool *hashmap.Map[string, v]
	c compress.Interface
	cmp bool // Compress status
}

func NewPoolWithSlice() *Pool[[]data.MessageModule] {
	return &Pool[[]data.MessageModule]{
		pool: hashmap.New[string, []data.MessageModule](),
	}
}

func NewPoolWithCache() *Pool[string] {
	return &Pool[string]{
		pool: &hashmap.Map[string, string]{},
	}
}
