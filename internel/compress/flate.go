package compress

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/flate"
)

type Flate struct{}

func (f *Flate) Encode(v []byte) []byte {
	var i bytes.Buffer
	w, _ := flate.NewWriter(&i, 7)
	w.Write(v)
	w.Close()
	return i.Bytes()
}

func (f *Flate) Decode(v []byte) []byte {
	b := bytes.NewReader(v)
	var o bytes.Buffer
	r := flate.NewReader(b)
	io.Copy(&o, r)
	return o.Bytes()
}
