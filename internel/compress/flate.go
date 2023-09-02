package compress

import (
	"bytes"
	"io"

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
	b := bytes.NewReader(v)
	var o bytes.Buffer
	defer o.Reset()
	r := flate.NewReader(b)
	defer r.Close()
	io.Copy(&o, r)
	return o.Bytes()
}
