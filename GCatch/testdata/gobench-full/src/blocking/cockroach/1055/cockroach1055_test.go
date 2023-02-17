package cockroach1055

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Stopper struct {
	stopper  chan struct{}
	stop     sync.WaitGroup
	mu       sync.Mutex
	draining int32
	drain    sync.WaitGroup
}

func (s *Stopper) AddWorker() {
	s.stop.Add(1)
}

func (s *Stopper) ShouldStop() <-chan struct{} {
	if s == nil {
		return nil
	}
	return s.stopper
}

func (s *Stopper) SetStopped() {
	if s != nil {
		s.stop.Done()
	}
}

func (s *Stopper) Quiesce() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.draining = 1
	s.drain.Wait()
	s.draining = 0
}

func (s *Stopper) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()
	atomic.StoreInt32(&s.draining, 1)
	s.drain.Wait()
	close(s.stopper)
	s.stop.Wait()
}

func (s *Stopper) StartTask() bool {
	if atomic.LoadInt32(&s.draining) == 0 {
		s.mu.Lock()
		defer s.mu.Unlock()
		s.drain.Add(1)
		return true
	}
	return false
}

func NewStopper() *Stopper {
	return &Stopper{
		stopper: make(chan struct{}),
	}
}

func TestCockroach1055(t *testing.T) {
	var stoppers []*Stopper
	for i := 0; i < 3; i++ {
		stoppers = append(stoppers, NewStopper())
	}

	for i := range stoppers {
		s := stoppers[i]
		s.AddWorker()
		go func() {
			s.StartTask()
			<-s.ShouldStop()
			s.SetStopped()
		}()
	}

	done := make(chan struct{})
	go func() {
		for _, s := range stoppers {
			s.Quiesce()
		}
		for _, s := range stoppers {
			s.Stop()
		}
		close(done)
	}()

	<-done
}
