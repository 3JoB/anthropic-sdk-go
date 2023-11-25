//go:build !windows
// +build !windows

package data

import "github.com/valyala/fasthttp/fasthttpproxy"

var dialer = fasthttpproxy.FasthttpProxyHTTPDialer()
