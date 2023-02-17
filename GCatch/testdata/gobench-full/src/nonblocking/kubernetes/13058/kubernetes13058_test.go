package kubernetes13058

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type ProcessFunc func(obj interface{})

type Config struct {
	Process ProcessFunc
}

type ResourceEventHandler interface {
	OnDelete(obj interface{})
}

type ResourceEventHandlerFuncs struct {
	DeleteFunc func(obj interface{})
}

func (r ResourceEventHandlerFuncs) OnDelete(obj interface{}) {
	if r.DeleteFunc != nil {
		r.DeleteFunc(obj)
	}
}

type Controller struct {
	config Config
}

func (c *Controller) processLoop() {
	for {
		c.config.Process(nil)
		break
	}
}

func (c *Controller) Run(stopCh <-chan struct{}) {
	Until(c.processLoop, 10*time.Millisecond, stopCh)
}

func New(c *Config) *Controller {
	ctlr := &Controller{config: *c}
	return ctlr
}

func NewInformer(h ResourceEventHandler) *Controller {
	cfg := &Config{
		Process: func(obj interface{}) {
			h.OnDelete(obj)
		},
	}
	return New(cfg)
}

func Until(f func(), period time.Duration, stopCh <-chan struct{}) {
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

func TestKubernetes13058(t *testing.T) {
	var testDoneWG sync.WaitGroup

	controller := NewInformer(ResourceEventHandlerFuncs{
		DeleteFunc: func(obj interface{}) {
			testDoneWG.Done()
		},
	})

	stop := make(chan struct{})
	go controller.Run(stop)

	tests := []func(string){
		func(name string) {},
	}

	const threads = 3
	var wg sync.WaitGroup
	wg.Add(threads * len(tests))
	testDoneWG.Add(threads * len(tests))
	for i := 0; i < threads; i++ {
		for j, f := range tests {
			go func(name string, f func(string)) {
				defer wg.Done()
				f(name)
			}(fmt.Sprintf("%v-%v", i, j), f)
		}
	}
	wg.Wait()
	testDoneWG.Wait()
	close(stop)
}
