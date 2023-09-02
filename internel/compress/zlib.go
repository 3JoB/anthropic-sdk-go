package compress

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/zlib"
)

type Zlib struct{}

func (z *Zlib) Encode(v []byte) []byte {
	var i bytes.Buffer
	defer i.Reset()
	w := zlib.NewWriter(&i)
	defer w.Close()
	w.Write(v)
	return i.Bytes()
}

func (z *Zlib) Decode(v []byte) []byte {
	b := bytes.NewReader(v)
	var i bytes.Buffer
	defer i.Reset()
	r, _ := zlib.NewReader(b)
	defer r.Close()
	io.Copy(&i, r)
	return i.Bytes()
}
