package compress

import (
	"bytes"

	"github.com/klauspost/compress/zlib"
)

type Zlib struct{}

func NewZlib() Interface {
	return &Zlib{}
}

func (z *Zlib) Encode(v []byte) []byte {
	var i bytes.Buffer
	defer i.Reset()
	w := zlib.NewWriter(&i)
	defer w.Close()
	w.Write(v)
	return i.Bytes()
}

func (z *Zlib) Decode(v []byte) []byte {
	var i bytes.Buffer
	i.Write(v)
	defer i.Reset()
	r, _ := zlib.NewReader(&i)
	defer r.Close()
	return reader(r)
}
