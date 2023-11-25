package data

import (
	"github.com/valyala/fasthttp"
)

var Client = &fasthttp.Client{
	NoDefaultUserAgentHeader:      true,
	DisableHeaderNamesNormalizing: false,
	Dial:                          dialer(),
}
