package main

import (
	"fmt"

	"github.com/3JoB/anthropic-sdk-go/v2"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

func main() {
	// fuck i accidentally leaked my keys and it's now disabled by me.
	c, err := anthropic.New(&anthropic.Config{Key: "your keys", DefaultModel: anthropic.Model.Full.Instant1})
	if err != nil {
		panic(err)
	}
	/*d, err := c.Send(&anthropic.Sender{
		Prompt:   "Do you know Golang, please answer me in the shortest possible way.",
		MaxToken: 1200,
	})*/
	d, err := c.Send(&anthropic.Opts{
		Message: data.MessageModule{
			Human: "Do you know Golang, please answer me in the shortest possible way.",
		},
		Sender: resp.Sender{MaxToken: 1200},
	})
	if err != nil {
		fmt.Println(d.ErrorResp.Message)
		panic(err)
	}
	fmt.Println(d.Response.String())

	ds, err := c.Send(&anthropic.Opts{
		Message: data.MessageModule{
			Human: "What is its current version number?",
		},
		ContextID: d.ID,
		Sender:    resp.Sender{MaxToken: 1200},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(ds.Response.String())

	// Set UserID
	dsr, err := c.Send(&anthropic.Opts{
		Message: data.MessageModule{
			Human: "What is its current version number?",
		},
		ContextID: d.ID,
		Sender: resp.Sender{
			MaxToken: 1200,
			MetaData: resp.MetaData{
				UserID: "rand id",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(dsr.Response.String())
}
