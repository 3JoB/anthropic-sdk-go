package anthropic

import (
	"io"

	"github.com/3JoB/ulib/err"
	"github.com/3JoB/unsafeConvert"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/encoder"
	// "github.com/google/uuid"

	"github.com/3JoB/anthropic-sdk-go/v2/context"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

type Sender struct {
	Message   data.MessageModule // Chunked message structure
	ContextID string             // Session ID. If empty, a new session is automatically created. If not empty, an attempt is made to find an existing session.
	Sender    resp.Sender
	client    *Client
}

func NewSender() *Sender {
	return &Sender{}
}

// Deprecated: This method will be deprecated in v2 sdk
// stable version and use new implementation.
func (s *Sender) newCtx() *context.Context {
	return &context.Context{
		Response: &resp.Response{},
		Human:    s.Message.Human,
	}
}

// Deprecated: This method will be deprecated in v2 sdk
// stable version and use new implementation.
func (s *Sender) With(client *Client) {
	s.client = client
}

func (s *Sender) SetHuman(v string) {

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

// Make a processed request to an API endpoint.
//
// Deprecated: This method will be deprecated in v2 sdk
// stable version and use new implementation.
func (opt *Sender) Complete(ctx *context.Context) (*context.Context, error) {
	// Get fasthttp object
	request, response := acquire()
	defer release(request, response)

	// Initialize Request
	opt.client.setHeaderWithURI(request)
	if errs := opt.setBody(request.BodyWriter()); errs != nil {
		return nil, errs
	}

	if errs := opt.client.do(request, response); errs != nil {
		return ctx, &err.Err{Op: "opts:61", Err: errs.Error()}
	}

	ctx.ID = opt.ContextID
	if errs := sonic.Unmarshal(response.Body(), ctx.Response); errs != nil {
		return ctx, &err.Err{Op: "opts:66", E: errs}
	}

	ctx.RawData = response.Body()

	if response.StatusCode() != 200 {
		errs, _ := resp.Error(response.StatusCode(), response.Body())
		if errs != nil {
			ctx.ErrorResp = errs
			return ctx, errs.Err()
		}
		return ctx, &err.Err{Op: "opts:77", Err: unsafeConvert.StringSlice(ctx.RawData)}
	}

	opt.Message.Assistant = ctx.Response.Completion

	if !ctx.Add() {
		return ctx, &err.Err{Op: "opts:83", Err: "Add failed"}
	}

	return ctx, nil
}

// Set Body for *fasthttp.Request.
//
// Need to export io.Writer in BodyWriter() as w.
func (opt *Sender) setBody(w io.Writer) error {
	return encoder.NewStreamEncoder(w).Encode(&opt.Sender)
}
