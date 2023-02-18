package hugo3251

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	remoteURLLock = &remoteLock{m: make(map[string]*sync.Mutex)}
)

type remoteLock struct {
	sync.RWMutex
	m map[string]*sync.Mutex
}

func (l *remoteLock) URLLock(url string) {
	l.Lock()
	if _, ok := l.m[url]; !ok {
		l.m[url] = &sync.Mutex{}
	}
	l.m[url].Lock()
	l.Unlock()
}

func (l *remoteLock) URLUnlock(url string) {
	l.RLock()
	defer l.RUnlock()
	if um, ok := l.m[url]; ok {
		um.Unlock()
	}
}

func resGetRemote(url string) error {
	remoteURLLock.URLLock(url)
	defer func() { remoteURLLock.URLUnlock(url) }()

	return nil
}

func TestHugo3251(t *testing.T) {
	url := "http://Foo.Bar/foo_Bar-Foo"
	for _ = range []bool{false, true} {
		var wg sync.WaitGroup
		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func(gor int) {
				defer wg.Done()
				for j := 0; j < 10; j++ {
					err := resGetRemote(url)
					if err != nil {
						fmt.Errorf("Error getting resource content: %s", err)
					}
					time.Sleep(300 * time.Nanosecond)
				}
			}(i)
		}
		wg.Wait()
	}
}
