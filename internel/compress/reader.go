package compress

import "io"

func reader(i io.Reader) []byte {
	r, _ := io.ReadAll(i)
	return r
}
