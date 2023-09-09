package resp

import (
	"github.com/3JoB/unsafeConvert"
	"github.com/bytedance/sonic"
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

type Response struct {
	Completion string `json:"completion"`          // The resulting completion up to and excluding the stop sequences.
	StopReason string `json:"stop_reason"`         // The reason we stopped sampling, either if we reached one of your provided , or if we exceeded `.stop_sequencestop_sequencesmax_tokensmax_tokens_to_sample`
	Stop       string `json:"stop"`                // If the is , this contains the actual stop sequence (of the list passed-in) that was `seenstop_reasonstop_sequencestop_sequences`
	LogID      string `json:"log_id"`              // The ID of the log that generated the response
	Exception  string `json:"exception,omitempty"` // exception
	Model      string `json:"model"`               // Model
	Truncated  bool   `json:"truncated"`           // truncated
}

func (resp Response) String() string {
	d, _ := sonic.Marshal(&resp)
	return unsafeConvert.StringSlice(d)
}

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"-"` // HTTP status code, please do not parse
}

type R struct {
	Error *ErrorResponse `json:"error"`
}

func Error(code int, v []byte) (*ErrorResponse, error) {
	var e = R{
		Error: &ErrorResponse{},
	}

	err := sonic.Unmarshal(v, &e)
	e.Error.Code = code
	return e.Error, err
}

/*
Example:

	func E() error {
		return &ErrorResponse{}
	}
*/
func (e *ErrorResponse) Error() string {
	return e.Message
}

// Return an error object
func (e *ErrorResponse) Err() error {
	return e
}

// Return HTTP status code
func (e *ErrorResponse) StatusCode() int {
	return e.Code
}

// Check if the HTTP status code matches the entered status code
func (e *ErrorResponse) IsStatusCode(code int) bool {
	return code == e.Code
}

// Return a String object
//
// if you need to call it multiple times, please assign or cache it directly,
// because the structure will call json to decode itself
func (e *ErrorResponse) String() string {
	d, _ := sonic.Marshal(e)
	return unsafeConvert.StringSlice(d)
}
