package hashpool

import "github.com/3JoB/anthropic-sdk-go/v2/internel/errors"

var (
	ErrSliceNoCmp       = errors.New("Slice cannot use compression!")
	ErrUnavaCmpAlg      = errors.New("Unavailable compression algorithms.")
	ErrDisableSwitchCmp = errors.New("Disable switching of compression algorithm in a pool that has completed compression initialization.")
)
