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
	e.e = errors.New(e.Err)
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
