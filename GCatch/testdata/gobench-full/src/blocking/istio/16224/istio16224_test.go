package istio16224

import (
	"sync"
	"testing"
)

type ConfigStoreCache interface {
	RegisterEventHandler(handler func())
	Run()
}

type Event int

type Handler func(Event)

type configstoreMonitor struct {
	handlers []Handler
	eventCh  chan Event
}

func (m *configstoreMonitor) Run(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			if _, ok := <-m.eventCh; ok {
				close(m.eventCh)
			}
			return
		case ce, ok := <-m.eventCh:
			if ok {
				m.processConfigEvent(ce)
			}
		}
	}
}

func (m *configstoreMonitor) processConfigEvent(ce Event) {
	m.applyHandlers(ce)
}

func (m *configstoreMonitor) AppendEventHandler(h Handler) {
	m.handlers = append(m.handlers, h)
}

func (m *configstoreMonitor) applyHandlers(e Event) {
	for _, f := range m.handlers {
		f(e)
	}
}
func (m *configstoreMonitor) ScheduleProcessEvent(configEvent Event) {
	m.eventCh <- configEvent
}

type Monitor interface {
	Run(<-chan struct{})
	AppendEventHandler(Handler)
	ScheduleProcessEvent(Event)
}

type controller struct {
	monitor Monitor
}

func (c *controller) RegisterEventHandler(f func(Event)) {
	c.monitor.AppendEventHandler(f)
}

func (c *controller) Run(stop <-chan struct{}) {
	c.monitor.Run(stop)
}

func (c *controller) Create() {
	c.monitor.ScheduleProcessEvent(Event(0))
}

func NewMonitor() Monitor {
	return NewBufferedMonitor()
}

func NewBufferedMonitor() Monitor {
	return &configstoreMonitor{
		eventCh: make(chan Event),
	}
}
func TestIstio16224(t *testing.T) {
	controller := &controller{monitor: NewMonitor()}
	done := make(chan bool)
	lock := sync.Mutex{}
	controller.RegisterEventHandler(func(event Event) {
		lock.Lock()
		defer lock.Unlock()
		done <- true
	})

	stop := make(chan struct{})
	go controller.Run(stop)

	controller.Create()

	lock.Lock()
	lock.Unlock()
	<-done

	close(stop)
}
