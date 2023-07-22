package compress

import (
	"bytes"
	"io"

	"github.com/andybalholm/brotli"
)

type Brotli struct{}

func (b *Brotli) Encode(v []byte) []byte {
	var i bytes.Buffer
	w := brotli.NewWriter(&i)
	w.Write(v)
	w.Close()
	return i.Bytes()
}

func (b *Brotli) Decode(v []byte) []byte {
	n := bytes.NewReader(v)
	var o bytes.Buffer
	r := brotli.NewReader(n)
	io.Copy(&o, r)
	return o.Bytes()
}
