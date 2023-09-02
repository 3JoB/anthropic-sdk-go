// The pool subpackage is the new hashmap cache pool, which is only for testing now.
//
// It will replace the Context cache pool in the future.
package pool

import (
	"github.com/cornelk/hashmap"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
)

type Pool[v []data.MessageModule | string] struct {
	pool *hashmap.Map[string, v]
}

func NewPoolWithSlice() *Pool[[]data.MessageModule] {
	return &Pool[[]data.MessageModule]{
		pool: hashmap.New[string, []data.MessageModule](),
	}
}

func NewPoolWithCache()
