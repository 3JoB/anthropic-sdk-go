package compress

import (
	"bytes"

	"github.com/klauspost/pgzip"
)

type PGZip struct{}

func NewPGZip() Interface {
	return &PGZip{}
}

func (f *PGZip) Encode(v []byte) []byte {
	var i bytes.Buffer
	defer i.Reset()
	w := pgzip.NewWriter(&i)
	defer w.Close()
	w.Write(v)
	return i.Bytes()
}

func (f *PGZip) Decode(v []byte) []byte {
	var i bytes.Buffer
	i.Write(v)
	defer i.Reset()
	r, _ := pgzip.NewReader(&i)
	defer r.Close()
	return reader(r)
}
