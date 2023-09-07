package workerpool

import (
	"github.com/cilium/workerpool"
	"github.com/cornelk/hashmap"
)

type WorkerPool struct {
	p *hashmap.Map[string, *workerpool.WorkerPool]
	l *hashmap.Map[string, bool]
}

func NewGlobal() *WorkerPool {
	return &WorkerPool{
		p: hashmap.New[string, *workerpool.WorkerPool](),
		l: hashmap.New[string, bool](),
	}
}

func (w *WorkerPool) NewPool(pool_id string, pool_process int) error {
	if pool_id == "" {
		return ErrPIDNotEmpty
	}
	if pool_process < 1 {
		pool_process = 1
	}
	_, ok := w.l.Get(pool_id)
	if ok {
		return ErrPIDHasLocked
	}
	w.p.Set(pool_id, workerpool.New(pool_process))
	return nil
}

func (w *WorkerPool) GetPool(pool_id string) (*workerpool.WorkerPool, bool) {
	if pool_id == "" {
		return nil, false
	}
	return w.p.Get(pool_id)
}

func (w *WorkerPool) DelPool(pool_id string) {
	w.p.Del(pool_id)
	w.l.Del(pool_id)
}
