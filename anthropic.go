package anthropic

import (
	"fmt"
	"time"

	"github.com/3JoB/ulib/net/ua"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

type AnthropicClient struct {
	Key          string        // API Keys
	DefaultModel string        // Choose the default AI model
	client       *resty.Client // http client
}

// Create a new Client object.
func New(key, defaultModel string) (*AnthropicClient, error) {
	if err := setHeaders(key); err != nil {
		return nil, err
	}
	conf := &AnthropicClient{
		client: resty.New().SetBaseURL(API).SetHeaders(Headers),
	}
	if defaultModel == "" {
		conf.DefaultModel = ModelClaudeV12
	}
	return conf, nil
}

// is minute
func (ah *AnthropicClient) SetTimeOut(times int) {
	if times == 0 {
		return
	}
	ah.client = ah.client.SetTimeout(time.Duration(times) * time.Minute)
}

func (ah *AnthropicClient) check(sender *Sender) (err error) {
	if sender.Prompt == "" {
		return ErrPromptEmpty
	}
	if sender.Model == "" {
		sender.Model = ah.DefaultModel
	}
	if len(sender.StopSequences) == 0 {
		sender.StopSequences = StopSequences
	}
	if sender.MaxToken < 1 {
		sender.MaxToken = 400
	}
	return nil
}

// Send data to the API endpoint. Before sending out, the data will be processed into a form that the API can recognize.
func (ah *AnthropicClient) Send(senderOpts *Opts) (ctx *Context, err error) {
	if err := ah.check(&senderOpts.Sender); err != nil {
		return nil, err
	}
	if (senderOpts.Context == MessageModule{}) {
		return nil, ErrContextNil
	}
	if senderOpts.ContextID == "" {
		senderOpts.ContextID = uuid.New().String()
		senderOpts.Sender.Prompt, err = buildPrompts(senderOpts.Context)
	} else {
		d, ok := ctx.Find()
		if !ok {
			return nil, ErrContextNotFound
		}
		senderOpts.Sender.Prompt, err = buildPrompts(d)
	}
	return senderOpts.Complete(ah.client)
}

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

func setHeaders(api string) error {
	if api == "" {
		return ErrApiKeyEmpty
	}
	Headers = map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"Client":       fmt.Sprintf("anthropic-sdk-go/%v", SDKVersion),
		"x-api-key":    api,
		"User-Agent":   ua.Chrome,
	}
	return nil
}
