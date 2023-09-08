package data

const (
	API        string = "https://api.anthropic.com/v1/complete"
	UserAgent  string = "Mozilla/5.0 (compatible; anthropic-sdk-go/2.0.3-Beta; +https://github.com/3JoB/anthropic-sdk-go/;) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"
	SDKVersion string = "2.0.3-Beta"
)

var StopSequences []string = []string{"\n\nHuman:"}
