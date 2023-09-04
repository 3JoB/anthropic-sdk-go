package compress

import (
	"bytes"

	"github.com/klauspost/compress/zlib"
)

type Zlib struct{}

func NewZlib() Interface {
	return &Zlib{}
}

func (z *Zlib) Encode(v []byte) *bytes.Buffer {
	var i bytes.Buffer
	w := zlib.NewWriter(&i)
	w.Write(v)
	w.Close()
	return &i
}

func (z *Zlib) Decode(v *bytes.Buffer) []byte {
	r, _ := zlib.NewReader(v)
	defer r.Close()
	return reader(r)
}
