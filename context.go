package anthropic

import (
	"sync"

	"github.com/3JoB/anthropic-sdk-go/data"
)

type Context struct {
	ID       string // Context ID
	Human    string
	RawData  string // Unprocessed raw json data returned by the API endpoint
	Response *Response
}

var pool sync.Map = sync.Map{}

func (c *Context) Find() (v []data.MessageModule, ok bool) {
	return FindContext(c.ID)
}

func (c *Context) Set(value any) bool {
	return SetContext(c.ID, value)
}

func (c *Context) Add() bool {
	return AddContext(c.ID, data.MessageModule{Assistant: c.Response.Completion, Human: c.Human})
}

func (c *Context) Delete() {
    DeleteContext(c.ID)
}

func (c *Context) Refresh() {
    RefreshContext()
}

func AddContext(key string, value data.MessageModule) bool {
	v, ok := FindContext(key)
	if !ok {
		return SetContext(key, value)
	}
	v = append(v, value)
	SetContext(key, v)
	return true
}

func FindContext(key string) (v []data.MessageModule, ok bool) {
	vs, ok := pool.Load(key)
	if !ok {
		return nil, ok
	}
	return vs.([]data.MessageModule), ok
}

func SetContext(key string, value any) bool {
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

func DeleteContext(key string) {
	pool.Delete(key)
}

func RefreshContext() {
	pool.Range(func(key, value any) bool {
		pool.Delete(key)
        return true
	})
}