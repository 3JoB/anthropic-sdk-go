package anthropic

import (
	"sync"

	fast "github.com/3JoB/fasthttp-client"
	"github.com/goccy/go-json"
)

var Pool sync.Pool = fast.NewClientPool()

func GetPool() *fast.Client {
	return Pool.Get().(*fast.Client)
}

func SetPool() {
	Pool.Put(GetPool().AddHeaders(Headers))
}

func request_Complete(data *Sender) (string, error) {
	d, err := GetPool().AddBodyStruct(data).Post(APIComplete)
	if err != nil {
		return "", &Err{Op: "request_Complete", Err: err.Error()}
	}
	var res *Response
	if err := json.Unmarshal(d.Body, res); err != nil {
		return "", &Err{Op: "request_Complete", Err: err.Error()}
	}
	if d.StatusCode != 200 {
		return "", &Err{Op: "request_Complete", Err: res.Stop}
	}
	return res.Completion, nil
}
