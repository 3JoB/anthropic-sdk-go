package data

import "github.com/3JoB/ulib/err"

var (
	ErrApiKeyEmpty      error = &err.Err{Op: "config", Err: "APIKey cannot be empty!"}
	ErrRegionBanned     error = &err.Err{Op: "403", Err: "Region blocked"}
	ErrContextNil       error = &err.Err{Op: "send", Err: "Context cannot be nil!"}
	ErrContextNotFound  error = &err.Err{Op: "send", Err: "Context not found"}
	ErrConfigEmpty      error = &err.Err{Op: "config", Err: "Configuration cannot be empty!"}
	ErrSenderNil        error = &err.Err{Op: "sender", Err: "Sender cannot be nil!"}
	ErrPromptHumanEmpty error = &err.Err{Op: "prompt", Err: "The value of human cannot be empty!"}
	ErrPromptCtxEmpty   error = &err.Err{Op: "prompt", Err: "The value of context cannot be empty!"}
	ErrPromptEmpty      error = &err.Err{Op: "send", Err: "The value of prompt cannot be empty!"}
)

// Chunked message structure
type MessageModule struct {
	Assistant string // returned data (do not modify)
	Human     string // input content
}
