package pool

import (
	"errors"
	"io"
	"sync"
)

type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
}

// 新建
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size must greater than 0")
	}
	return &Pool{
		res:     make(chan io.Closer, size),
		factory: fn,
	}, nil
}

// 获取
func (p *Pool) Get() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		if !ok {
			return nil, errors.New("res having closed")
		}
		return r, nil
	default:
		return p.factory()
	}
}

// 关闭
func (p *Pool) Close() error {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return nil
	}
	p.closed = true
	close(p.res)
	for r := range p.res {
		if err := r.Close(); nil != err {
			return err
		}
	}

	return nil
}
