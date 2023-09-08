package hashpool

import (
	"github.com/cornelk/hashmap"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
)

type slice_pool struct {
	pool *hashmap.Map[string, []data.MessageModule]
}

// Enable Compress
func (p *slice_pool) UseCompress(compress_model string) error {
	// Compression cannot be used in slice mode.
	return ErrSliceNoCmp
}

// Get retrieves an element from the map under given hash key.
func (p *slice_pool) Get(k string) ([]data.MessageModule, bool) {
	return p.pool.Get(k)
}

// Set sets the value under the specified key to the map.
// An existing item for this key will be overwritten.
// If a resizing operation is happening concurrently while calling Set,
// the item might show up in the map after the resize operation is finished.
func (p *slice_pool) Set(k string, v []data.MessageModule) bool {
	p.pool.Set(k, v)
	return true
}

// Del deletes the key from the map and returns whether the key was deleted.
func (p *slice_pool) Del(k string) bool {
	return p.pool.Del(k)
}

// Insert sets the value under the specified key to the map if it does not exist yet.
// If a resizing operation is happening concurrently while calling Insert,
// the item might show up in the map after the resize operation is finished.
// Returns true if the item was inserted or false if it existed.
func (p *slice_pool) Insert(k string, v []data.MessageModule) bool {
	return p.pool.Insert(k, v)
}

// Flush will clear all data in the Pool.
func (p *slice_pool) ResetPool() {
	p.pool.Range(func(k string, v []data.MessageModule) bool {
		return p.pool.Del(k)
	})
}

// Append will take out the data,
// and then append a new piece of data to the end before saving it.
func (p *slice_pool) Append(k string, v []data.MessageModule) {}

// Len returns the number of elements within the map.
func (p *slice_pool) Len() int {
	return p.pool.Len()
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
func (p *slice_pool) Range(f func(string, []data.MessageModule) bool) {
	p.pool.Range(f)
}
