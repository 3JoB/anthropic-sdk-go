package anthropic

import (
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/pkg/hashpool"
)

// Anthropic-SDK-Go configuration
type Config[T []data.MessageModule | string] struct {
	Key          string        // API Keys
	DefaultModel string        // Choose the default AI model
	PoolConfig   PoolConfig[T] // Pool Config
}

// Pool Config
type PoolConfig[T []data.MessageModule | string] struct {
	WorkerPool WorkerPoolConfig
	HashPool   HashPoolConfig[T]
}

type WorkerPoolConfig struct {
	Process int
}

type HashPoolConfig[T []data.MessageModule | string] struct {
	Pool     hashpool.Interface[T]
	Compress bool // Only CachePool
}
