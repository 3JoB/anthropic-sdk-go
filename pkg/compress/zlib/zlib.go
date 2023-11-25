package zlib

import (
	"bytes"

	"github.com/3JoB/ulib/pool"
	"github.com/klauspost/compress/zlib"

	"github.com/3JoB/anthropic-sdk-go/v2/pkg/compress"
)

type Zlib struct{}

// Initialize a new Zlib object based on the Interface interface
func New() compress.Interface {
	return &Zlib{}
}

// Encode compresses the given bytes using ZLIB compression,
// returning the compressed data in a new bytes.Buffer.
func (z Zlib) Encode(v []byte) (*bytes.Buffer, error) {
	i := pool.NewBuffer()
	w := zlib.NewWriter(i)
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
func (z Zlib) Decode(v *bytes.Buffer) {
	r, _ := zlib.NewReader(v)
	d := compress.Reader(r)
	_ = r.Close()
	v.Reset()
	_, _ = v.Write(d)
}
