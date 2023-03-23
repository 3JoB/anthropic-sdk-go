package anthropic

const (
	API string = "https://api.anthropic.com"

	SDKVersion string = "v1.0.0"

	ModelClaudeV1             string = "claude-v1"
	ModelClaudeDefault        string = "claude-v1.0"
	ModelClaudeV12            string = "claude-v1.2"
	ModelClaudeInstantV1      string = "claude-instant-v1"
	ModelClaudeInstantDefault string = "claude-instant-v1.0"
)

var (
	Headers map[string]string

	ErrApiKeyEmpty error = &Err{Op: "config", Err: "APIKey cannot be empty!"}
	ErrConfigEmpty error = &Err{Op: "config", Err: "Configuration cannot be empty!"}
	ErrHumanEmpty  error = &Err{Op: "prompt", Err: "The value of human cannot be empty!"}
	ErrPromptEmpty error = &Err{Op: "send", Err: "The value of prompt cannot be empty!"}
)
