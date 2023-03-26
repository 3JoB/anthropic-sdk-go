package main

import (
	"fmt"

	"github.com/3JoB/anthropic-sdk-go"
)

func main() {
	c, err := anthropic.NewClient(&anthropic.AnthropicClient{
		Key: "your keys",
	})
	if err != nil {
		panic(err)
	}
	/*d, err := c.Send(&anthropic.Sender{
		Prompt:   "Do you know Golang, please answer me in the shortest possible way.",
		MaxToken: 1200,
	})*/
	d, err := c.Send(&anthropic.Opts{
		Context: []anthropic.MessageModule{
			{
				Human: "Do you know Golang, please answer me in the shortest possible way.",
			},
		},
		Sender: anthropic.Sender{MaxToken: 1200},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(d.Response.String())

	ds, err := c.SendWithContext(&anthropic.Sender{
		Prompt:   "What is its current version number?",
		MaxToken: 1200,
	},
		d.CtxData)
	if err != nil {
		panic(err)
	}
	fmt.Println(ds.Response.String())
}
