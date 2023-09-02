package compress

type Interface interface {
	Encode([]byte) []byte
	Decode([]byte) []byte
}

func New() {}
