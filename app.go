package anthropic

// Deprecated: This method will be deprecated in v2 sdk
// stable version and use new implementation.
type AppModel struct {
	Major ModelMajor
	Full  ModelFull
}

// Deprecated: This method will be deprecated in v2 sdk
// stable version and use new implementation.
type ModelMajor struct {
	Instant1 string
	Claude2  string
}

// Deprecated: This method will be deprecated in v2 sdk
// stable version and use new implementation.
type ModelFull struct {
	Instant1 string
	Claude2  string
}

// Deprecated: This method will be deprecated in v2 sdk
// stable version and use new implementation.
var Model = AppModel{
	Major: ModelMajor{
		Instant1: "claude-instant-1",
		Claude2:  "claude-2",
	},
	Full: ModelFull{
		Instant1: "claude-instant-1.2",
		Claude2:  "claude-2.0",
	},
}

const (
	API        string = "https://api.anthropic.com/v1/complete"
	UserAgent  string = "Mozilla/5.0 (compatible; anthropic-sdk-go/2.0.2-Beta; +https://github.com/3JoB/anthropic-sdk-go/;) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"
	SDKVersion string = "2.0.2-Beta"
)

var StopSequences []string = []string{"\n\nHuman:"}
