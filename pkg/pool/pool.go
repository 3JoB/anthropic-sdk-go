// The pool subpackage is the new hashmap cache pool, which is only for testing now.
//
// It will replace the Context cache pool in the future.
package pool

import (
	"errors"

	"github.com/cornelk/hashmap"
)

var (
	ErrUnavaCmpAlg      = errors.New("unavailable compression algorithms")
	ErrDisableSwitchCmp = errors.New("disable switching of compression algorithm in a pool that has completed compression initialization")
)

// Create a new pool in cached mode.
func NewPool() *Pool {
	return &Pool{
		pool: hashmap.New[string, string](),
	}
}
