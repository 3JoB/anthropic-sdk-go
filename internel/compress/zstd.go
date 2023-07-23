package compress

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/zstd"
)

type ZSTD struct{}

func (zs *ZSTD) Encode(v []byte) []byte {
	var i bytes.Buffer
	w, _ := zstd.NewWriter(&i)
	w.Write(v)
	w.Close()
	return i.Bytes()
}

func (zs *ZSTD) Decode(v []byte) []byte {
	b := bytes.NewReader(v)
	var o bytes.Buffer
	r, _ := zstd.NewReader(b)
	io.Copy(&o, r)
	return o.Bytes()
}
