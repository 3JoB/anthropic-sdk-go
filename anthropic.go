package anthropic

import (
	"errors"
	"sync"
)

// Create a new Client object.
func New(conf *Client) (*Client, error) {
	if conf == nil {
		return nil, errors.New("client is nil")
	}
	if err := conf.headers(); err != nil {
		return nil, err
	}
	if conf.DefaultModel == "" {
		conf.DefaultModel = Model.Major.Instant1
	}
	return conf, nil
}

func NewPool(conf *Client) sync.Pool {
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
