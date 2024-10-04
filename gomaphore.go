package gomaphore

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	errTimeout = errors.New("timeout")
)

type Gomaphore struct {
	resources chan struct{}
	once      sync.Once
}

func New(maxConcurrency int) *Gomaphore {
	return &Gomaphore{
		resources: make(chan struct{}, maxConcurrency),
	}
}

func (g *Gomaphore) Wait() {
	g.resources <- struct{}{}
}

func (g *Gomaphore) WaitWithTimeout(timeout int) error {
	timer := time.NewTimer(time.Duration(timeout) * time.Millisecond)
	defer timer.Stop()

	select {
	case g.resources <- struct{}{}:
		return nil
	case <-timer.C:
		return errTimeout
	}
}

func (g *Gomaphore) WaitWithContext(ctx context.Context) error {
	select {
	case g.resources <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (g *Gomaphore) Release() {
	<-g.resources
}

func (g *Gomaphore) Close() {
	g.once.Do(func() {
		close(g.resources)
	})
}
