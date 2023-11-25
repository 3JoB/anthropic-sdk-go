package anthropic

import "github.com/3JoB/anthropic-sdk-go/v2/pkg/compress"

// Anthropic-SDK-Go configuration
type Config struct {
	Key          string             // API Keys
	DefaultModel string             // Choose the default AI model
	Compress     compress.Interface // Data Compress
}
