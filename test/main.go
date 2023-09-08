package main

import (
	"fmt"

	"github.com/3JoB/anthropic-sdk-go/v2"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/internel/compress"
	"github.com/3JoB/anthropic-sdk-go/v2/pkg/hashpool"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

var c_data = `Seasons come and seasons go
The world is always changing
New life springs, old life fades
A cycle ever ranging

The sun will rise, the moon will set
The stars blink in the sky
Clouds will form, rain will fall
And time will pass us by

But some things remain through it all
The love we have inside
Bindings of family, bonds of friends
With us they will abide

So cherish every moment
Of laughter, joy and mirth
Find beauty in each season
And wonder in the earth`

func main() {
	p := hashpool.NewPoolWithSlice()
	if err := p.UseComress("br"); err != nil {
		panic(err)
	}
	b := []data.MessageModule{
		{
			Human:     "123",
			Assistant: "456",
		},
	}
	p.Set("c_data", b)
	d, ok := p.Get("c_data")
	if !ok {
		panic("get failed")
	}
	fmt.Println(d)
}

func cmain() {
	data := []byte(c_data)
	fmt.Println("Size: ", len(data))
	cp := compress.NewSnappy()
	en_buf, err := cp.Encode(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encode Size: ", en_buf.Len())
	fmt.Println("Encode Data: ", string(en_buf.Bytes()))
	fmt.Println("Decode Data: ")
	cp.Decode(en_buf)
	fmt.Println(string(en_buf.Bytes()))
}

func xmain() {
	// fuck i accidentally leaked my keys and it's now disabled by me.
	c, err := anthropic.New(&anthropic.Config{Key: "your keys", DefaultModel: data.ModelFullInstant})
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
