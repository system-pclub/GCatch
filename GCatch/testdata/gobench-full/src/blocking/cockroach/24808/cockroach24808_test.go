package cockroach24808

import (
	"context"
	"sync"
	"testing"
)

type Compactor struct {
	ch chan struct{}
}

type Stopper struct {
	stop    sync.WaitGroup
	stopper chan struct{}
}

func (s *Stopper) RunWorker(ctx context.Context, f func(context.Context)) {
	s.stop.Add(1)
	go func() {
		defer s.stop.Done()
		f(ctx)
	}()
}

func (s *Stopper) ShouldStop() <-chan struct{} {
	if s == nil {
		return nil
	}
	return s.stopper
}

func (s *Stopper) Stop() {
	close(s.stopper)
}

func NewStopper() *Stopper {
	s := &Stopper{
		stopper: make(chan struct{}),
	}
	return s
}

func NewCompactor() *Compactor {
	return &Compactor{ch: make(chan struct{}, 1)}
}

func (c *Compactor) Start(ctx context.Context, stopper *Stopper) {
	c.ch <- struct{}{}
	stopper.RunWorker(ctx, func(ctx context.Context) {
		for {
			select {
			case <-stopper.ShouldStop():
				return
			case <-c.ch:
			}
		}
	})
}

func TestCockroach24808(t *testing.T) {
	stopper := NewStopper()
	defer stopper.Stop()

	compactor := NewCompactor()
	compactor.ch <- struct{}{}

	compactor.Start(context.Background(), stopper)
}
