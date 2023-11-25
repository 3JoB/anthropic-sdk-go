package main

import (
	// "fmt"

	"github.com/3JoB/anthropic-sdk-go/v2"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

func main() {
	c, err := anthropic.New(&anthropic.Config{Key: "your token", DefaultModel: data.ModelFullInstant})
	if err != nil {
		panic(err)
	}
	d, err := c.Send(&anthropic.Sender{
		Message: data.MessageModule{
			Human: "Do you know Golang, please answer me in the shortest possible way.",
		},
		Sender: &resp.Sender{MaxToken: 1200},
	})
	if err != nil {
		panic(err)
	}
	// fmt.Println(d.Response.String())

	_, err = c.Send(&anthropic.Sender{
		Message: data.MessageModule{
			Human: "What is its current version number?",
		},
		SessionID: d.ID,
		Sender:    &resp.Sender{MaxToken: 1200},
	})
	if err != nil {
		panic(err)
	}
	// fmt.Println(ds.Response.String())

	// Set UserID
	dsr, err := c.Send(&anthropic.Sender{
		Message: data.MessageModule{
			Human: "What is its current version number?",
		},
		SessionID: d.ID,
		Sender: &resp.Sender{
			MaxToken: 1200,
			MetaData: resp.MetaData{
				UserID: "rand id",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	// fmt.Println(dsr.Response.String())
	c.CloseSession(dsr)
}
