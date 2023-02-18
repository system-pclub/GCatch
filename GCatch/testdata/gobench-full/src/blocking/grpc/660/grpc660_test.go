/*
 * Project: grpc-go
 * Issue or PR  : https://github.com/grpc/grpc-go/pull/660
 * Buggy version: db85417dd0de6cc6f583672c6175a7237e5b5dd2
 * fix commit-id: ceacfbcbc1514e4e677932fd55938ac455d182fb
 * Flaky: 100/100
 * Description:
 *   The parent function could return without draining the done channel.
 */
package grpc660

import (
	"math/rand"
	"testing"
)

type benchmarkClient struct {
	stop chan bool
}

func (bc *benchmarkClient) doCloseLoopUnary() {
	for {
		done := make(chan bool)
		go func() { // G2
			if rand.Intn(10) > 7 {
				done <- false
				return
			}
			done <- true
		}()
		select {
		case <-bc.stop:
			return
		case <-done:
		}
	}
}

///
/// G1 						G2 				helper goroutine
/// doCloseLoopUnary()
///											bc.stop <- true
/// <-bc.stop
/// return
/// 						done <-
/// ----------------------G2 leak--------------------------
///

func TestGrpc660(t *testing.T) {
	bc := &benchmarkClient{
		stop: make(chan bool),
	}
	go bc.doCloseLoopUnary() // G1
	go func() {              // helper goroutine
		bc.stop <- true
	}()
}
