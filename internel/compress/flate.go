package compress

import (
	"bytes"

	"github.com/klauspost/compress/flate"
)

type Flate struct{}

func NewFlate() Interface {
	return &Flate{}
}

func (f *Flate) Encode(v []byte) *bytes.Buffer {
	var i bytes.Buffer
	w, _ := flate.NewWriter(&i, 9)
	w.Write(v)
	w.Close()
	return &i
}

func (f *Flate) Decode(v *bytes.Buffer) []byte {
	r := flate.NewReader(v)
	defer r.Close()
	return reader(r)
}
