package anthropic

import (
	"time"

	"github.com/3JoB/resty-ilo"
	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/ulid"
	"pgregory.net/rand"

	"github.com/3JoB/anthropic-sdk-go/v2/context"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

type Client struct {
	Key          string        // API Keys
	DefaultModel string        // Choose the default AI model
	UseCache     bool          // Enable Prompt build cache
	client       *resty.Client // http client
}

// is minute
func (ah *Client) SetTimeOut(times int) {
	if times == 0 {
		return
	}
	ah.client = ah.client.SetTimeout(time.Duration(times) * time.Minute)
}

// Send data to the API endpoint. Before sending out, the data will be processed into a form that the API can recognize.
func (ah *Client) Send(senderOpts *Opts) (*context.Context, error) {
	var err error
	if err = ah.check(&senderOpts.Sender); err != nil {
		return nil, err
	}
	if (senderOpts.Message == data.MessageModule{}) {
		return nil, data.ErrContextNil
	}
	ctx := senderOpts.newCtx()
	if senderOpts.ContextID == "" {
		id, _ := ulid.New(ulid.Timestamp(time.Now()), rand.New())
		senderOpts.ContextID = id.String()
		senderOpts.Sender.Prompt, err = context.Set(senderOpts.Message.Human, "")
	} else {
		ctx.ID = senderOpts.ContextID
		d, ok := ctx.Find()
		if !ok {
			return nil, data.ErrContextNotFound
		}
		d = append(d, senderOpts.Message)
		senderOpts.Sender.Prompt, err = ctx.Build(d)
	}
	if err != nil {
		return ctx, err
	}
	return senderOpts.Complete(ctx, ah.client)
}

func (ah *Client) check(sender *resp.Sender) (err error) {
	if sender.Model == "" {
		sender.Model = ah.DefaultModel
	}
	if len(sender.StopSequences) == 0 {
		sender.StopSequences = StopSequences
	}
	if sender.MaxToken < 400 {
		sender.MaxToken = 400
	}
	return nil
}

func (c *Client) headers() error {
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
