package context

import (
	"sync"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

type Context struct {
	ID        string // Context ID
	Human     string
	RawData   []byte // Unprocessed raw json data returned by the API endpoint
	Response  *resp.Response
	ErrorResp *resp.ErrorResponse
}

// Deprecated: Due to design and performance issues,
// ContextPool will be deprecated in the v2 sdk stable
// version soon, and the relevant code will be removed
// at that time. Please prepare to migrate to HashPool
// which will replace it as soon as possible.
var pool sync.Map = sync.Map{}

// Deprecated: Due to design and performance issues,
// ContextPool will be deprecated in the v2 sdk stable
// version soon, and the relevant code will be removed
// at that time. Please prepare to migrate to HashPool
// which will replace it as soon as possible.
func (c *Context) Find() (v []data.MessageModule, ok bool) {
	return _FindContext(c.ID)
}

// Deprecated: Due to design and performance issues,
// ContextPool will be deprecated in the v2 sdk stable
// version soon, and the relevant code will be removed
// at that time. Please prepare to migrate to HashPool
// which will replace it as soon as possible.
func (c *Context) Set(value any) bool {
	return _SetContext(c.ID, value)
}

// Add a prompt to the context storage pool
//
// Deprecated: Due to design and performance issues,
// ContextPool will be deprecated in the v2 sdk stable
// version soon, and the relevant code will be removed
// at that time. Please prepare to migrate to HashPool
// which will replace it as soon as possible.
func (c *Context) Add() bool {
	return _AddContext(c.ID, data.MessageModule{Assistant: c.Response.Completion, Human: c.Human})
}

// Deprecated: Due to design and performance issues,
// ContextPool will be deprecated in the v2 sdk stable
// version soon, and the relevant code will be removed
// at that time. Please prepare to migrate to HashPool
// which will replace it as soon as possible.
func (c *Context) Close() {
	_DeleteContext(c.ID)
}

// Refresh the context storage pool (clear all data)
//
// Deprecated: Due to design and performance issues,
// ContextPool will be deprecated in the v2 sdk stable
// version soon, and the relevant code will be removed
// at that time. Please prepare to migrate to HashPool
// which will replace it as soon as possible.
func (c *Context) Refresh() {
	RefreshContext()
}

// Deprecated: Due to design and performance issues,
// ContextPool will be deprecated in the v2 sdk stable
// version soon, and the relevant code will be removed
// at that time. Please prepare to migrate to HashPool
// which will replace it as soon as possible.
func _AddContext(key string, value data.MessageModule) bool {
	v, ok := _FindContext(key)
	if !ok {
		return _SetContext(key, value)
	}
	v = append(v, value)
	_ = _SetContext(key, v)
	return true
}

// Deprecated: Due to design and performance issues,
// ContextPool will be deprecated in the v2 sdk stable
// version soon, and the relevant code will be removed
// at that time. Please prepare to migrate to HashPool
// which will replace it as soon as possible.
func _FindContext(key string) (v []data.MessageModule, ok bool) {
	if vs, ok := pool.Load(key); !ok {
		return nil, ok
	} else {
		return vs.([]data.MessageModule), ok
	}
}

// Deprecated: Due to design and performance issues,
// ContextPool will be deprecated in the v2 sdk stable
// version soon, and the relevant code will be removed
// at that time. Please prepare to migrate to HashPool
// which will replace it as soon as possible.
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

// Deprecated: Due to design and performance issues,
// ContextPool will be deprecated in the v2 sdk stable
// version soon, and the relevant code will be removed
// at that time. Please prepare to migrate to HashPool
// which will replace it as soon as possible.
func _DeleteContext(key string) {
	pool.Delete(key)
}

// Deprecated: Due to design and performance issues,
// ContextPool will be deprecated in the v2 sdk stable
// version soon, and the relevant code will be removed
// at that time. Please prepare to migrate to HashPool
// which will replace it as soon as possible.
func RefreshContext() {
	pool.Range(func(key, value any) bool {
		pool.Delete(key)
		return true
	})
}
