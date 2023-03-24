package main

import (
	"fmt"

	"github.com/3JoB/anthropic-sdk-go"
)

func main() {
	c, err := anthropic.NewClient(&anthropic.AnthropicClient{
		Key: "",
	})
	if err != nil {
		panic(err)
	}
	d, err := c.Send(&anthropic.Sender{
		Prompt: "what's your name?",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
}
