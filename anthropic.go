package anthropic

import (
	"sync"

	"github.com/cornelk/hashmap"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
)

// Create a new Client object.
func New(c *Config) (*Client, error) {
	if c == nil {
		return nil, data.ErrConfigEmpty
	}
	client := &Client{
		cfg:    c,
		header: hashmap.New[string, string](),
	}
	client.setDefaultClient()
	if err := client.headers(); err != nil {
		return nil, err
	}
	if c.DefaultModel == "" {
		c.DefaultModel = Model.Major.Instant1
	}
	return client, nil
}

func NewPool(c *Config) sync.Pool {
	return sync.Pool{
		New: func() any {
			if client, err := New(c); err != nil {
				panic(err)
			} else {
				return client
			}
		},
	}
}
