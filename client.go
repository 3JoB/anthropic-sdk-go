package anthropic

import (
	"pgregory.net/rand"
	"sync"
	"time"

	"github.com/3JoB/anthropic-sdk-go/data"
	"github.com/3JoB/anthropic-sdk-go/internal/prompt"
	"github.com/3JoB/resty-ilo"
	"github.com/3JoB/ulid"
)

type AnthropicClient struct {
	Key          string        // API Keys
	DefaultModel string        // Choose the default AI model
	client       *resty.Client // http client
}

// Create a new Client object.
func New(key, defaultModel string) (*AnthropicClient, error) {
	conf := &AnthropicClient{}
	if headers, err := initHeaders(key); err != nil {
		return nil, err
	} else {
		conf.client = resty.New().SetBaseURL(API).SetHeaders(headers)
	}
	if defaultModel == "" {
		conf.DefaultModel = ModelClaudeV13
	}
	if conf.TestBan() {
		return nil, data.ErrRegionBanned
	}
	return conf, nil
}

func NewPool(key, defaultModel string) sync.Pool {
	return sync.Pool{
		New: func() any {
			if client, err := New(key, defaultModel); err != nil {
				panic(err)
			} else {
				return client
			}
		},
	}
}

// is minute
func (ah *AnthropicClient) SetTimeOut(times int) {
	if times == 0 {
		return
	}
	ah.client = ah.client.SetTimeout(time.Duration(times) * time.Minute)
}

// Send data to the API endpoint. Before sending out, the data will be processed into a form that the API can recognize.
func (ah *AnthropicClient) Send(senderOpts *Opts) (*Context, error) {
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
		senderOpts.Sender.Prompt, err = prompt.Set(senderOpts.Message.Human, "")
	} else {
		ctx.ID = senderOpts.ContextID
		d, ok := ctx.Find()
		if !ok {
			return nil, data.ErrContextNotFound
		}
		d = append(d, senderOpts.Message)
		senderOpts.Sender.Prompt, err = prompt.Build(d)
	}
	if err != nil {
		return ctx, err
	}
	return senderOpts.Complete(ctx, ah.client)
}

func (ah *AnthropicClient) ResetContextPool() {
	RefreshContext()
}

func (ah *AnthropicClient) TestBan() bool {
	req := ah.client.R()
	req.RawRequest.Close = true
	req.RawRequest.Response.Close = true
	resp, err := req.Get("/")
	if err != nil {
		return true
	}
	defer resp.RawBody().Close()
	return resp.StatusCode() == 403
}

func (ah *AnthropicClient) check(sender *Sender) (err error) {
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