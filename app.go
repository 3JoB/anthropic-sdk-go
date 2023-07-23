package anthropic

var Model = struct {
	Major struct {
		Instant1 string
		Claude2  string
	}
	Full struct {
		Instant1 string
		Claude2  string
	}
}{
	Major: struct {
		Instant1 string
		Claude2  string
	}{
		Instant1: "claude-instant-1",
		Claude2:  "claude-2",
	},
	Full: struct {
		Instant1 string
		Claude2  string
	}{
		Instant1: "claude-instant-1.1",
		Claude2:  "claude-2.0",
	},
}

const (
	API         string = "https://api.anthropic.com"
	APIComplete string = "/v1/complete"
	UserAgent   string = "Mozilla/5.0 (compatible; anthropic-sdk-go/1.7.0; +https://github.com/3JoB/anthropic-sdk-go/;) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"
	SDKVersion  string = "1.7.0"
)

var StopSequences []string = []string{"\n\nHuman:"}
