package anthropic

import (
	"sync"

	"github.com/cornelk/hashmap"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
)

// Create a new Client object.
func New[T []data.MessageModule | string](c *Config[T]) (*Client[T], error) {
	if c == nil {
		return nil, data.ErrConfigEmpty
	}
	client := &Client[T]{
		cfg:    c,
		header: hashmap.New[string, string](),
		client: &fasthttp.Client{
			NoDefaultUserAgentHeader:      true,
			DisableHeaderNamesNormalizing: false,
			Dial:                          fasthttpproxy.FasthttpProxyHTTPDialer(),
		},
	}
	if err := client.headers(); err != nil {
		return nil, err
	}
	if c.DefaultModel == "" {
		c.DefaultModel = data.ModelMajorInstant
	}
	return client, nil
}

func NewPool[T []data.MessageModule | string](c *Config[T]) sync.Pool {
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
