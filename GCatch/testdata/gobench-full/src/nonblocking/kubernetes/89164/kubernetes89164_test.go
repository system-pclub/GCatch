package kubernetes89164

import (
	"sync"
	"testing"
)

type cacheWatcher int

type Cacher struct {
	sync.RWMutex
	watcherBuffer []*cacheWatcher
}

func (c *Cacher) startDispatching() {
	c.Lock()
	defer c.Unlock()

	c.watcherBuffer = c.watcherBuffer[:0]
}

func (c *Cacher) dispatchEvent() {
	c.startDispatching()
	for _ = range c.watcherBuffer {
	}
}

func (c *Cacher) dispatchEvents() {
	c.dispatchEvent()
}

func NewCacherFromConfig() *Cacher {
	cacher := &Cacher{}
	go cacher.dispatchEvents()
	return cacher
}

func newTestCacher() *Cacher {
	return NewCacherFromConfig()
}

func TestKubernetes89164(t *testing.T) {
	cacher := newTestCacher()
	for i := 0; i < 3; i++ {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			cacher.dispatchEvent()
			wg.Done()
		}()
		wg.Wait()
	}
}
