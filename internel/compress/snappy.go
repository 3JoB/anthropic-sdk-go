package compress

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/snappy"
)

type Snappy struct{}

func (s *Snappy) Encode(v []byte) []byte {
	var i bytes.Buffer
	defer i.Reset()
	w := snappy.NewBufferedWriter(&i)
	defer w.Close()
	w.Write(v)
	w.Flush()
	return i.Bytes()
}

func (s *Snappy) Decode(v []byte) []byte {
	b := bytes.NewReader(v)
	var o bytes.Buffer
	defer o.Reset()

	r := snappy.NewReader(b)
	io.Copy(&o, r)
	return o.Bytes()
}
