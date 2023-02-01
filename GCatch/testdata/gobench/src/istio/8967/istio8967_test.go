package istio8967

import (
	"sync"
	"testing"
	"time"
)

type Source interface {
	Start()
	Stop()
}

type fsSource struct {
	donec chan struct{}
}

func (s *fsSource) Start() {
	go func() {
		for {
			select {
			case <-s.donec:
				return
			}
		}
	}()
}

func (s *fsSource) Stop() {
	close(s.donec)
	s.donec = nil
}

func newFsSource() *fsSource {
	return &fsSource{
		donec: make(chan struct{}),
	}
}

func New() Source {
	return newFsSource()
}

func TestIstio8967(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s := New()
		s.Start()
		s.Stop()
		time.Sleep(5 * time.Millisecond)
	}()
	wg.Wait()
}
