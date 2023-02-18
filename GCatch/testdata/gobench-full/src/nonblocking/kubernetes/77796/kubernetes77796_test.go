package kubernetes77796

import (
	"sync"
	"testing"
	"time"
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

func TestKubernetes77796(t *testing.T) {
	cacher := newTestCacher()
	for i := 0; i < 3; i++ {
		go func() {
			cacher.dispatchEvent()
		}()
		time.Sleep(10 * time.Millisecond)
	}
}
