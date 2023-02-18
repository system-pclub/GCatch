package etcd4876

import (
	"sync"
	"testing"
	"time"
)

var ProgressReportInterval = 10 * time.Second

type Watcher interface {
	Watch()
}
type ServerStream interface{}

type Watch_WatchServer interface {
	Send()
	ServerStream
}
type watchWatchServer struct {
	ServerStream
}

func (x *watchWatchServer) Send() {}

type WatchServer interface {
	Watch(Watch_WatchServer)
}

type serverWatchStream struct{}

func (sws *serverWatchStream) sendLoop() {
	_ = time.NewTicker(ProgressReportInterval)
}

type watchServer struct{}

func (ws *watchServer) Watch(stream Watch_WatchServer) {
	sws := serverWatchStream{}
	go sws.sendLoop()
}

func TestEtcd4876(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		w := &watchServer{}
		go func() {
			defer wg.Done()
			testInterval := 3 * time.Second
			ProgressReportInterval = testInterval
		}()
		go func() {
			defer wg.Done()
			w.Watch(&watchWatchServer{})
		}()
	}()
	wg.Wait()
}
