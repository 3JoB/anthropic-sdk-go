package resp

import (
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/ulib/litefmt"
)

type Sender struct {
	Prompt        string   `json:"prompt"`                   // (required) The prompt you want Claude to complete. For proper response generation you will most likely want to format your prompt as follows:See [our comments on prompts](https://console.anthropic.com/docs/prompt-design#what-is-a-prompt) for more context.
	Model         string   `json:"model"`                    // (required) As we improve Claude, we develop new versions of it that you can query. This controls which version of Claude answers your request
	StopSequences []string `json:"stop_sequences,omitempty"` // (optional) A list of strings upon which to stop generating. You probably want , as that's the cue for the next turn in the dialog agent. Our client libraries provide a constant for this value (see examples below["\n\nHuman:"])
	Stream        bool     `json:"stream"`                   // (optional) Amount of randomness injected into the response. Ranges from 0 to 1. Use temp closer to 0 for analytical / multiple choice, and temp closer to 1 for creative and generative tasks.
	Temperature   float64  `json:"temperature,omitempty"`    // (required) Amount of randomness injected into the response. Defaults to 1. Ranges from 0 to 1. Use temp closer to 0 for analytical / multiple choice, and closer to 1 for creative and generative tasks.
	TopK          float64  `json:"top_k,omitempty"`          // (optional) Only sample from the top K options for each subsequent token. Used to remove "long tail" low probability responses. Defaults to -1, which disables it.
	TopP          uint     `json:"top_p,omitempty"`          // (optional) Does nucleus sampling, in which we compute the cumulative distribution over all the options for each subsequent token in decreasing probability order and cut it off once it reaches a particular probability specified by . Defaults to -1, which disables it. Note that you should either alter or , but not both.`top_ptemperaturetop_p``
	MaxToken      uint     `json:"max_tokens_to_sample"`     // (required) A maximum number of tokens to generate before stopping.
	MetaData      MetaData `json:"metadata,omitempty"`       // (optional) An object describing metadata about the request.
}

type MetaData struct {
	UserID string `json:"user_id,omitempty"` // (optional) A uuid, hash value, or other external identifier for the user who is associated with the request. Anthropic may use this id to help detect abuse. Do not include any identifying information such as name, email address, or phone number.
}

// A uuid, hash value, or other external identifier for the user who is associated with the request. Anthropic may use this id to help detect abuse. Do not include any identifying information such as name, email address, or phone number.
func (s *Sender) SetUserID(userID string) {
	s.MetaData = MetaData{UserID: userID}
}

func (s *Sender) Set(d *data.MessageModule) error {
	if d.Human == "" {
		return data.ErrPromptHumanEmpty
	}
	if d.Assistant == "" {
		s.Prompt = litefmt.Sprint("\n\nHuman: ", d.Human, "\n\nAssistant:")
	} else {
		s.Prompt = litefmt.Sprint(d.Human, d.Assistant)
	}
	return nil
}

func (s *Sender) Build(next string,  d *data.MessageModule) error {
	if d.Human == "" {
		return data.ErrPromptHumanEmpty
	}
	if d.Assistant == "" {
		s.Prompt = litefmt.Sprint(next, "\n\nHuman: ", d.Human, "\n\nAssistant:")
	} else {
		s.Prompt = litefmt.Sprint(next, "\n\nHuman: ", d.Human, "\n\nAssistant:", d.Assistant)
	}
	return nil
}
