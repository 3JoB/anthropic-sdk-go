// The pool subpackage is the new hashmap cache pool, which is only for testing now.
//
// It will replace the Context cache pool in the future.
package hashpool

import (
	"github.com/cornelk/hashmap"
)

// Create a new pool in cached mode.
func NewPool() *Pool {
	return &Pool{
		pool: hashmap.New[string, string](),
	}
}
