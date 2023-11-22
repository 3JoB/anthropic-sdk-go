package anthropic

import (
	"github.com/3JoB/anthropic-sdk-go/v2/pkg/pool"
)

// Anthropic-SDK-Go configuration
type Config struct {
	Key          string         // API Keys
	DefaultModel string         // Choose the default AI model
	DataPool     DataPoolConfig // Pool Config
}

type DataPoolConfig struct {
	Pool     *pool.Pool
	Compress bool
}
