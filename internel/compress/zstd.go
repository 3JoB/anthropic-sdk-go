package compress

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/zstd"
)

type ZSTD struct{}

func NewZSTD() Interface {
	return &ZSTD{}
}

func (zs *ZSTD) Encode(v []byte) []byte {
	var i bytes.Buffer
	defer i.Reset()
	w, _ := zstd.NewWriter(&i)
	defer w.Close()
	w.Write(v)
	return i.Bytes()
}

func (zs *ZSTD) Decode(v []byte) []byte {
	b := bytes.NewReader(v)
	var i bytes.Buffer
	defer i.Reset()
	r, _ := zstd.NewReader(b)
	defer r.Close()
	io.Copy(&i, r)
	return i.Bytes()
}
