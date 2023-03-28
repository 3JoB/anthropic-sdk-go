package anthropic

import (
	"github.com/3JoB/ulib/err"
	"github.com/3JoB/ulib/json"
	"github.com/3JoB/unsafeConvert"
	"github.com/go-resty/resty/v2"
)

// Make a processed request to an API endpoint.
func (req *Opts) Complete(ctx *Context, client *resty.Client) (*Context, error) {
	r, errs := client.R().SetBody(json.Marshal(req.Sender).Bytes()).Post(APIComplete)
	if errs != nil {
		return nil, &err.Err{Op: "request", Err: errs.Error()}
	}
	defer r.RawBody().Close()

	ctx.ID = req.ContextID
	if errs := json.Unmarshal(r.Body(), ctx.Response); errs != nil {
		return nil, &err.Err{Op: "request", Err: errs.Error()}
	}

	ctx.RawData = unsafeConvert.StringReflect(r.Body())

	if r.StatusCode() != 200 {
		return nil, &err.Err{Op: "request", Err: ctx.RawData}
	}

	req.Context.Assistant = ctx.Response.Completion

	if !ctx.Add() {
		return nil, &err.Err{Op: "request", Err: "Add failed"}
	}

	return ctx, nil
}
