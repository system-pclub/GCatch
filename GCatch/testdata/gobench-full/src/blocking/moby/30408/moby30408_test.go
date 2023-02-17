package moby30408

import (
	"errors"
	"sync"
	"testing"
)

type Manifest struct {
	Implements []string
}

type Plugin struct {
	activateWait *sync.Cond
	activateErr  error
	Manifest     *Manifest
}

func (p *Plugin) waitActive() error {
	p.activateWait.L.Lock()
	for !p.activated() {
		p.activateWait.Wait()
	}
	p.activateWait.L.Unlock()
	return p.activateErr
}

func (p *Plugin) activated() bool {
	return p.Manifest != nil
}

func testActive(p *Plugin) {
	done := make(chan struct{})
	go func() {
		p.waitActive()
		close(done)
	}()
	<-done
}
func TestMoby30408(t *testing.T) {
	p := &Plugin{activateWait: sync.NewCond(&sync.Mutex{})}
	p.activateErr = errors.New("some junk happened")

	testActive(p)
}
