package anthropic

import (
	"fmt"
	"os"

	"github.com/3JoB/ulib/err"
)

func Exit(e error, code ...int) {
	fmt.Println(e.Error())
	if len(code) != 0 {
		os.Exit(code[0])
	} else {
		os.Exit(-99)
	}
}

var (
	ErrApiKeyEmpty      error = &err.Err{Op: "config", Err: "APIKey cannot be empty!"}
	ErrConfigEmpty      error = &err.Err{Op: "config", Err: "Configuration cannot be empty!"}
	ErrSenderNil        error = &err.Err{Op: "sender", Err: "Sender cannot be nil!"}
	ErrPromptHumanEmpty error = &err.Err{Op: "prompt", Err: "The value of human cannot be empty!"}
	ErrPromptCtxEmpty   error = &err.Err{Op: "prompt", Err: "The value of context cannot be empty!"}
	ErrPromptEmpty      error = &err.Err{Op: "send", Err: "The value of prompt cannot be empty!"}
)
