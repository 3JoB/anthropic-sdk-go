//go:build windows
// +build windows

package data

import "github.com/mattn/go-ieproxy"

// use ieproxy to automatically read proxy from the system.
//
// to avoid cgo, only build on windows.
func init() {
	ieproxy.OverrideEnvWithStaticProxy()
}
