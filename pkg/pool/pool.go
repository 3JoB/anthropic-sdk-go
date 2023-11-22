// The pool subpackage is the new hashmap cache pool, which is only for testing now.
//
// It will replace the Context cache pool in the future.
package pool

import (
	"github.com/cornelk/hashmap"

	"github.com/3JoB/anthropic-sdk-go/v2/internel/errors"
)

var (
	ErrUnavaCmpAlg      = errors.New("Unavailable compression algorithms.")
	ErrDisableSwitchCmp = errors.New("Disable switching of compression algorithm in a pool that has completed compression initialization.")
)

// Create a new pool in cached mode.
func NewPool() *Pool {
	return &Pool{
		pool: hashmap.New[string, string](),
	}
}
