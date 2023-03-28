package anthropic

import "sync"

type Context struct {
	ID       string // Context ID
	Human    string
	RawData  string // Unprocessed raw json data returned by the API endpoint
	Response *Response
}

type MessageModule struct {
	Assistant string // returned data (do not modify)
	Human     string // input content
}

var pool sync.Map

func (c *Context) Find() (v []MessageModule, ok bool) {
	return FindContext(c.ID)
}

func (c *Context) Set(value any) bool {
	return SetContext(c.ID, value)
}

func (c *Context) Add(human string) bool {
	return AddContext(c.ID, MessageModule{Assistant: c.Response.Completion, Human: human})
}

func (c *Context) Delete() {
    DeleteContext(c.ID)
}

func (c *Context) Refresh() {
    RefreshContext()
}

func AddContext(key string, value MessageModule) bool {
	v, ok := FindContext(key)
	if !ok {
		return SetContext(key, value)
	}
	v = append(v, value)
	SetContext(key, v)
	return true
}

func FindContext(key string) (v []MessageModule, ok bool) {
	vs, ok := pool.Load(key)
	if !ok {
		return nil, ok
	}
	return vs.([]MessageModule), ok
}

func SetContext(key string, value any) bool {
	switch v := value.(type) {
	case MessageModule:
		r := []MessageModule{
			v,
		}
		pool.Store(key, r)
	case []MessageModule:
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