package brotli

import (
	"bytes"

	"github.com/3JoB/anthropic-sdk-go/v2/pkg/compress"
	"github.com/3JoB/ulib/pool"
	"github.com/andybalholm/brotli"
)

type Brotli struct{}

// Initialize a new Brotli object based on the Interface interface
func New() compress.Interface {
	return &Brotli{}
}

// Encode compresses the given bytes using brotli compression,
// returning the compressed data in a new bytes.Buffer.
func (b Brotli) Encode(v []byte) (*bytes.Buffer, error) {
	i := pool.NewBuffer()
	w := brotli.NewWriter(i)
	if _, err := w.Write(v); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return i, nil
}

// The Decode method will first decode and then
// overwrite the data in the input *bytes.Buffer.
func (b Brotli) Decode(v *bytes.Buffer) {
	d := compress.Reader(brotli.NewReader(v))
	v.Reset()
	_, _ = v.Write(d)
}
