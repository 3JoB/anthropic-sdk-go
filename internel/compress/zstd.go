package compress

import (
	"bytes"

	"github.com/klauspost/compress/zstd"
)

type ZSTD struct{}

func NewZSTD() Interface {
	return &ZSTD{}
}

func (zs *ZSTD) Encode(v []byte) *bytes.Buffer {
	var i bytes.Buffer
	w, _ := zstd.NewWriter(&i)
	w.Write(v)
	w.Close()
	return &i
}

func (zs *ZSTD) Decode(v *bytes.Buffer) []byte {
	r, _ := zstd.NewReader(v)
	defer r.Close()
	return reader(r)
}
