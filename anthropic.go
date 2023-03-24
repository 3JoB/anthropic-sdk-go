package anthropic

import (
	"fmt"

	"github.com/3JoB/ulib/net/ua"
	"github.com/go-resty/resty/v2"
)

// Create a new Client object.
func NewClient(conf *AnthropicClient) (*AnthropicClient, error) {
	if conf == nil {
		return nil, ErrConfigEmpty
	}
	if err := setHeaders(conf.Key); err != nil {
		return nil, err
	}
	conf.client = resty.New().SetBaseURL(API).SetHeaders(Headers)
	if conf.DefaultModel == "" {
		conf.DefaultModel = ModelClaudeInstantV1
	}
	return conf, nil
}

func (ah *AnthropicClient) c(sender *Sender) (err error) {
	if sender == nil {
		return ErrSenderNil
	}
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
func (ah *AnthropicClient) Send(sender *Sender) (ctx *Context, err error) {
	if err := ah.c(sender); err != nil {
		return nil, err
	}
	sender.Prompt, err = setPrompt(sender.Prompt, "")
	if err != nil {
		return nil, err
	}
	return sender.Complete(ah.client)
}

// Send data to the API endpoint. Before sending out, the data will be processed into a form that the API can recognize.
//
// This method will be used to handle context requests.
//
// The context parameter comes from *Context.CtxData, please do not modify or process it by yourself, the context will be automatically processed when the previous request is executed.
func (ah *AnthropicClient) SendWithContext(sender *Sender, context string) (ctx *Context, err error) {
	if err := ah.c(sender); err != nil {
		return nil, err
	}
	sender.Prompt, err = addPrompt(context, sender.Prompt)
	if err != nil {
		return nil, err
	}
	return sender.Complete(ah.client)
}

func setPrompt(human, assistant string) (string, error) {
	if human == "" {
		return "", ErrPromptHumanEmpty
	}
	if assistant == "" {
		return fmt.Sprintf("\n\nHuman: %v\n\nAssistant:", human), nil
	}
	return fmt.Sprintf("%v%v", human, assistant), nil
}

func addPrompt(context, human string) (string, error) {
	if human == "" {
		return "", ErrPromptHumanEmpty
	}
	if context == "" {
		return "", ErrPromptCtxEmpty
	}
	return fmt.Sprintf("%v\n\nHuman: %v\n\nAssistant:", context, human), nil
}

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
