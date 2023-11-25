package pool

import "github.com/3JoB/anthropic-sdk-go/v2/resp"

type Session struct {
	ID        string // Session ID
	Human     string
	RawData   []byte // Unprocessed raw json data returned by the API endpoint
	Response  resp.Response
	ErrorResp *resp.ErrorResponse
}
