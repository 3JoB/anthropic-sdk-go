package main

import (
	"fmt"

	"github.com/3JoB/anthropic-sdk-go"
	"github.com/3JoB/anthropic-sdk-go/data"
)

func main() {
	c, err := anthropic.New("your keys","")
	if err != nil {
		panic(err)
	}
	/*d, err := c.Send(&anthropic.Sender{
		Prompt:   "Do you know Golang, please answer me in the shortest possible way.",
		MaxToken: 1200,
	})*/
	d, err := c.Send(&anthropic.Opts{
		Context: data.MessageModule{
			Human: "Do you know Golang, please answer me in the shortest possible way.",
		},
		Sender: anthropic.Sender{MaxToken: 1200},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(d.Response.String())

	ds, err := c.Send(&anthropic.Opts{
		Context: data.MessageModule{
            Human: "What is its current version number?",
        },
		ContextID: d.ID,
        Sender: anthropic.Sender{MaxToken: 1200},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(ds.Response.String())
}
