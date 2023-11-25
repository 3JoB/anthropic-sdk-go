package xz

import (
	"bytes"

	"github.com/3JoB/ulib/pool"
	"github.com/ulikunitz/xz"

	"github.com/3JoB/anthropic-sdk-go/v2/pkg/compress"
)

type XZ struct{}

// Initialize a new xz object based on the Interface interface
func New() compress.Interface {
	return &XZ{}
}

// Encode compresses the given bytes using xz compression,
// returning the compressed data in a new bytes.Buffer.
func (x XZ) Encode(v []byte) (*bytes.Buffer, error) {
	i := pool.NewBuffer()
	w, _ := xz.NewWriter(i)
	if _, err := w.Write(v); err != nil {
		return i, err
	}
	if err := w.Close(); err != nil {
		return i, err
	}
	return i, nil
}

// The Decode method will first decode and then
// overwrite the data in the input *bytes.Buffer.
func (x XZ) Decode(v *bytes.Buffer) {
	r, _ := xz.NewReader(v)
	d := compress.Reader(r)
	v.Reset()
	_, _ = v.Write(d)
}
