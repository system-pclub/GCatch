/*
 * Project: grpc-go
 * Issue or PR  : https://github.com/grpc/grpc-go/pull/1275
 * Buggy version: (missing)
 * fix commit-id: 0669f3f89e0330e94bb13fa1ce8cc704aab50c9c
 * Flaky: 100/100
 * Description:
 *   Two goroutines are invovled in this deadlock. The first goroutine
 * is the main goroutine. It is blocked at case <- donec, and it is
 * waiting for the second goroutine to close the channel.
 *   The second goroutine is created by the main goroutine. It is blocked
 * when calling stream.Read(). stream.Read() invokes recvBufferRead.Read().
 * The second goroutine is blocked at case i := r.recv.get(), and it is
 * waiting for someone to send a message to this channel.
 *   It is the client.CloseSream() method called by the main goroutine that
 * should send the message, but it is not. The patch is to send out this message.
 */
package grpc1275

import (
	"io"
	"testing"
	"time"
)

type recvBuffer struct {
	c chan bool
}

func (b *recvBuffer) get() <-chan bool {
	return b.c
}

type recvBufferReader struct {
	recv *recvBuffer
}

func (r *recvBufferReader) Read(p []byte) (int, error) {
	select {
	case <-r.recv.get(): // G2 block here
	}
	return 0, nil
}

type Stream struct {
	trReader io.Reader
}

func (s *Stream) Read(p []byte) (int, error) {
	return io.ReadFull(s.trReader, p)
}

type http2Client struct{}

func (t *http2Client) CloseStream(s *Stream) {
	// It is the client.CloseSream() method called by the
	// main goroutine that should send the message, but it
	// is not. The patch is to send out this message.
}

func (t *http2Client) NewStream() *Stream {
	return &Stream{
		trReader: &recvBufferReader{
			recv: &recvBuffer{
				c: make(chan bool),
			},
		},
	}
}

func testInflightStreamClosing() {
	client := &http2Client{}
	stream := client.NewStream()
	donec := make(chan bool)
	go func() { // G2
		defer close(donec)
		stream.Read([]byte{1})
	}()

	client.CloseStream(stream)

	timeout := time.NewTimer(300 * time.Nanosecond)
	select {
	case <-donec:
		if !timeout.Stop() {
			<-timeout.C
		}
	case <-timeout.C:
	}
}

///
/// G1 									G2
/// testInflightStreamClosing()
/// 									stream.Read()
/// 									io.ReadFull()
/// 									<- r.recv.get()
/// CloseStream()
/// <- donec
/// ------------G1 timeout, G2 leak---------------------
///

func TestGrpc1293(t *testing.T) {
	testInflightStreamClosing() // G1
}
