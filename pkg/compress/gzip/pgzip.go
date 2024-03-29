package gzip

import (
	"bytes"

	"github.com/3JoB/ulib/pool"
	"github.com/klauspost/pgzip"

	"github.com/3JoB/anthropic-sdk-go/v2/pkg/compress"
)

type PGZip struct{}

// Initialize a new PGZip object based on the Interface interface
func New() compress.Interface {
	return &PGZip{}
}

// Encode compresses the given bytes using pgzip compression,
// returning the compressed data in a new bytes.Buffer.
//
// The reason why we choose pgzip instead of gzip is because
// it has special advantages when compressing large amounts of data.
// When the data block exceeds 1MB, pgzip will obtain a very considerable performance improvement.
func (f PGZip) Encode(v []byte) (*bytes.Buffer, error) {
	i := pool.NewBuffer()
	w := pgzip.NewWriter(i)
	if _, err := w.Write(v); err != nil {
		return i, err
	}
	if err := w.Close(); err != nil {
		return i, err
	}
	return i, nil
}

// The Decode method will first decode and then overwrite the data in the input *bytes.Buffer.
func (f PGZip) Decode(v *bytes.Buffer) {
	r, _ := pgzip.NewReader(v)
	d := compress.Reader(r)
	_ = r.Close()
	v.Reset()
	_, _ = v.Write(d)
}
