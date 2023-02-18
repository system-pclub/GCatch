package cockroach4407

import (
	"sync"
	"testing"
	"time"
)

type Stopper struct {
	stopper chan struct{}
	stop    sync.WaitGroup
	mu      sync.Mutex
}

func (s *Stopper) RunWorker(f func()) {
	s.stop.Add(1)
	go func() {
		defer s.stop.Done()
		f()
	}()
}

func (s *Stopper) SetStopped() {
	if s != nil {
		s.stop.Done()
	}
}

func (s *Stopper) Stop() {
	close(s.stopper)
	s.stop.Wait()
	s.mu.Lock()
	defer s.mu.Unlock()
}

type server struct {
	mu      sync.Mutex
	stopper *Stopper
}

func (s *server) Gossip() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stopper.RunWorker(func() {
		s.gossipSender()
	})
}

func (s *server) gossipSender() {
	s.mu.Lock()
	defer s.mu.Unlock()
}

func NewStopper() *Stopper {
	return &Stopper{
		stopper: make(chan struct{}),
	}
}

func TestCockroach4407(t *testing.T) {
	stopper := NewStopper()
	defer stopper.Stop()
	s := &server{
		stopper: stopper,
	}
	for i := 0; i < 2; i++ {
		go s.Gossip()
	}
	time.Sleep(time.Millisecond)
}
