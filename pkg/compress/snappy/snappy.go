package snappy

import (
	"bytes"

	"github.com/3JoB/anthropic-sdk-go/v2/pkg/compress"
	"github.com/3JoB/ulib/pool"
	"github.com/klauspost/compress/snappy"
)

type Snappy struct{}

// Initialize a new Snappy object based on the Interface interface
func NewSnappy() compress.Interface {
	return &Snappy{}
}

// Encode compresses the given bytes using Snappy compression,
// returning the compressed data in a new bytes.Buffer.
//
// Snappy only has the advantage of compression speed, and its
// compression ratio is ridiculously low.
// This compression engine is not recommended under normal circumstances
// and it causes additional overhead.
func (s Snappy) Encode(v []byte) (*bytes.Buffer, error) {
	i := pool.NewBuffer()
	w := snappy.NewBufferedWriter(i)
	if _, err := w.Write(v); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return i, nil
}

// The Decode method will first decode and then overwrite the data in the input *bytes.Buffer.
func (s Snappy) Decode(v *bytes.Buffer) {
	d := compress.Reader(snappy.NewReader(v))
	v.Reset()
	_, _ = v.Write(d)
}
