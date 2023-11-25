package data

import "errors"

var (
	ErrApiKeyEmpty      = errors.New("apikey cannot be empty")
	ErrRegionBanned     = errors.New("region blocked")
	ErrClientIsNil      = errors.New("client cannot be nil")
	ErrSessionIsNil     = errors.New("session cannot be nil")
	ErrSessionNotFound  = errors.New("session not found")
	ErrConfigIsNil      = errors.New("configuration cannot be nil")
	ErrSenderIsNil      = errors.New("sender cannot be nil")
	ErrPromptHumanEmpty = errors.New("the value of human cannot be empty")
	ErrPromptCtxEmpty   = errors.New("the value of session cannot be empty")
	ErrPromptEmpty      = errors.New("the value of prompt cannot be empty")
)

// Chunked message structure
type MessageModule struct {
	Assistant string // returned data (do not modify)
	Human     string // input content
}
