package anthropic

import (
	"github.com/3JoB/resty-ilo"
	"github.com/3JoB/ulib/litefmt"
	// "github.com/google/uuid"

	"github.com/3JoB/anthropic-sdk-go/data"
)

type Opts struct {
	Message   data.MessageModule // Chunked message structure
	ContextID string             // Session ID. If empty, a new session is automatically created. If not empty, an attempt is made to find an existing session.
	Sender    Sender
}

func (opts *Opts) newCtx() *Context {
	return &Context{
		Response: &Response{},
		Human:    opts.Message.Human,
	}
}

// Send data to the API endpoint. Before sending out, the data will be processed into a form that the API can recognize.
//
// This method will be used to handle context requests.
//
// The context parameter comes from *Context.CtxData, please do not modify or process it by yourself, the context will be automatically processed when the previous request is executed.
/*func (ah *AnthropicClient) SendWithContext(sender *Sender, context string) (ctx *Context, err error) {
	if err := ah.c(sender); err != nil {
		return nil, err
	}
	sender.Prompt, err = addPrompt(context, sender.Prompt)
	if err != nil {
		return nil, err
	}
	return sender.Complete(ah.client)
}*/

func (c *Client) initHeaders() error {
	if c.Key == "" {
		return data.ErrApiKeyEmpty
	}
	c.client = resty.New().SetBaseURL(API).SetHeaders(map[string]string{
		"Accept":            "application/json",
		"Content-Type":      "application/json",
		"Client":            litefmt.Sprint("anthropic-sdk-go/", SDKVersion),
		"anthropic-version": "2023-06-01",
		"x-api-key":         c.Key,
		"User-Agent":        UserAgent,
	})
	return nil
}
