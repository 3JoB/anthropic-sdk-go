package compress

import (
	"bytes"

	"github.com/klauspost/compress/zstd"
)

type ZST struct{}

func NewZST() Interface {
	return &ZST{}
}

func (zs *ZST) Encode(v []byte) *bytes.Buffer {
	var i bytes.Buffer
	w, _ := zstd.NewWriter(&i)
	w.Write(v)
	w.Close()
	return &i
}

func (zs *ZST) Decode(v *bytes.Buffer) []byte {
	r, _ := zstd.NewReader(v)
	defer r.Close()
	return reader(r)
}
