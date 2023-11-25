package resp

import (
	"github.com/3JoB/unsafeConvert"
	"github.com/sugawarayuuta/sonnet"
)

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
	d, _ := sonnet.Marshal(&resp)
	return unsafeConvert.StringPointer(d)
}
