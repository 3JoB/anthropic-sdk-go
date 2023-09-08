package anthropic

// Anthropic-SDK-Go configuration
type Config struct {
	Key          string // API Keys
	DefaultModel string // Choose the default AI model
	PoolConfig PoolConfig // Pool Config
}

// Pool Config
type PoolConfig struct{
	WorkerPool WorkerPoolConfig
	HashPool HashPoolConfig
}

type WorkerPoolConfig struct{

}

type HashPoolConfig struct{

}