package anthropic

import (
	"github.com/3JoB/ulib/err"
)

const (
	API         string = "https://api.anthropic.com"
	APIComplete string = "/v1/complete"
	SDKVersion  string = "1.1.0"

	ModelClaudeV1             string = "claude-v1"
	ModelClaudeDefault        string = "claude-v1.0"
	ModelClaudeV12            string = "claude-v1.2"
	ModelClaudeInstantV1      string = "claude-instant-v1"
	ModelClaudeInstantDefault string = "claude-instant-v1.0"
)

var (
	Headers       map[string]string
	StopSequences []string = []string{
		"\n\nHuman:",
	}

	ErrApiKeyEmpty      error = &err.Err{Op: "config", Err: "APIKey cannot be empty!"}
	ErrConfigEmpty      error = &err.Err{Op: "config", Err: "Configuration cannot be empty!"}
	ErrSenderNil        error = &err.Err{Op: "sender", Err: "Sender cannot be nil!"}
	ErrPromptHumanEmpty error = &err.Err{Op: "prompt", Err: "The value of human cannot be empty!"}
	ErrPromptCtxEmpty   error = &err.Err{Op: "prompt", Err: "The value of context cannot be empty!"}
	ErrPromptEmpty      error = &err.Err{Op: "send", Err: "The value of prompt cannot be empty!"}
)
