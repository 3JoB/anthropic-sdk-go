package hashpool

import (
	"bytes"

	"github.com/3JoB/unsafeConvert"
	"github.com/cornelk/hashmap"

	"github.com/3JoB/anthropic-sdk-go/v2/internel/compress"
)

type cache_pool struct {
	pool *hashmap.Map[string, string]
	c    compress.Interface
	cmp  bool // Compress status
}

// Enable Compress
func (p *cache_pool) UseComress(compress_model string) error {
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

// Get retrieves an element from the map under given hash key.
func (p *cache_pool) Get(k string) (string, bool) {
	d, ok := p.pool.Get(k)
	if p.cmp {
		if !ok {
			return d, ok
		}
		var b *bytes.Buffer
		b.WriteString(d)
		p.c.Decode(b)
		if b.Len() == 0 {
			return "", ok
		}
		b.Reset()
		return unsafeConvert.StringSlice(b.Bytes()), ok
	}
	return d, ok
}

// Set sets the value under the specified key to the map.
// An existing item for this key will be overwritten.
// If a resizing operation is happening concurrently while calling Set,
// the item might show up in the map after the resize operation is finished.
func (p *cache_pool) Set(k string, v string) bool {
	if p.cmp {
		buf, err := p.c.Encode(unsafeConvert.ByteSlice(v))
		if err != nil {
			return false
		}
		v = unsafeConvert.StringSlice(buf.Bytes())
		defer buf.Reset()
	}
	p.pool.Set(k, v)
	return true
}

// Del deletes the key from the map and returns whether the key was deleted.
func (p *cache_pool) Del(k string) bool {
	return p.pool.Del(k)
}

// Insert sets the value under the specified key to the map if it does not exist yet.
// If a resizing operation is happening concurrently while calling Insert,
// the item might show up in the map after the resize operation is finished.
// Returns true if the item was inserted or false if it existed.
func (p *cache_pool) Insert(k, v string) bool {
	return p.pool.Insert(k, v)
}

// Flush will clear all data in the Pool.
func (p *cache_pool) ResetPool() {
	p.pool.Range(func(k, v string) bool {
		return p.pool.Del(k)
	})
}

// Append will take out the data,
// and then append a new piece of data to the end before saving it.
func (p *cache_pool) Append(k, v string) {}

// Len returns the number of elements within the map.
func (p *cache_pool) Len() int {
	return p.pool.Len()
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
func (p *cache_pool) Range(f func(k string, v string) bool) {
	p.pool.Range(f)
}
