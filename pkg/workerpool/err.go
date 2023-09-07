package workerpool

import "github.com/3JoB/anthropic-sdk-go/v2/internel/errors"

var (
	ErrPIDNotEmpty  = errors.New("pool_id cannot be empty.")
	ErrPIDHasLocked = errors.New("pool_id has been locked by other pool.")
	ErrNoPoolToPID  = errors.New("No pool corresponding to the pool_id.")
)
