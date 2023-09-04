package compress

import "bytes"

//	Done

type Interface interface {
	Encode([]byte) *bytes.Buffer
	Decode(*bytes.Buffer) []byte
}
