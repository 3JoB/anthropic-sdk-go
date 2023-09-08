package anthropic

import (
	"io"

	"github.com/bytedance/sonic/encoder"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

// Set Body for *fasthttp.Request.
//
// Need to export io.Writer in BodyWriter() as w.
func (opt *Opts) setBody(w io.Writer) error {
	return encoder.NewStreamEncoder(w).Encode(&opt.Sender)
}

// Initialize a fasthttp.Client object for Client
//
// Deprecated: This method will be deprecated in v2 sdk
// stable version and use new implementation.
func (c *Client) setDefaultClient() {
	c.client = &fasthttp.Client{
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: false,
		Dial:                          fasthttpproxy.FasthttpProxyHTTPDialer(),
	}
}

func (c *Client) setHeaderWithURI(req *fasthttp.Request) {
	c.header.Range(func(k, v string) bool {
		req.Header.Set(k, v)
		return true
	})
	req.SetRequestURI(API)
	req.Header.SetMethod("POST")
}

func (c *Client) do(req *fasthttp.Request, res *fasthttp.Response) error {
	return c.client.Do(req, res)
}

// Acquire returns an empty fasthttp instance from request pool.
//
// The returned fasthttp instance may be passed to Release when it is no longer needed.
// This allows Request recycling, reduces GC pressure and usually improves performance.
func acquire() (*fasthttp.Request, *fasthttp.Response) {
	return fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
}

// Release returns req and resp acquired via Acquire to request pool.
//
// It is forbidden accessing req and/or its' members after returning it to request pool.
func release(req *fasthttp.Request, res *fasthttp.Response) {
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}
