package kubernetes49404

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Handler interface {
	ServeHTTP()
}

type websocket_Handler func()

func (h websocket_Handler) ServeHTTP() {
	h()
}

type muxEntry struct {
	h Handler
}

type ServeMux struct {
	es []muxEntry
}

func (mux *ServeMux) match() Handler {
	for _, e := range mux.es {
		return e.h
	}
	return nil
}

func (mux *ServeMux) handler() (h Handler) {
	h = mux.match()
	return
}

func (mux *ServeMux) Handler() Handler {
	return mux.handler()
}

func (mux *ServeMux) Handle(handler Handler) {
	e := muxEntry{h: handler}
	mux.es = appendSorted(mux.es, e)
}

func (mux *ServeMux) ServeHTTP() {
	h := mux.Handler()
	h.ServeHTTP()
}

func appendSorted(es []muxEntry, e muxEntry) []muxEntry {
	n := len(es)
	i := 0
	if i == n {
		return append(es, e)
	}
	es = append(es, muxEntry{})
	copy(es[i+1:], es[i:])
	es[i] = e
	return es
}

func NewServeMux() *ServeMux {
	return new(ServeMux)
}

type Server struct {
	Config *http_Server
	wg     sync.WaitGroup
}

func (s *Server) StartTLS() {
	s.goServe()
}

func (s *Server) goServe() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.Config.Serve()
	}()
}

type conn struct {
	server *http_Server
}

func (c *conn) serve() {
	serverHandler{c.server}.ServeHTTP()
}

type serverHandler struct {
	srv *http_Server
}

func (sh serverHandler) ServeHTTP() {
	handler := sh.srv.Handler
	handler.ServeHTTP()
}

type http_Server struct {
	Handler Handler
}

func (srv *http_Server) Serve() {
	c := srv.newConn()
	go c.serve()
}

func (srv *http_Server) newConn() *conn {
	c := &conn{
		server: srv,
	}
	return c
}

func NewUnstartedServer(handler Handler) *Server {
	return &Server{Config: &http_Server{Handler: handler}}
}

func TestKubernetes49404(t *testing.T) {
	ExpectCalled := true
	called := false
	func() {
		backendHandler := NewServeMux()
		backendHandler.Handle(websocket_Handler(func() {
			called = true
		}))

		backendServer := NewUnstartedServer(backendHandler)

		backendServer.StartTLS()

		defer func() {
			if called != ExpectCalled {
				_ = fmt.Sprintf("Error")
			}
		}()
	}()
	time.Sleep(10 * time.Millisecond)
}
