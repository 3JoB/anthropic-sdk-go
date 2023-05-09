package anthropic

import (
	"github.com/3JoB/resty-ilo"
	"github.com/3JoB/ulib/err"
)

// Make a processed request to an API endpoint.
func (req *Opts) Complete(ctx *Context, client *resty.Client) (*Context, error) {
	r, errs := client.R().SetBody(req.Sender).Post(APIComplete)
	if errs != nil {
		return ctx, &err.Err{Op: "request", Err: errs.Error()}
	}
	defer r.RawBody().Close()
	r.RawResponse.Close = true

	ctx.ID = req.ContextID
	if errs := r.Bind(ctx.Response); errs != nil {
		return ctx, &err.Err{Op: "request", Err: errs.Error()}
	}

	ctx.RawData = r.String()

	if !r.IsStatusCode(200) {
		return ctx, &err.Err{Op: "request", Err: ctx.RawData}
	}

	req.Message.Assistant = ctx.Response.Completion

	if !ctx.Add() {
		return ctx, &err.Err{Op: "request", Err: "Add failed"}
	}

	return ctx, nil
}
