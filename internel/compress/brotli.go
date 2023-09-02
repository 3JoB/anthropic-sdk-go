package compress

import (
	"bytes"
	"io"

	"github.com/3JoB/brotli"
)

type Brotli struct{}

func (b *Brotli) Encode(v []byte) []byte {
	var i bytes.Buffer
	defer i.Reset()
	w := brotli.NewWriter(&i)
	defer w.Close()
	w.Write(v)
	return i.Bytes()
}

func (b *Brotli) Decode(v []byte) []byte {
	n := bytes.NewReader(v)
	var i bytes.Buffer
	defer i.Reset()
	r := brotli.NewReader(n)
	defer r.Close()
	io.Copy(&i, r)
	return i.Bytes()
}
