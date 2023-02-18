/*
 * Project: kubernetes
 * Issue or PR  : https://github.com/kubernetes/kubernetes/pull/13135
 * Buggy version: 6ced66249d4fd2a81e86b4a71d8df0139fe5ceae
 * fix commit-id: a12b7edc42c5c06a2e7d9f381975658692951d5a
 * Flaky: 93/100
 */
package kubernetes13135

import (
	"sync"
	"testing"
	"time"
)

var (
	StopChannel chan struct{}
)

func Util(f func(), period time.Duration, stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		default:
		}
		func() {
			f()
		}()
		time.Sleep(period)
	}
}

type Store interface {
	Add(obj interface{})
	Replace(obj interface{})
}

type Reflector struct {
	store Store
}

func (r *Reflector) ListAndWatch(stopCh <-chan struct{}) error {
	r.syncWith()
	return nil
}

func NewReflector(store Store) *Reflector {
	return &Reflector{
		store: store,
	}
}

func (r *Reflector) syncWith() {
	r.store.Replace(nil)
}

type Cacher struct {
	sync.Mutex
	initialized sync.WaitGroup
	initOnce    sync.Once
	watchCache  *WatchCache
	reflector   *Reflector
}

func (c *Cacher) processEvent() {
	c.Lock()
	defer c.Unlock()
}

func (c *Cacher) startCaching(stopChannel <-chan struct{}) {
	c.Lock()
	for {
		err := c.reflector.ListAndWatch(stopChannel)
		if err == nil {
			break
		}
	}
}

type WatchCache struct {
	sync.RWMutex
	onReplace func()
	onEvent   func()
}

func (w *WatchCache) SetOnEvent(onEvent func()) {
	w.Lock()
	defer w.Unlock()
	w.onEvent = onEvent
}

func (w *WatchCache) SetOnReplace(onReplace func()) {
	w.Lock()
	defer w.Unlock()
	w.onReplace = onReplace
}

func (w *WatchCache) processEvent() {
	w.Lock()
	defer w.Unlock()
	if w.onEvent != nil {
		w.onEvent()
	}
}

func (w *WatchCache) Add(obj interface{}) {
	w.processEvent()
}

func (w *WatchCache) Replace(obj interface{}) {
	w.Lock()
	defer w.Unlock()
	if w.onReplace != nil {
		w.onReplace()
	}
}

func NewCacher() *Cacher {
	watchCache := &WatchCache{}
	cacher := &Cacher{
		initialized: sync.WaitGroup{},
		watchCache:  watchCache,
		reflector:   NewReflector(watchCache),
	}
	cacher.initialized.Add(1)
	watchCache.SetOnReplace(func() {
		cacher.initOnce.Do(func() { cacher.initialized.Done() })
		cacher.Unlock()
	})
	watchCache.SetOnEvent(cacher.processEvent)
	stopCh := StopChannel
	go Util(func() { cacher.startCaching(stopCh) }, 0, stopCh) // G2
	cacher.initialized.Wait()
	return cacher
}

///
/// G1								G2								G3
/// NewCacher()
/// watchCache.SetOnReplace()
/// watchCache.SetOnEvent()
/// 								cacher.startCaching()
///									c.Lock()
/// 								c.reflector.ListAndWatch()
/// 								r.syncWith()
/// 								r.store.Replace()
/// 								w.Lock()
/// 								w.onReplace()
/// 								cacher.initOnce.Do()
/// 								cacher.Unlock()
/// return cacher
///																	c.watchCache.Add()
///																	w.processEvent()
///																	w.Lock()
///									cacher.startCaching()
///									c.Lock()
///									...
///																	c.Lock()
///									w.Lock()
///--------------------------------G2,G3 deadlock-------------------------------------
///
func TestKubernetes13135(t *testing.T) {
	StopChannel = make(chan struct{})
	c := NewCacher()         // G1
	go c.watchCache.Add(nil) // G3
	go close(StopChannel)
}
