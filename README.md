<h1 align="center">Anthropic-SDK-Go</h1>

<p align="center">
        <a href="https://pkg.go.dev/github.com/3JoB/anthropic-sdk-go/v2"><img src="https://pkg.go.dev/badge/github.com/3JoB/anthropic-sdk-go/v2.svg" alt="Go Reference"></a>
        <a href="https://github.com/3JoB/anthropic-sdk-go/blob/master/LICENSE"><img src="https://img.shields.io/github/license/3JoB/anthropic-sdk-go?style=flat-square" alt="MIT"></a>
        <a href="#"><img src="https://img.shields.io/github/go-mod/go-version/3JoB/anthropic-sdk-go?label=Go%20Version&style=flat-square" alt="Go Version"></a>
        <a href="https://github.com/3JoB/anthropic-sdk-go/releases"><img src="https://img.shields.io/github/v/release/3JoB/anthropic-sdk-go?label=Release%20Version&style=flat-square" alt="GitHub release (latest by date)"></a>
    </p>
    <p align="center">
        <a href="https://github.com/3JoB/anthropic-sdk-go/issues"><img src="https://img.shields.io/github/issues/3JoB/anthropic-sdk-go?label=Issues&style=flat-square" alt="GitHub Issues"></a>
        <a href="https://github.com/3JoB/anthropic-sdk-go/stargazers"><img src="https://img.shields.io/github/stars/3JoB/anthropic-sdk-go?label=Stars&style=flat-square" alt="GitHub Repo stars"></a>
        <a href="#"><img src="https://img.shields.io/github/repo-size/3JoB/anthropic-sdk-go?style=flat-square" alt="GitHub repo size"></a>
        <a href="#"><img src="https://img.shields.io/github/commit-activity/m/3JoB/anthropic-sdk-go?style=flat-square" alt="GitHub commit activity"></a>
    </p>
<p align="center">Golang SDK for AnthRopic Claude AI</p>

<br>

## Features
- Contextual sequential memory 
- Prompt automatically handles / Contextual automated processing
- Concise and easy-to-use API
- Fast data processing


Claude Docs: [https://console.anthropic.com/docs](https://console.anthropic.com/docs)

<br><br>

## Start

Usage:
```sh
$ go get github.com/3JoB/anthropic-sdk-go/v2@v2.1.0
```

<br>
Example usage:

```go
package main

import (
	"fmt"

	"github.com/3JoB/anthropic-sdk-go/v2"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

func main() {
	c, err := anthropic.New(&anthropic.Config{Key: "your keys", DefaultModel: data.ModelFullInstant})
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

	fmt.Println(d.Response.String())
}
```

Return:
```json
{"detail":null,"completion":"Hello world! \nfmt.Println(\"Hello world!\")\n\nDone.","stop_reason":"stop_sequence","stop":"\n\nHuman:","log_id":"nop","exception":"","model":"claude-instant-v1.2","truncated":false}
```

<br>

Context Example:
```go
package main

import (
	"fmt"

	"github.com/3JoB/anthropic-sdk-go/v2"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
)

func main() {
	c, err := anthropic.New(&anthropic.Config{Key: "your keys", DefaultModel: data.ModelFullInstant})
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

	fmt.Println(d.Response.String())

	ds, err := c.Send(&anthropic.Sender{
		Message: data.MessageModule{
			Human: "What is its current version number?",
		},
		SessionID: d.ID,
		Sender: &resp.Sender{MaxToken: 1200},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(ds.Response.String())
}
```

Return:
```json
{"detail":null,"completion":"Hello world! \nfmt.Println(\"Hello world!\")\n\nDone.","stop_reason":"stop_sequence","stop":"\n\nHuman:","log_id":"nop","exception":"","model":"claude-instant-v1","truncated":false}
{"detail":null,"completion":"1.14.4 ","stop_reason":"stop_sequence","stop":"\n\nHuman:","log_id":"nop","exception":"","model":"claude-instant-v1.2","truncated":false}
```

### Delete the session in an ID
```golang
c, err := anthropic.New(&anthropic.Config{Key: "your keys", DefaultModel: data.Model_FullInstant_1_0})
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

c.CloseSession(d)
```

<br>


# Contribute
Move => [CONTRIBUTING](/CONTRIBUTING.md)


# Contact
Organize EMAIL: `admin#zxda.top` [# => @]

<br>

# License
This software is distributed under MIT license.
