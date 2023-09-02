package anthropic

import (
	"sync"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
)

// Create a new Client object.
func New(conf *Config) (*Client, error) {
	if conf == nil {
		return nil, data.ErrConfigEmpty
	}
	client := &Client{
		cfg: conf,
	}
	if err := client.headers(); err != nil {
		return nil, err
	}
	if conf.DefaultModel == "" {
		conf.DefaultModel = Model.Major.Instant1
	}
	return client, nil
}

func NewPool(conf *Config) sync.Pool {
	return sync.Pool{
		New: func() any {
			if client, err := New(conf); err != nil {
				panic(err)
			} else {
				return client
			}
		},
	}
}
