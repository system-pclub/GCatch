/*
 * Project: moby
 * Issue or PR  : https://github.com/moby/moby/pull/7559
 * Buggy version: 64579f51fcb439c36377c0068ccc9a007b368b5a
 * fix commit-id: 6cbb8e070d6c3a66bf48fbe5cbf689557eee23db
 * Flaky: 100/100
 */
package moby7559

import (
	"net"
	"sync"
	"testing"
)

type UDPProxy struct {
	connTrackLock sync.Mutex
}

func (proxy *UDPProxy) Run() {
	for i := 0; i < 2; i++ {
		proxy.connTrackLock.Lock()
		_, err := net.DialUDP("udp", nil, nil)
		if err != nil {
			/// Missing unlock here
			continue
		}
		if i == 0 {
			break
		}
	}
	proxy.connTrackLock.Unlock()
}
func TestMoby7559(t *testing.T) {
	proxy := &UDPProxy{}
	go proxy.Run()
}
