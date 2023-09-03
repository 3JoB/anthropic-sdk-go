package compress

import (
	"bytes"

	"github.com/klauspost/compress/flate"
)

type Flate struct{}

func NewFlate() Interface {
	return &Flate{}
}

func (f *Flate) Encode(v []byte) []byte {
	var i bytes.Buffer
	defer i.Reset()
	w, _ := flate.NewWriter(&i, 7)
	defer w.Close()
	w.Write(v)
	return i.Bytes()
}

func (f *Flate) Decode(v []byte) []byte {
	var i bytes.Buffer
	i.Write(v)
	defer i.Reset()
	r := flate.NewReader(&i)
	defer r.Close()
	return reader(r)
}
