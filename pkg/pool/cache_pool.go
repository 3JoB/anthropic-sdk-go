package pool

import (
	"github.com/3JoB/ulib/litefmt"
	"github.com/cornelk/hashmap"

	"github.com/3JoB/anthropic-sdk-go/v2/pkg/compress"
)

type Pool struct {
	pool *hashmap.Map[string, string]
}

// Enable Compress (no use)
func (p *Pool) UseCompress(compress compress.Interface) error {
	return nil
}

// Get retrieves an element from the map under given hash key.
func (p *Pool) Get(k string) (string, bool) {
	return p.pool.Get(k)
}

// Set sets the value under the specified key to the map.
// An existing item for this key will be overwritten.
// If a resizing operation is happening concurrently while calling Set,
// the item might show up in the map after the resize operation is finished.
func (p *Pool) Set(k, v string) bool {
	p.pool.Set(k, v)
	return true
}

// Del deletes the key from the map and returns whether the key was deleted.
func (p *Pool) Del(k string) bool {
	return p.pool.Del(k)
}

// Insert sets the value under the specified key to the map if it does not exist yet.
// If a resizing operation is happening concurrently while calling Insert,
// the item might show up in the map after the resize operation is finished.
// Returns true if the item was inserted or false if it existed.
func (p *Pool) Insert(k, v string) bool {
	return p.pool.Insert(k, v)
}

// Flush will clear all data in the Pool.
func (p *Pool) ResetPool() {
	p.pool.Range(func(k, v string) bool {
		return p.pool.Del(k)
	})
}

// Append will take out the data,
// and then append a new piece of data to the end before saving it.
func (p *Pool) Append(k, v string) bool {
	r, ok := p.Get(k)
	if !ok {
		return p.Set(k, v)
	}
	return p.Set(k, litefmt.Sprint(r, v))
}

// Len returns the number of elements within the map.
func (p *Pool) Len() int {
	return p.pool.Len()
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
func (p *Pool) Range(f func(k string, v string) bool) {
	p.pool.Range(f)
}
