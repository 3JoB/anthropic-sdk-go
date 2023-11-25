package data

const (
	API        string = "https://api.anthropic.com/v1/complete"
	UserAgent  string = "Mozilla/5.0 (compatible; anthropic-sdk-go/2.1.0; +https://github.com/3JoB/anthropic-sdk-go/;) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
	SDKVersion string = "2.1.0"
)

var StopSequences []string = []string{"\n\nHuman:"}
