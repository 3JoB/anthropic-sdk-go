package anthropic

import (
	"errors"
	"io"

	"github.com/3JoB/unsafeConvert"
	"github.com/sugawarayuuta/sonnet"

	"github.com/3JoB/anthropic-sdk-go/v2/context"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

type Sender struct {
	Message   data.MessageModule // Chunked message structure
	ContextID string             // Session ID. If empty, a new session is automatically created. If not empty, an attempt is made to find an existing session.
	Sender    resp.Sender
}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) newCtx() *context.Context {
	return &context.Context{
		Response: resp.Response{},
		Human:    s.Message.Human,
	}
}

func (s *Sender) SetHuman(v string) {}

// Make a processed request to an API endpoint.
func (s *Sender) Complete(client *Client, ctx *context.Context) (*context.Context, error) {
	// Get fasthttp object
	request, response := client.Acquire()
	defer release(request, response)
	if errs := s.setBody(request.BodyWriter()); errs != nil {
		return nil, errs
	}

	if errs := client.do(request, response); errs != nil {
		return ctx, errs
	}

	ctx.ID = s.ContextID
	if errs := sonnet.Unmarshal(response.Body(), &ctx.Response); errs != nil {
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
func (opt *Sender) setBody(w io.Writer) error {
	return sonnet.NewEncoder(w).Encode(&opt.Sender)
}
