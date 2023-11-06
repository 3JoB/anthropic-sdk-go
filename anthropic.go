package anthropic

import (
	"sync"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/ulib/litefmt"
)

// Create a new Client object.
func New(c *Config) (*Client, error) {
	client := &Client{
		cfg: c,
		header: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
			"Client":       litefmt.Sprint("anthropic-sdk-go/", data.SDKVersion),
			"anthropic-version": "2023-06-01",
			"User-Agent": data.UserAgent,
			"x-api-key": c.Key,
		},
		client: &fasthttp.Client{
			NoDefaultUserAgentHeader:      true,
			DisableHeaderNamesNormalizing: false,
			Dial:                          fasthttpproxy.FasthttpProxyHTTPDialer(),
		},
	}
	if c.DefaultModel == "" {
		c.DefaultModel = data.ModelMajorInstant
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
