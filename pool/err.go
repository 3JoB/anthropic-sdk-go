package pool

import "errors"

var (
	ErrSliceNoCmp = errors.New("Slice cannot use compression!")
	ErrUnavaCmpAlg = errors.New("Unavailable compression algorithms.")
	ErrDisableSwitchCmp = errors.New("Disable switching of compression algorithm in a pool that has completed compression initialization.")
)