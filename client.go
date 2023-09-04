package anthropic

import (
	"time"

	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/ulid"
	"github.com/cornelk/hashmap"
	"github.com/valyala/fasthttp"
	"pgregory.net/rand"

	"github.com/3JoB/anthropic-sdk-go/v2/context"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

type Client struct {
	client  *fasthttp.Client
	timeout time.Duration
	cfg     *Config                      // Config
	header  *hashmap.Map[string, string] // http header
}

type Config struct {
	Key          string // API Keys
	DefaultModel string // Choose the default AI model
	UseCache     bool   // Enable Prompt build cache
}

// is minute
func (ah *Client) SetTimeOut(times int) {
	if times == 0 {
		return
	}
	ah.timeout = time.Duration(times) * time.Minute
}

// Send data to the API endpoint. Before sending out, the data will be processed into a form that the API can recognize.
func (ah *Client) Send(senderOpts *Opts) (*context.Context, error) {
	var err error
	if err = ah.check(&senderOpts.Sender); err != nil {
		return nil, err
	}
	if (senderOpts.Message == data.MessageModule{}) {
		return nil, data.ErrContextIsNil
	}
	ctx := senderOpts.new()
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
	senderOpts.With(ah)
	return senderOpts.Complete(ctx)
}

func (ah *Client) check(sender *resp.Sender) (err error) {
	if sender.Model == "" {
		sender.Model = ah.cfg.DefaultModel
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
	if c.cfg.Key == "" {
		return data.ErrApiKeyEmpty
	}
	c.header.Set("Accept", "application/json")
	c.header.Set("Content-Type", "application/json")
	c.header.Set("Client", litefmt.Sprint("anthropic-sdk-go/", SDKVersion))
	c.header.Set("anthropic-version", "2023-06-01")
	c.header.Set("x-api-key", c.cfg.Key)
	c.header.Set("User-Agent", UserAgent)
	return nil
}
