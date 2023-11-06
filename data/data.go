package data

import "github.com/3JoB/anthropic-sdk-go/v2/internel/errors"

var (
	ErrApiKeyEmpty      = errors.New("APIKey cannot be empty!")
	ErrRegionBanned     = errors.New("Region blocked")
	ErrClientIsNil      = errors.New("Client cannot be nil!")
	ErrContextIsNil     = errors.New("Context cannot be nil!")
	ErrContextNotFound  = errors.New("Context not found")
	ErrConfigIsNil      = errors.New("Configuration cannot be nil!")
	ErrSenderIsNil      = errors.New("Sender cannot be nil!")
	ErrPromptHumanEmpty = errors.New("The value of human cannot be empty!")
	ErrPromptCtxEmpty   = errors.New("The value of context cannot be empty!")
	ErrPromptEmpty      = errors.New("The value of prompt cannot be empty!")
)

// Chunked message structure
type MessageModule struct {
	Assistant string // returned data (do not modify)
	Human     string // input content
}
