package compress

import (
	"bytes"
	"compress/zlib"
	"io"
)

type Zlib struct{}

func (z *Zlib) Encode(v []byte) []byte {
	var i bytes.Buffer
	w := zlib.NewWriter(&i)
	w.Write(v)
	w.Close()
	return i.Bytes()
}

func (z *Zlib) Decode(v []byte) []byte {
	b := bytes.NewReader(v)
	var o bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&o, r)
	return o.Bytes()
}
