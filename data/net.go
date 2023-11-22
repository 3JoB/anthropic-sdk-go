package data

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

var Client = &fasthttp.Client{
	NoDefaultUserAgentHeader:      true,
	DisableHeaderNamesNormalizing: false,
	Dial:                          fasthttpproxy.FasthttpProxyHTTPDialer(),
}
