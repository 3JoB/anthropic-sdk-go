package compress

import (
	"bytes"

	"github.com/3JoB/brotli"
)

type Brotli struct{}

// Initialize a new Brotli object based on the Interface interface
func NewBrotli() Interface {
	return &Brotli{}
}

// Encode compresses the given bytes using brotli compression,
// returning the compressed data in a new bytes.Buffer.
func (b *Brotli) Encode(v []byte) (*bytes.Buffer, error) {
	var i bytes.Buffer
	w := brotli.NewWriter(&i)
	if _, err := w.Write(v); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return &i, nil
}

// The Decode method will first decode and then
// overwrite the data in the input *bytes.Buffer.
func (b *Brotli) Decode(v *bytes.Buffer) {
	r := brotli.NewReader(v)
	d := reader(r)
	_ = r.Close()
	v.Reset()
	_, _ = v.Write(d)
}
