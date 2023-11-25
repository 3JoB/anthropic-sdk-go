// +build windows

package data

import (
	"github.com/mattn/go-ieproxy"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

func dialer() fasthttp.DialFunc {
	ieproxy.OverrideEnvWithStaticProxy()
	return fasthttpproxy.FasthttpProxyHTTPDialer()
}
