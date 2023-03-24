package anthropic

type Context struct {
	CtxData  string // This is automatically processed context data, do not modify
	ID string
	RawData  string // Unprocessed raw json data returned by the API endpoint
	Response *Response
}

type MessageModule struct{
	Assistant string
	Human string
}

var contextMaps map[string][]MessageModule

func init(){
	contextMaps = make(map[string][]MessageModule)
}

func FindContext(key string) (v []MessageModule) {
	return contextMaps[key]
}

func (c *Context) Find() (v []MessageModule) {
	return FindContext(c.ID)
}

func SetContext(key string, value []MessageModule) {}

func (c *Context) Set() {}

func AddContextMaps(key string, value MessageModule) {
	
}