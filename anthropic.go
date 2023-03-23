package anthropic

import (
	"fmt"

	"github.com/3JoB/ulib/maps"
	"github.com/3JoB/ulib/net/ua"
)

func SetPrompt(human string) string {
	return fmt.Sprintf(`\n\nHuman: %v\n\nAssistant:`, human)
}

func SetHeaders(api string) {
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
