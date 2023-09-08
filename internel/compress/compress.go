package compress

import "bytes"

// The Interface interface is a unified interface of compress,
// and the custom interface is not available for now.
type Interface interface {
	// The Decode method will first decode and then overwrite the data in the input *bytes.Buffer.
	Encode([]byte) (*bytes.Buffer, error)

	Decode(*bytes.Buffer)
}
