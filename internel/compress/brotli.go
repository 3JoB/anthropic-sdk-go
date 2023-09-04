package compress

import (
	"bytes"

	"github.com/3JoB/brotli"
)

type Brotli struct{}

func NewBrotli() Interface {
	return &Brotli{}
}

func (b *Brotli) Encode(v []byte) *bytes.Buffer {
	// Done
	var i bytes.Buffer
	w := brotli.NewWriter(&i)
	w.Write(v)
	w.Close()
	return &i
}

func (b *Brotli) Decode(v *bytes.Buffer) []byte {
	r := brotli.NewReader(v)
	defer r.Close()
	return reader(r)
}
