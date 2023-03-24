package anthropic

import (
	"fmt"

	"github.com/3JoB/ulib/net/ua"
)

func NewClient(conf *AnthropicClient) (*AnthropicClient, error) {
	if conf == nil {
		return nil, ErrConfigEmpty
	}
	if err := setHeaders(conf.Key); err != nil {
		return nil, err
	}
	if conf.DefaultModel == "" {
		conf.DefaultModel = ModelClaudeInstantV1
	}
	return conf, nil
}

func (ah *AnthropicClient) Send(sender *Sender) (data *Response, err error) {
	if sender == nil {
		return nil, ErrSenderNil
	}
	if sender.Prompt == "" {
		return nil, ErrPromptEmpty
	}
	sender.Prompt, err = setPrompt(sender.Prompt, "")
	if err != nil {
		return nil, err
	}
	sender.Stream = false
	if sender.Model == "" {
		sender.Model = ah.DefaultModel
	}
	if len(sender.StopSequences) == 0 {
		sender.StopSequences = StopSequences
	}
	if sender.MaxToken < 1 {
		sender.MaxToken = 400
	}
	sender.Complete()
	return nil, nil
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
