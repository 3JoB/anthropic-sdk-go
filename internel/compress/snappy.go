package compress

import (
	"github.com/klauspost/compress/snappy"
)

type Snappy struct{}

func NewSnappy() Interface {
	return &Snappy{}
}

func (s *Snappy) Encode(v []byte) []byte {
	return snappy.Encode(nil, v)
}

func (s *Snappy) Decode(v []byte) []byte {
	b, _ := snappy.Decode(nil, v)
	return b
}
