package cockroach1462

import (
	"sync"
	"testing"
)

type Stopper struct {
	stopper  chan struct{}
	stopped  chan struct{}
	stop     sync.WaitGroup
	mu       sync.Mutex
	drain    *sync.Cond
	draining bool
	numTasks int
}

func NewStopper() *Stopper {
	s := &Stopper{
		stopper: make(chan struct{}),
		stopped: make(chan struct{}),
	}
	s.drain = sync.NewCond(&s.mu)
	return s
}

func (s *Stopper) RunWorker(f func()) {
	s.AddWorker()
	go func() {
		defer s.SetStopped()
		f()
	}()
}

func (s *Stopper) AddWorker() {
	s.stop.Add(1)
}
func (s *Stopper) StartTask() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.draining {
		return false
	}
	s.numTasks++
	return true
}

func (s *Stopper) FinishTask() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.numTasks--
	s.drain.Broadcast()
}
func (s *Stopper) SetStopped() {
	if s != nil {
		s.stop.Done()
	}
}
func (s *Stopper) ShouldStop() <-chan struct{} {
	if s == nil {
		return nil
	}
	return s.stopper
}

func (s *Stopper) Quiesce() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.draining = true
	for s.numTasks > 0 {
		// Unlock s.mu, wait for the signal, and lock s.mu.
		s.drain.Wait()
	}
}

func (s *Stopper) Stop() {
	s.Quiesce()
	close(s.stopper)
	s.stop.Wait()
	s.mu.Lock()
	defer s.mu.Unlock()
	close(s.stopped)
}

type interceptMessage int

type localInterceptableTransport struct {
	mu      sync.Mutex
	Events  chan interceptMessage
	stopper *Stopper
}

func (lt *localInterceptableTransport) Close() {}

type Transport interface {
	Close()
}

func NewLocalInterceptableTransport(stopper *Stopper) Transport {
	lt := &localInterceptableTransport{
		Events:  make(chan interceptMessage),
		stopper: stopper,
	}
	lt.start()
	return lt
}

func (lt *localInterceptableTransport) start() {
	lt.stopper.RunWorker(func() {
		for {
			select {
			case <-lt.stopper.ShouldStop():
				return
			default:
				lt.Events <- interceptMessage(0)
			}
		}
	})
}

func processEventsUntil(ch <-chan interceptMessage, stopper *Stopper) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		case <-stopper.ShouldStop():
			return
		}
	}
}
func TestCockroach1462(t *testing.T) {
	stopper := NewStopper()
	transport := NewLocalInterceptableTransport(stopper).(*localInterceptableTransport)
	stopper.RunWorker(func() {
		processEventsUntil(transport.Events, stopper)
	})
	stopper.Stop()
}
