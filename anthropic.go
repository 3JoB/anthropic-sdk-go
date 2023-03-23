package anthropic

import (
	"fmt"

	"github.com/3JoB/ulib/maps"
	"github.com/3JoB/ulib/net/ua"
)

func New(conf *Anthropic) *Anthropic {
	if conf == nil {
        panic(ErrConfigEmpty)
	}
	setHeaders(conf.Key)
	if conf.DefaultModel == "" {

	}
	return conf
}

func setPrompt(human string) (string, error) {
	if human == "" {
		return "", ErrHumanEmpty
	}
	return fmt.Sprintf(`\n\nHuman: %v\n\nAssistant:`, human), nil
}

func setHeaders(api string) {
	if api == "" {
		panic(ErrApiKeyEmpty)
	}
	Headers = maps.New(Headers)
	Headers = map[string]string{
		"Accept":        "application/json",
		"Client":        fmt.Sprintf("anthropic-sdk-go/%v", SDKVersion),
		"X-SDK-Version": SDKVersion,
		"X-SDK-Repo":    "https://github.com/3JoB/anthropic-sdk-go",
		"X-API-Key":     api,
		"User-Agent":    ua.ULIBDefault,
	}
}
