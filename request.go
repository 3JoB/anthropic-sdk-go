package anthropic

import (
	"github.com/3JoB/ulib/err"
	"github.com/3JoB/ulib/json"
	"github.com/go-resty/resty/v2"
)

func (req *Sender) Complete() (res *Response, errs error) {
	r, errs := resty.New().R().SetHeaders(Headers).SetBody(json.Marshal(req).Bytes()).Post(APIComplete)
	if errs != nil {
		return nil, &err.Err{Op: "request_Complete", Err: errs.Error()}
	}
	defer r.RawBody().Close()
	if errs := json.Unmarshal(r.Body(), res); errs != nil {
		return nil, &err.Err{Op: "request_Complete", Err: errs.Error()}
	}
	if r.StatusCode() != 200 {
		return nil, &err.Err{Op: "request_Complete", Err: res.Detail.(string)}
	}
	res.Context, _ = setPrompt(req.Prompt, res.Completion)
	return res, nil
}
