package anthropic

import (
	"sync"

	"github.com/3JoB/ulib/litefmt"
	"github.com/valyala/fasthttp"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/pkg/pool"
)

// Create a new Client object.
func New(c *Config) (*Client, error) {
	client := &Client{
		key:   c.Key,
		model: c.DefaultModel,
		header: map[string]string{
			"Accept":            "application/json",
			"Content-Type":      "application/json",
			"Client":            litefmt.Sprint("anthropic-sdk-go/", data.SDKVersion),
			"anthropic-version": "2023-06-01",
			"User-Agent":        data.UserAgent,
			"x-api-key":         c.Key,
		},
		pool:   pool.NewPool(),
		client: data.Client,
	}
	if c.Compress != nil {
		_ = client.pool.UseCompress(c.Compress)
	}
	if client.model == "" {
		client.model = data.ModelMajorInstant
	}
	return client, nil
}

// Create a new Client object with a sync.Pool.
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

// Acquire returns an empty fasthttp instance from request pool.
//
// The returned fasthttp instance may be passed to Release when it is no longer needed.
// This allows Request recycling, reduces GC pressure and usually improves performance.
func (c *Client) Acquire() (*fasthttp.Request, *fasthttp.Response) {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	for k, v := range c.header {
		req.Header.Set(k, v)
	}
	req.SetRequestURI(data.API)
	req.Header.SetMethod("POST")
	return req, resp
}

// Release returns req and resp acquired via Acquire to request pool.
//
// It is forbidden accessing req and/or its' members after returning it to request pool.
func release(req *fasthttp.Request, res *fasthttp.Response) {
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}
