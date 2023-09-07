package errors

import (
	"runtime"
	"strconv"

	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/unsafeConvert"
)

type errorF struct {
	s string
}

func New(err string) *errorF {
	ptr, s, l, _ := runtime.Caller(1)
	return &errorF{s: litefmt.Sprint(s, ":", unsafeConvert.IntToString(l), " 0x", strconv.FormatUint(uint64(ptr), 16), " ", err)}
}

func (e *errorF) Error() string {
	return e.s
}
