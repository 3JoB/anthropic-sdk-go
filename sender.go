package anthropic

import (
	"errors"

	"github.com/3JoB/unsafeConvert"
	"github.com/sugawarayuuta/sonnet"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/pkg/pool"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

type Sender struct {
	Message   data.MessageModule // Chunked message structure
	SessionID string             // Session ID. If empty, a new session is automatically created. If not empty, an attempt is made to find an existing session.
	Sender    *resp.Sender
}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) newSession() *pool.Session {
	return &pool.Session{
		Response: resp.Response{},
		Human:    s.Message.Human,
	}
}

// Make a processed request to an API endpoint.
func (s *Sender) Complete(client *Client, session *pool.Session) error {
	// Get fasthttp object
	request, response := client.Acquire()
	defer client.Release(request, response)
	if err := sonnet.NewEncoder(request.BodyWriter()).Encode(&s.Sender); err != nil {
		return err
	}

	if err := client.do(request, response); err != nil {
		return err
	}

	session.ID = s.SessionID
	if err := sonnet.Unmarshal(response.Body(), &session.Response); err != nil {
		return err
	}
	session.RawData = response.Body()

	if response.StatusCode() != 200 {
		err, errp := resp.Error(response.StatusCode(), response.Body())
		if errp != nil {
			return err
		}
		if err != nil {
			session.ErrorResp = err
			return err
		}
		return errors.New(unsafeConvert.StringPointer(session.RawData))
	}
	s.Message.Assistant = session.Response.Completion

	return nil
}
