package snappy

import (
	"bytes"

	"github.com/3JoB/ulib/pool"
	"github.com/klauspost/compress/s2"

	"github.com/3JoB/anthropic-sdk-go/v2/pkg/compress"
)

type S2 struct{}

// Initialize a new S2 object based on the Interface interface
func NewS2() compress.Interface {
	return &S2{}
}

// Encode compresses the given bytes using S2 compression,
// returning the compressed data in a new bytes.Buffer.
func (s S2) Encode(v []byte) (*bytes.Buffer, error) {
	i := pool.NewBuffer()
	w := s2.NewWriter(i, s2.WriterBetterCompression())
	if _, err := w.Write(v); err != nil {
		return i, err
	}
	if err := w.Close(); err != nil {
		return i, err
	}
	return i, nil
}

// The Decode method will first decode and then overwrite the data in the input *bytes.Buffer.
func (s S2) Decode(v *bytes.Buffer) {
	d := compress.Reader(s2.NewReader(v))
	v.Reset()
	_, _ = v.Write(d)
}
