package pool

import (
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/cornelk/hashmap"
)

type slice_pool struct {
	pool *hashmap.Map[string, []data.MessageModule]
}

// Enable Compress
func (p *slice_pool) UseComress(compress_model string) error {
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
func (p *slice_pool) Set(k string, v []data.MessageModule) {
	p.pool.Set(k, v)
}

// Del deletes the key from the map and returns whether the key was deleted.
func (p *slice_pool) Del(k string) bool {
	return p.pool.Del(k)
}

// Len returns the number of elements within the map.
func (p *slice_pool) Len() int {
	return p.pool.Len()
}

// Range calls f sequentially for each key and value present in the map. 
// If f returns false, range stops the iteration.
func (p *slice_pool) Range(f func(string, []data.MessageModule) bool) {
	p.pool.Range(f)
}