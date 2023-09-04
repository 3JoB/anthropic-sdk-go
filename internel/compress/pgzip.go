package compress

import (
	"bytes"

	"github.com/klauspost/pgzip"
)

type PGZip struct{}

func NewPGZip() Interface {
	return &PGZip{}
}

func (f *PGZip) Encode(v []byte) *bytes.Buffer {
	var i bytes.Buffer
	w := pgzip.NewWriter(&i)
	w.Write(v)
	w.Close()
	return &i
}

func (f *PGZip) Decode(v *bytes.Buffer) []byte {
	r, _ := pgzip.NewReader(v)
	defer r.Close()
	return reader(r)
}
