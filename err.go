package anthropic

import (
	"errors"
	"fmt"
)

type Err struct {
	Op  string
	Err string
	e error
}

func (e *Err) Error() string {
	e.e = errors.New(e.Err)
	return fmt.Sprintf("%v: %v", e.Op, e.e.Error())
}

func (e *Err) Unwrap() error {
	e.e = errors.New(e.Err)
    return e.e
}