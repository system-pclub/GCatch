/*
 * Project: cockroach
 * Issue or PR  : https://github.com/cockroachdb/cockroach/pull/18101
 * Buggy version: f7a8e2f57b6bcf00b9abaf3da00598e4acd3a57f
 * fix commit-id: 822bd176cc725c6b50905ea615023200b395e14f
 * Flaky: 100/100
 * Description:
 *   context.Done() signal only stops the goroutine who pulls data
 * from a channel, while does not stops goroutines which send data
 * to the channel. This causes all goroutines trying to send data
 * through the channel to block.
 */

package cockroach18101

import (
	"context"
	"testing"
)

const chanSize = 6

func restore(ctx context.Context) bool {
	readyForImportCh := make(chan bool, chanSize)
	go func() { // G2
		defer close(readyForImportCh)
		splitAndScatter(ctx, readyForImportCh)
	}()
	for readyForImportSpan := range readyForImportCh {
		select {
		case <-ctx.Done():
			return readyForImportSpan
		}
	}
	return true
}

func splitAndScatter(ctx context.Context, readyForImportCh chan bool) {
	for i := 0; i < chanSize+2; i++ {
		readyForImportCh <- (false || i != 0)
	}
}

///
/// G1					G2					helper goroutine
/// restore()
/// 					splitAndScatter()
/// <-readyForImportCh
/// 					readyForImportCh<-
/// ...					...
/// 										cancel()
/// return
/// 					readyForImportCh<-
/// -----------------------G2 leak-------------------------
///
func TestCockroach18101(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go restore(ctx) // G1
	go cancel()     // helper goroutine
}
