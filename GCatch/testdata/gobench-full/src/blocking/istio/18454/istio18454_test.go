package istio18454

import (
	"context"

	"sync"
	"testing"
	"time"
)

type Worker struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
}

func (w *Worker) Start(setupFn func(), runFn func(c context.Context)) {
	if setupFn != nil {
		setupFn()
	}
	go func() {
		runFn(w.ctx)
	}()
}

func (w *Worker) Stop() {
	w.ctxCancel()
}

type Strategy struct {
	timer          *time.Timer
	timerFrequency time.Duration
	stateLock      sync.Mutex
	resetChan      chan struct{}
	worker         *Worker
	startTimerFn   func()
}

func (s *Strategy) OnChange() {
	s.stateLock.Lock()
	if s.timer != nil {
		s.stateLock.Unlock()
		s.resetChan <- struct{}{}
		return
	}
	s.startTimerFn()
	s.stateLock.Unlock()
}

func (s *Strategy) startTimer() {
	s.timer = time.NewTimer(s.timerFrequency)
	eventLoop := func(ctx context.Context) {
		for {
			select {
			case <-s.timer.C:
			case <-s.resetChan:
				if !s.timer.Stop() {
					<-s.timer.C
				}
				s.timer.Reset(s.timerFrequency)
			case <-ctx.Done():
				s.timer.Stop()
				return
			}
		}
	}
	s.worker.Start(nil, eventLoop)
}

func (s *Strategy) Close() {
	s.worker.Stop()
}

type Event int

type Processor struct {
	stateStrategy *Strategy
	worker        *Worker
	eventCh       chan Event
}

func (p *Processor) processEvent() {
	p.stateStrategy.OnChange()
}

func (p *Processor) Start() {
	setupFn := func() {
		for i := 0; i < 1024; i++ {
			p.eventCh <- Event(0)
		}
	}
	runFn := func(ctx context.Context) {
		defer func() {
			p.stateStrategy.Close()
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case <-p.eventCh:
				p.processEvent()
			}
		}
	}
	p.worker.Start(setupFn, runFn)
}

func (p *Processor) Stop() {
	p.worker.Stop()
}

func NewWorker() *Worker {
	worker := &Worker{}
	worker.ctx, worker.ctxCancel = context.WithCancel(context.Background())
	return worker
}

func TestIstio18454(t *testing.T) {
	stateStrategy := &Strategy{
		timerFrequency: time.Nanosecond,
		resetChan:      make(chan struct{}, 1),
		worker:         NewWorker(),
	}
	stateStrategy.startTimerFn = stateStrategy.startTimer

	p := &Processor{
		stateStrategy: stateStrategy,
		worker:        NewWorker(),
		eventCh:       make(chan Event, 1024),
	}

	p.Start()
	defer p.Stop()
}
