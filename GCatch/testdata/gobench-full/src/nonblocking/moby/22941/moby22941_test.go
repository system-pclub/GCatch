package moby22941

import (
	"sync"
	"testing"
	"time"
)

type Conn interface {
	Write(b []byte)
}

type pipe struct {
	wrMu sync.Mutex
}

func (p *pipe) Write(b []byte) {
	p.wrMu.Lock()
	defer p.wrMu.Unlock()
	b = b[1:]
}

func Pipe() Conn {
	return &pipe{}
}

func TestMoby22941(t *testing.T) {
	srv := Pipe()
	tests := [][2][]byte{
		{
			[]byte("GET /foo\nHost: /var/run/docker.sock\nUser-Agent: Docker\r\n\r\n"),
			[]byte("GET /foo\nHost: \r\nConnection: close\r\nUser-Agent: Docker\r\n\r\n"),
		},
		{
			[]byte("GET /foo\nHost: /var/run/docker.sock\nUser-Agent: Docker\nFoo: Bar\r\n"),
			[]byte("GET /foo\nHost: \r\nConnection: close\r\nUser-Agent: Docker\nFoo: Bar\r\n"),
		},
	}
	for _, pair := range tests {
		go func() {
			srv.Write(pair[0])
		}()
	}
	time.Sleep(10 * time.Millisecond)
}
