package compress

import (
	"bytes"

	"github.com/klauspost/compress/flate"
)

type Flate struct{}

// Initialize a new Deflate object based on the Interface interface
func NewFlate() Interface {
	return &Flate{}
}

// Encode compresses the given bytes using Deflate compression,
// returning the compressed data in a new bytes.Buffer.
func (f *Flate) Encode(v []byte) (*bytes.Buffer, error) {
	var i bytes.Buffer
	w, _ := flate.NewWriter(&i, 9)
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
func (f *Flate) Decode(v *bytes.Buffer) {
	r := flate.NewReader(v)
	d := reader(r)
	_ = r.Close()
	v.Reset()
	_, _ = v.Write(d)
}
