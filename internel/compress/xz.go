package compress

import (
	"bytes"

	"github.com/ulikunitz/xz"
)

type XZ struct{}

// Initialize a new xz object based on the Interface interface
func NewXZ() Interface {
	return &XZ{}
}

// Encode compresses the given bytes using xz compression,
// returning the compressed data in a new bytes.Buffer.
func (x *XZ) Encode(v []byte) (*bytes.Buffer, error) {
	var i bytes.Buffer
	w, _ := xz.NewWriter(&i)
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
func (x *XZ) Decode(v *bytes.Buffer) {
	r, _ := xz.NewReader(v)
	d := reader(r)
	v.Reset()
	_, _ = v.Write(d)
}
