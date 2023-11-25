// +build !windows

package data

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

func dialer() fasthttp.DialFunc {
	return fasthttpproxy.FasthttpProxyHTTPDialer()
}

