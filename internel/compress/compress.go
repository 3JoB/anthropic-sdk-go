package compress

type Interface interface {
	Encode(string) string
	Decode(string) string
}
