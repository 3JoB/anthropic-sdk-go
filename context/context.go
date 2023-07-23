package context

import (
	"sync"
	"time"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/internel/compress"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

type CacheItem struct {
	Key        string
	Value      string
	String     bool
	Compress   bool
	IO         compress.Interface
	ExpireTime time.Time
}

type ContextItem struct {
}

type Context struct {
	ID       string // Context ID
	Human    string
	RawData  string // Unprocessed raw json data returned by the API endpoint
	Response *resp.Response
}

var pool sync.Map = sync.Map{}

func (c *Context) Find() (v []data.MessageModule, ok bool) {
	return _FindContext(c.ID)
}

func (c *Context) Set(value any) bool {
	return _SetContext(c.ID, value)
}

// Add a prompt to the context storage pool
func (c *Context) Add() bool {
	return _AddContext(c.ID, data.MessageModule{Assistant: c.Response.Completion, Human: c.Human})
}

// Deprecated: Use Close() instead.
func (c *Context) Delete() {
	_DeleteContext(c.ID)
}

func (c *Context) Close() error {
	_DeleteContext(c.ID)
	return nil
}

// Refresh the context storage pool (clear all data)
func (c *Context) Refresh() {
	RefreshContext()
}

func _AddContext(key string, value data.MessageModule) bool {
	v, ok := _FindContext(key)
	if !ok {
		return _SetContext(key, value)
	}
	v = append(v, value)
	_SetContext(key, v)
	return true
}

func _FindContext(key string) (v []data.MessageModule, ok bool) {
	if vs, ok := pool.Load(key); !ok {
		return nil, ok
	} else {
		return vs.([]data.MessageModule), ok
	}
}

func _SetContext(key string, value any) bool {
	switch v := value.(type) {
	case data.MessageModule:
		r := []data.MessageModule{
			v,
		}
		pool.Store(key, r)
	case []data.MessageModule:
		pool.Store(key, v)
	default:
		return false
	}
	return true
}

func _DeleteContext(key string) {
	pool.Delete(key)
}

func RefreshContext() {
	pool.Range(func(key, value any) bool {
		pool.Delete(key)
		return true
	})
}
