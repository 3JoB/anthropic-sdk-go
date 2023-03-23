package anthropic

import (
	"sync"

	fast "github.com/3JoB/fasthttp-client"
)

var (
	Pool sync.Pool = fast.NewClientPool()
)

func GetPool() *fast.Client {
	return Pool.Get().(*fast.Client)
}

func SetPool() {
	Pool.Put(GetPool().AddHeaders(Headers))
}