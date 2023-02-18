package grpc1687

import (
	"testing"
	"time"
)

type ResponseWriter interface {
}

type testHandlerResponseWriter struct {
}

func newTestHandlerResponseWriter() ResponseWriter {
	return testHandlerResponseWriter{}
}

type serverHandlerTransport struct {
	closedCh chan struct{}
	writes   chan func()
}

func (ht *serverHandlerTransport) do(fn func()) {
	select {
	case <-ht.closedCh:
		return
	default:
		select {
		case ht.writes <- fn:
			return
		case <-ht.closedCh:
			return
		}
	}
}

func (ht *serverHandlerTransport) WriteStatus() {
	ht.do(func() {})
	close(ht.writes)
}

func (ht *serverHandlerTransport) Write() {
	ht.do(func() {})
}

func (ht *serverHandlerTransport) runStream() {
	for {
		select {
		case fn, ok := <-ht.writes:
			if !ok {
				return
			}
			fn()
		case <-ht.closedCh:
			return
		}
	}
}

func (ht *serverHandlerTransport) HandleStreams(startStream func()) {
	startStream()

	ht.runStream()
}

type ServerTransport interface {
	HandleStreams(func())
	Write()
	WriteStatus()
}

func NewServerHandlerTransport(writer ResponseWriter) ServerTransport {
	st := &serverHandlerTransport{
		closedCh: make(chan struct{}),
		writes:   make(chan func()),
	}
	return st
}

type handleStreamTest struct {
	t  *testing.T
	rw testHandlerResponseWriter
	ht *serverHandlerTransport
}

func newHandleStreamTest(t *testing.T) *handleStreamTest {
	rw := newTestHandlerResponseWriter().(testHandlerResponseWriter)
	ht := NewServerHandlerTransport(rw)
	return &handleStreamTest{
		t:  t,
		rw: rw,
		ht: ht.(*serverHandlerTransport),
	}
}

func testHandlerTransportHandleStreams(t *testing.T, handleStream func(st *handleStreamTest)) {
	st := newHandleStreamTest(t)
	st.ht.HandleStreams(func() { go handleStream(st) })
}

func TestGrpc1687(t *testing.T) {
	testHandlerTransportHandleStreams(t, func(st *handleStreamTest) {
		st.ht.WriteStatus()
		st.ht.Write()
	})
	time.Sleep(10 * time.Millisecond)
}
