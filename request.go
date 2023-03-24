package anthropic

import (
	"github.com/3JoB/ulib/err"
	"github.com/3JoB/ulib/json"
	"github.com/3JoB/unsafeConvert"
	"github.com/go-resty/resty/v2"
)

func (req *Sender) Complete() (*Context, error) {
	r, errs := resty.New().R().SetHeaders(Headers).SetBody(json.Marshal(req).Bytes()).Post(APIComplete)
	if errs != nil {
		return nil, &err.Err{Op: "request_Complete", Err: errs.Error()}
	}
	ctx := &Context{
		Response: &Response{},
	}
	if errs := json.Unmarshal(r.Body(), ctx.Response); errs != nil {
		return nil, &err.Err{Op: "request_Complete", Err: errs.Error()}
	}
	if r.StatusCode() != 200 {
		return nil, &err.Err{Op: "request_Complete", Err: ctx.Response.Detail.(string)}
	}
	ctx.RawData = unsafeConvert.StringReflect(r.Body())
	ctx.CtxData, _ = setPrompt(req.Prompt, ctx.Response.Completion)
	return ctx, nil
}
