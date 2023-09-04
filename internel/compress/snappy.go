package compress

import (
	"bytes"

	"github.com/klauspost/compress/snappy"
)

type Snappy struct{}

func NewSnappy() Interface {
	return &Snappy{}
}

func (s *Snappy) Encode(v []byte) *bytes.Buffer {
	var i bytes.Buffer
	w := snappy.NewBufferedWriter(&i)
	w.Write(v)
	w.Close()
	return &i
}

func (s *Snappy) Decode(v *bytes.Buffer) []byte {
	return reader(snappy.NewReader(v))
}
