package main

import (
	"fmt"

	"github.com/3JoB/anthropic-sdk-go"
	"github.com/3JoB/anthropic-sdk-go/data"
)

func main() {
	// fuck i accidentally leaked my keys and it's now disabled by me.
	c, err := anthropic.New("keys ", "")
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
		Sender: anthropic.Sender{MaxToken: 1200},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(d.Response.String())

	ds, err := c.Send(&anthropic.Opts{
		Message: data.MessageModule{
			Human: "What is its current version number?",
		},
		ContextID: d.ID,
		Sender:    anthropic.Sender{MaxToken: 1200},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(ds.Response.String())
}
