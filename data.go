package anthropic

const (
	API string = "https://api.anthropic.com"

	SDKVersion string = "v1.0.0"
)

var (
	Headers map[string]string

	ErrApiKeyEmpty error = &Err{Op: "config", Err: "APIKey cannot be empty!"}
	ErrConfigEmpty error = &Err{Op: "config", Err: "Configuration cannot be empty!"}
	ErrHumanEmpty error = &Err{Op: "prompt", Err: "The value of human cannot be empty!"}
)