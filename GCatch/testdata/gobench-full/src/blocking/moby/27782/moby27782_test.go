/*
 * Project: moby
 * Issue or PR  : https://github.com/moby/moby/pull/27782
 * Buggy version: 18768fdc2e76ec6c600c8ab57d2d487ee7877794
 * fix commit-id: a69a59ffc7e3d028a72d1195c2c1535f447eaa84
 * Flaky: 2/100
 */
package moby27782

import (
	"errors"
	"sync"
	"testing"
	"time"
)

type Event struct {
	Op Op
}

type Op uint32

const (
	Create Op = 1 << iota
	Write
	Remove
	Rename
	Chmod
)

func newEvent(op Op) Event {
	return Event{op}
}

func (e *Event) ignoreLinux(w *Watcher) bool {
	if e.Op != Write {
		w.mu.Lock()
		defer w.mu.Unlock()
		w.cv.Broadcast()
		return true
	}
	return false
}

type Watcher struct {
	Events chan Event
	mu     sync.Mutex
	cv     *sync.Cond
	done   chan struct{}
}

func NewWatcher() *Watcher {
	w := &Watcher{
		Events: make(chan Event),
		done:   make(chan struct{}),
	}
	w.cv = sync.NewCond(&w.mu)
	go w.readEvents() // G3
	return w
}

func (w *Watcher) readEvents() {
	defer close(w.Events)
	for {
		if w.isClosed() {
			return
		}
		event := newEvent(Write) // MODIFY event
		if !event.ignoreLinux(w) {
			time.Sleep(300 * time.Nanosecond)
			select {
			case w.Events <- event:
			case <-w.done:
				return
			}
		}
	}
}

func (w *Watcher) isClosed() bool {
	select {
	case <-w.done:
		return true
	default:
		return false
	}
}

func (w *Watcher) Close() {
	if w.isClosed() {
		return
	}
	close(w.done)
}

func (w *Watcher) Remove() {
	w.mu.Lock()
	defer w.mu.Unlock()
	exists := true
	for exists {
		w.cv.Wait()
	}
}

type FileWatcher interface {
	Events() <-chan Event
	Remove()
	Close()
}

func New() FileWatcher {
	return NewEventWatcher()
}

func NewEventWatcher() FileWatcher {
	return &fsNotifyWatcher{NewWatcher()}
}

type fsNotifyWatcher struct {
	*Watcher
}

func (w *fsNotifyWatcher) Events() <-chan Event {
	return w.Watcher.Events
}

func watchFile() FileWatcher {
	fileWatcher := New()
	return fileWatcher
}

type LogWatcher struct {
	closeOnce     sync.Once
	closeNotifier chan struct{}
}

func (w *LogWatcher) Close() {
	w.closeOnce.Do(func() {
		close(w.closeNotifier)
	})
}

func (w *LogWatcher) WatchClose() <-chan struct{} {
	return w.closeNotifier
}

func NewLogWatcher() *LogWatcher {
	return &LogWatcher{
		closeNotifier: make(chan struct{}),
	}
}

func followLogs(logWatcher *LogWatcher) {
	fileWatcher := watchFile()
	defer func() {
		fileWatcher.Close()
	}()
	waitRead := func() {
		time.Sleep(300 * time.Nanosecond)
		select {
		case <-fileWatcher.Events():
		case <-logWatcher.WatchClose():
			fileWatcher.Remove()
			return
		}
	}
	handleDecodeErr := func() {
		waitRead()
	}
	handleDecodeErr()
}

type Container struct {
	LogDriver *JSONFileLogger
}

func (container *Container) InitializeStdio() {
	if err := container.startLogging(); err != nil {
		container.Reset()
	}
}

func (container *Container) startLogging() error {
	l := &JSONFileLogger{
		readers: make(map[*LogWatcher]struct{}),
	}
	container.LogDriver = l
	l.ReadLogs()
	return errors.New("Some error")
}

func (container *Container) Reset() {
	if container.LogDriver != nil {
		container.LogDriver.Close()
	}
}

type JSONFileLogger struct {
	readers map[*LogWatcher]struct{}
}

func (l *JSONFileLogger) ReadLogs() *LogWatcher {
	logWatcher := NewLogWatcher()
	go l.readLogs(logWatcher)
	return logWatcher
}

func (l *JSONFileLogger) readLogs(logWatcher *LogWatcher) {
	l.readers[logWatcher] = struct{}{}
	followLogs(logWatcher)
}

func (l *JSONFileLogger) Close() {
	for r := range l.readers {
		r.Close()
		delete(l.readers, r)
	}
}

///
/// G1 						G2							G3
/// InitializeStdio()
/// startLogging()
/// l.ReadLogs()
/// NewLogWatcher()
/// 						l.readLogs()
/// container.Reset()
/// LogDriver.Close()
/// r.Close()
/// close(w.closeNotifier)
/// 						followLogs(logWatcher)
/// 						watchFile()
/// 						New()
/// 						NewEventWatcher()
/// 						NewWatcher()
/// 													w.readEvents()
/// 													event.ignoreLinux()
/// 													return false
/// 						<-logWatcher.WatchClose()
/// 						fileWatcher.Remove()
/// 						w.cv.Wait()
/// 													w.Events <- event
/// ------------------------------G2,G3 deadlock---------------------------
///

func TestMoby27782(t *testing.T) {
	c := &Container{}
	go c.InitializeStdio() // G1
}
