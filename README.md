# anthropic-sdk-go

<p align="center">
        <a href="https://godoc.org/github.com/3JoB/anthropic-sdk-go"><img src="https://pkg.go.dev/badge/github.com/3JoB/anthropic-sdk-go.svg" alt="Go Reference"></a>
        <a href="https://github.com/3JoB/anthropic-sdk-go/blob/master/LICENSE"><img src="https://img.shields.io/github/license/3JoB/anthropic-sdk-go?style=flat-square" alt="MIT"></a>
        <a href="#"><img src="https://img.shields.io/github/go-mod/go-version/3JoB/anthropic-sdk-go?label=Go%20Version&style=flat-square" alt="Go Version"></a>
        <a href="https://github.com/3JoB/anthropic-sdk-go/release"><img src="https://img.shields.io/github/v/release/3JoB/anthropic-sdk-go?label=Release%20Version&style=flat-square" alt="GitHub release (latest by date)"></a>
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
- Prompt automatically handles
- Concise and easy-to-use API
- Support context memory (automatically handle context)

Claude Docs: [https://console.anthropic.com/docs](https://console.anthropic.com/docs)

<br><br>

## Start
Usage:
```sh
$ go get github.com/3JoB/anthropic-sdk-go@v1.2.2
```

<br>
Example usage:

```go
package main

import (
	"fmt"

	"github.com/3JoB/anthropic-sdk-go"
	"github.com/3JoB/anthropic-sdk-go/data"
)

func main() {
	c, err := anthropic.New("api keys","modules")
	if err != nil {
		panic(err)
	}

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
}
```

Return:
```json
{"detail":null,"completion":"Hello world! \nfmt.Println(\"Hello world!\")\n\nDone.","stop_reason":"stop_sequence","stop":"\n\nHuman:","log_id":"nop","exception":"","model":"claude-instant-v1","truncated":false}
```

<br>

Context Example:
```go
package main

import (
	"fmt"

	"github.com/3JoB/anthropic-sdk-go"
	"github.com/3JoB/anthropic-sdk-go/data"
)

func main() {
	c, err := anthropic.New("api keys","modules")
	if err != nil {
		panic(err)
	}

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
```

Return:
```json
{"detail":null,"completion":"Hello world! \nfmt.Println(\"Hello world!\")\n\nDone.","stop_reason":"stop_sequence","stop":"\n\nHuman:","log_id":"nop","exception":"","model":"claude-instant-v1","truncated":false}
{"detail":null,"completion":"1.14.4 ","stop_reason":"stop_sequence","stop":"\n\nHuman:","log_id":"nop","exception":"","model":"claude-instant-v1","truncated":false}
```

### Delete the context in an ID
```golang
c, err := anthropic.New("api keys","modules")
if err != nil {
	panic(err)
}

d, err := c.Send(&anthropic.Opts{
	Context: data.MessageModule{
		Human: "Do you know Golang, please answer me in the shortest possible way.",
	},
	Sender: anthropic.Sender{MaxToken: 1200},
})

if err != nil {
	panic(err)
}

d.Delete()
```

<br>

# Other
This project only guarantees basic usability, if you need new features or improvements, please create a `Pull Requests`

<br>

# License
This software is distributed under MIT license.