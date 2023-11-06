package anthropic

import (
	"time"

	"github.com/3JoB/ulid"
	"github.com/valyala/fasthttp"
	"pgregory.net/rand"

	"github.com/3JoB/anthropic-sdk-go/v2/context"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

type Client struct {
	client  *fasthttp.Client
	cfg     *Config           // Config
	header  map[string]string // http
	timeout time.Duration
}

// Set the response timeout in minutes.
func (ah *Client) SetTimeOut(times int) {
	if times == 0 {
		return
	}
	ah.timeout = time.Duration(times) * time.Minute
}

// Send data to the API endpoint. Before sending out,
// the data will be processed into a form that the API can recognize.
func (ah *Client) Send(sender *Sender) (*context.Context, error) {
	var err error
	if err = ah.check(&sender.Sender); err != nil {
		return nil, err
	}
	if (sender.Message == data.MessageModule{}) {
		return nil, data.ErrContextIsNil
	}
	ctx := sender.newCtx()
	if sender.ContextID == "" {
		id, _ := ulid.New(ulid.Timestamp(time.Now()), rand.New())
		sender.ContextID = id.String()
		sender.Sender.Prompt, err = context.Set(sender.Message.Human, "")
	} else {
		ctx.ID = sender.ContextID
		d, ok := ctx.Find()
		if !ok {
			return nil, data.ErrContextNotFound
		}
		d = append(d, sender.Message)
		sender.Sender.Prompt, err = ctx.Build(d)
	}
	if err != nil {
		return ctx, err
	}
	sender.With(ah)
	return sender.Complete(ctx)
}

// Basic check
func (ah *Client) check(sender *resp.Sender) (err error) {
	if sender.Model == "" {
		sender.Model = ah.cfg.DefaultModel
	}
	if len(sender.StopSequences) == 0 {
		sender.StopSequences = data.StopSequences
	}
	if sender.MaxToken < 400 {
		sender.MaxToken = 400
	}
	return nil
}

func (c *Client) setHeaderWithURI(req *fasthttp.Request) {
	for k, v := range c.header {
		req.Header.Set(k, v)
	}
	req.SetRequestURI(data.API)
	req.Header.SetMethod("POST")
}

func (c *Client) do(req *fasthttp.Request, res *fasthttp.Response) error {
	return c.client.Do(req, res)
}
