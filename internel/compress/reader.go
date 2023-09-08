package compress

import "io"

// Read io.Reader and return []byte
func reader(i io.Reader) []byte {
	r, _ := io.ReadAll(i)
	return r
}
