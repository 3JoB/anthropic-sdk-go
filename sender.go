package anthropic

import (
	"errors"
	"io"

	"github.com/3JoB/unsafeConvert"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/encoder"
	// "github.com/google/uuid"

	"github.com/3JoB/anthropic-sdk-go/v2/context"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

type Sender[T []data.MessageModule | string] struct {
	Message   data.MessageModule // Chunked message structure
	ContextID string             // Session ID. If empty, a new session is automatically created. If not empty, an attempt is made to find an existing session.
	Sender    resp.Sender
	client    *Client[T]
}

func NewSenderWithSlice() *Sender[[]data.MessageModule] {
	return &Sender[[]data.MessageModule]{}
}

func NewSenderWithCache() *Sender[string] {
	return &Sender[string]{}
}

func (s *Sender[T]) newCtx() *context.Context {
	return &context.Context{
		Response: resp.Response{},
		Human:    s.Message.Human,
	}
}

func (s *Sender[T]) With(client *Client[T]) {
	s.client = client
}

func (s *Sender[T]) SetHuman(v string) {}

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
func (s *Sender[T]) Complete(ctx *context.Context) (*context.Context, error) {
	// Get fasthttp object
	request, response := acquire()
	defer release(request, response)

	// Initialize Request
	s.client.setHeaderWithURI(request)
	if errs := s.setBody(request.BodyWriter()); errs != nil {
		return nil, errs
	}

	if errs := s.client.do(request, response); errs != nil {
		return ctx, errs
	}

	ctx.ID = s.ContextID
	if errs := sonic.Unmarshal(response.Body(), &ctx.Response); errs != nil {
		return ctx, errs
	}

	ctx.RawData = response.Body()

	if response.StatusCode() != 200 {
		errs, _ := resp.Error(response.StatusCode(), response.Body())
		if errs != nil {
			ctx.ErrorResp = errs
			return ctx, errs
		}
		return ctx, errors.New(unsafeConvert.StringSlice(ctx.RawData))
	}

	s.Message.Assistant = ctx.Response.Completion

	if !ctx.Add() {
		return ctx, errors.New("add failed")
	}

	return ctx, nil
}

// Set Body for *fasthttp.Request.
//
// Need to export io.Writer in BodyWriter() as w.
func (opt *Sender[T]) setBody(w io.Writer) error {
	return encoder.NewStreamEncoder(w).Encode(&opt.Sender)
}
