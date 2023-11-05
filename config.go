package anthropic

import (
	"github.com/3JoB/anthropic-sdk-go/v2/pkg/hashpool"
)

// Anthropic-SDK-Go configuration
type Config struct {
	Key          string        // API Keys
	DefaultModel string        // Choose the default AI model
	PoolConfig   PoolConfig // Pool Config
}

// Pool Config
type PoolConfig struct {
	WorkerPool WorkerPoolConfig
	DataPool   DataPoolConfig
}

type WorkerPoolConfig struct {
	Process int
}

type DataPoolConfig struct {
	Pool     *hashpool.Pool
	Compress bool // Only CachePool
}
