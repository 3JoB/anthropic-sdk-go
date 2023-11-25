package anthropic

import (
	"time"

	"github.com/3JoB/ulid"
	"github.com/valyala/fasthttp"
	"pgregory.net/rand"

	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/pkg/pool"
)

type Client struct {
	client  *fasthttp.Client
	key     string
	model   string
	pool    *pool.Pool
	header  map[string]string
	timeout time.Duration
}

// Set the response timeout in minutes.
func (c *Client) SetTimeOut(times int) {
	if times != 0 {
		c.timeout = time.Duration(times) * time.Minute
	}
}

// Send data to the API endpoint. Before sending out,
// the data will be processed into a form that the API can recognize.
func (c *Client) Send(sender *Sender) (*pool.Session, error) {
	var err error
	if (sender.Message == data.MessageModule{}) {
		return nil, data.ErrSessionIsNil
	}
	c.check(sender)
	ss := sender.newSession()
	if sender.SessionID == "" {
		id, _ := ulid.New(ulid.Timestamp(time.Now()), rand.New())
		sender.SessionID = id.String()
		ss.ID = sender.SessionID
		err = sender.Sender.Set(sender.Message)
	} else {
		ss.ID = sender.SessionID
		p, ok := c.pool.Get(ss.ID)
		if !ok {
			return nil, data.ErrSessionNotFound
		}
		err = sender.Sender.Build(p, sender.Message)
	}
	if err != nil {
		return ss, err
	}
	if err := sender.Complete(c, ss); err != nil {
		return nil, err
	}
	_ = c.pool.Append(ss.ID, ss.Response.Completion)
	return ss, nil
}

// Basic check
func (c *Client) check(s *Sender) {
	if s.Sender.Model == "" {
		s.Sender.Model = c.model
	}
	if len(s.Sender.StopSequences) == 0 {
		s.Sender.StopSequences = data.StopSequences
	}
	if s.Sender.MaxToken < 400 {
		s.Sender.MaxToken = 400
	}
}

func (c *Client) do(req *fasthttp.Request, res *fasthttp.Response) error {
	return c.client.Do(req, res)
}
