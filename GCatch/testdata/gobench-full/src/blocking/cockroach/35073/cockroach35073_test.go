package cockroach35073

import (
	"sync"
	"sync/atomic"
	"testing"
)

type ConsumerStatus uint32

const (
	NeedMoreRows ConsumerStatus = iota
	DrainRequested
	ConsumerClosed
)

const rowChannelBufSize = 16
const outboxBufRows = 16

type rowSourceBase struct {
	consumerStatus ConsumerStatus
}

func (rb *rowSourceBase) consumerClosed() {
	atomic.StoreUint32((*uint32)(&rb.consumerStatus), uint32(ConsumerClosed))
}

type RowChannelMsg int

type RowChannel struct {
	rowSourceBase
	dataChan chan RowChannelMsg
}

func (rc *RowChannel) ConsumerClosed() {
	rc.consumerClosed()
	select {
	case <-rc.dataChan:
	default:
	}
}

func (rc *RowChannel) Push() ConsumerStatus {
	consumerStatus := ConsumerStatus(
		atomic.LoadUint32((*uint32)(&rc.consumerStatus)))
	switch consumerStatus {
	case NeedMoreRows:
		rc.dataChan <- RowChannelMsg(0)
	case DrainRequested:
	case ConsumerClosed:
	}
	return consumerStatus
}

func (rc *RowChannel) InitWithNumSenders() {
	rc.initWithBufSizeAndNumSenders(rowChannelBufSize)
}

func (rc *RowChannel) initWithBufSizeAndNumSenders(chanBufSize int) {
	rc.dataChan = make(chan RowChannelMsg, chanBufSize)
}

type outbox struct {
	RowChannel
}

func (m *outbox) init() {
	m.RowChannel.InitWithNumSenders()
}

func (m *outbox) start(wg *sync.WaitGroup) {
	if wg != nil {
		wg.Add(1)
	}
	go m.run(wg)
}

func (m *outbox) run(wg *sync.WaitGroup) {
	m.mainLoop()
	if wg != nil {
		wg.Done()
	}
}

func (m *outbox) mainLoop() {
	return
}

func TestCockroach35073(t *testing.T) {
	outbox := &outbox{}
	outbox.init()

	var wg sync.WaitGroup
	for i := 0; i < outboxBufRows; i++ {
		outbox.Push()
	}

	var blockedPusherWg sync.WaitGroup
	blockedPusherWg.Add(1)
	go func() {
		outbox.Push()
		blockedPusherWg.Done()
	}()

	outbox.start(&wg)

	wg.Wait()
	outbox.RowChannel.Push()
}
