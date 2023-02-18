package moby29733

import (
	"sync"
	"testing"
)

type Plugin struct {
	activated    bool
	activateWait *sync.Cond
}

type plugins struct {
	sync.Mutex
	plugins map[int]*Plugin
}

func (p *Plugin) waitActive() {
	p.activateWait.L.Lock()
	for !p.activated {
		p.activateWait.Wait()
	}
	p.activateWait.L.Unlock()
}

type extpointHandlers struct {
	sync.RWMutex
	extpointHandlers map[int]struct{}
}

var (
	storage  = plugins{plugins: make(map[int]*Plugin)}
	handlers = extpointHandlers{extpointHandlers: make(map[int]struct{})}
)

func Handle() {
	handlers.Lock()
	for _, p := range storage.plugins {
		p.activated = false
	}
	handlers.Unlock()
}

func testActive(p *Plugin) {
	done := make(chan struct{})
	go func() {
		p.waitActive()
		close(done)
	}()
	<-done
}

func TestMoby29733(t *testing.T) {
	p := &Plugin{activateWait: sync.NewCond(&sync.Mutex{})}
	storage.plugins[0] = p

	testActive(p)
	Handle()
	testActive(p)
}
