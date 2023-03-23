package anthropic

import (
	"errors"
	"fmt"
	"os"
)

type Err struct {
	Op  string
	Err string
	e   error
}

func (e *Err) Error() string {
	e.e = fmt.Errorf(e.Err)
	return fmt.Sprintf("%v: %v", e.Op, e.e.Error())
}

func (e *Err) Unwrap() error {
	e.e = errors.New(e.Err)
	return e.e
}

func Exit(e error, code ...int) {
	fmt.Println(e.Error())
	if len(code) != 0 {
		os.Exit(code[0])
	} else {
		os.Exit(-99)
	}
}

var (
	ErrApiKeyEmpty error = &Err{Op: "config", Err: "APIKey cannot be empty!"}
	ErrConfigEmpty error = &Err{Op: "config", Err: "Configuration cannot be empty!"}
	ErrSenderNil   error = &Err{Op: "sender", Err: "Sender cannot be nil!"}
	ErrHumanEmpty  error = &Err{Op: "prompt", Err: "The value of human cannot be empty!"}
	ErrPromptEmpty error = &Err{Op: "send", Err: "The value of prompt cannot be empty!"}
)
