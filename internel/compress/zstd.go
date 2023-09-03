package compress

import (
	"bytes"

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
	var i bytes.Buffer
	defer i.Reset()
	i.Write(v)
	r, _ := zstd.NewReader(&i)
	defer r.Close()
	return reader(r)
}
