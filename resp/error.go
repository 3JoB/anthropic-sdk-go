package resp

import (
	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/unsafeConvert"
	"github.com/sugawarayuuta/sonnet"
)

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"-"` // HTTP status code, please do not parse
}

type R struct {
	Error *ErrorResponse `json:"error"`
}

func Error(code int, v []byte) (*ErrorResponse, error) {
	var e = R{
		Error: &ErrorResponse{},
	}
	if err := sonnet.Unmarshal(v, &e); err != nil {
		return nil, err
	}
	e.Error.Code = code
	return e.Error, nil
}

/*
Example:

	func E() error {
		return &ErrorResponse{}
	}
*/
func (e *ErrorResponse) Error() string {
	return litefmt.Sprint("anthropic: ", e.Message, " (", unsafeConvert.IntToString(e.Code), ")")
}

// Return an error object
func (e *ErrorResponse) Err() error {
	return e
}

// Return HTTP status code
func (e *ErrorResponse) StatusCode() int {
	return e.Code
}

// Check if the HTTP status code matches the entered status code
func (e *ErrorResponse) IsStatusCode(code int) bool {
	return code == e.Code
}
