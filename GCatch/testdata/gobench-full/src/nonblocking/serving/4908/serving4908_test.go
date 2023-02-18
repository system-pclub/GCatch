package serving4908

import (
	"sync"
	"testing"
)

type TestingT interface {
	Logf(string, ...interface{})
}

type WriteSyncer interface {
	Write()
}

type CheckedEntry struct {
	ErrorOutput WriteSyncer
	cores       []Core
}

func (ce *CheckedEntry) Write() {
	for i := range ce.cores {
		ce.cores[i].Write()
	}
}

type testingWriter struct {
	t TestingT
}

func newTestingWriter(t TestingT) testingWriter {
	return testingWriter{t: t}
}

func (w testingWriter) Write() {
	w.t.Logf("%s", "1")
}

type Logger struct {
	core Core
}

func (log *Logger) clone() *Logger {
	copy := *log
	return &copy
}

func (log *Logger) Check() *CheckedEntry {
	ent := &CheckedEntry{}
	ent.cores = append(ent.cores, log.core)
	return ent
}

func NewLogger(t TestingT) *Logger {
	writer := newTestingWriter(t)
	return New(NewCore(writer))
}

func New(core Core) *Logger {
	return &Logger{
		core: core,
	}
}

type Core interface {
	Write()
}

type ioCore struct {
	out WriteSyncer
}

func (c *ioCore) Write() {
	c.out.Write()
}

func NewCore(ws WriteSyncer) Core {
	return &ioCore{
		out: ws,
	}
}

func testing_TestLogger(t *testing.T) *SugaredLogger {
	return NewLogger(t).Sugar()
}

func (log *Logger) Sugar() *SugaredLogger {
	return &SugaredLogger{log.clone()}
}

type SugaredLogger struct {
	base *Logger
}

func (s *SugaredLogger) log() {
	ce := s.base.Check()
	ce.Write()
}

func (s *SugaredLogger) Info(args ...interface{}) {
	s.log()
}

type revisionWatcher struct {
	logger *SugaredLogger
}

func newRevisionWatcher(logger *SugaredLogger) *revisionWatcher {
	return &revisionWatcher{
		logger: logger,
	}
}

func (rw *revisionWatcher) runWithTickCh() {
	rw.checkDests()
}

func (rw *revisionWatcher) checkDests() {
	go func() {
		rw.logger.Info("1")
	}()
}

func TestServing4908(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		t.Run("TestServing4908", func(t *testing.T) {
			rw := newRevisionWatcher(
				testing_TestLogger(t),
			)
			var _wg sync.WaitGroup
			_wg.Add(1)
			go func() {
				rw.runWithTickCh()
				_wg.Done()
			}()
			_wg.Wait()
		})
	}()
	wg.Wait()
}
