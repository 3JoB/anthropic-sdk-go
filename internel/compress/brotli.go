package compress

import (
	"bytes"

	"github.com/3JoB/brotli"
)

type Brotli struct{}

func NewBrotli() Interface {
	return &Brotli{}
}

func (b *Brotli) Encode(v []byte) []byte {
	var i bytes.Buffer
	defer i.Reset()
	w := brotli.NewWriter(&i)
	defer w.Close()
	w.Write(v)
	return i.Bytes()
}

func (b *Brotli) Decode(v []byte) []byte {
	var i bytes.Buffer
	i.Write(v)
	defer i.Reset()
	r := brotli.NewReader(&i)
	defer r.Close()
	return reader(r)
}
