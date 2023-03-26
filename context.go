package anthropic

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

var cPools map[string][]MessageModule

func init() {
	cPools = make(map[string][]MessageModule)
}

func FindContext(key string) (v []MessageModule) {
	return cPools[key]
}

func (c *Context) Find() (v []MessageModule) {
	return FindContext(c.ID)
}

func SetContext(key string, value []MessageModule) {
	cPools[key] = value
}

func (c *Context) Set() {}

func SetLastContext(key string, value MessageModule) {
	num := len(cPools[key])
	cPools[key][num-1] = value
}

func (c *Context) SetLast() {}

func AddContextMaps(key string, value MessageModule) {
	cPools[key] = append(cPools[key], value)
}

func (c *Context) Add() {
	AddContextMaps(c.ID, MessageModule{Assistant: c.Response.Completion})
}

func RefreshContext(key string) {
	delete(cPools, key)
}
