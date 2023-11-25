package anthropic

import "github.com/3JoB/ulib/litefmt"

type Context struct {
	id string
}

func (c *Client) NewContext(id ...string) *Context {
	ctx := &Context{}
	if len(id) > 0 {
		ctx.id = litefmt.Sprint(id...)
	}
	return ctx
}

func (c *Client) Close() error {
	return nil
}