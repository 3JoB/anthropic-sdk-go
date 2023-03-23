package anthropic

const (
	API         string = "https://api.anthropic.com"
	APIComplete string = "https://api.anthropic.com/v1/complete"
	SDKVersion  string = "v1.0.0"

	ModelClaudeV1             string = "claude-v1"
	ModelClaudeDefault        string = "claude-v1.0"
	ModelClaudeV12            string = "claude-v1.2"
	ModelClaudeInstantV1      string = "claude-instant-v1"
	ModelClaudeInstantDefault string = "claude-instant-v1.0"
)

var (
	Headers map[string]string
)
