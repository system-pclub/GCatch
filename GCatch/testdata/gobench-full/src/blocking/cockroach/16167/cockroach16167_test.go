/*
 * Project: cockroach
 * Issue or PR  : https://github.com/cockroachdb/cockroach/pull/16167
 * Buggy version: 36fa784aa846b46c29e077634c4e362635f6e74a
 * fix commit-id: d064942b067ab84628f79cbfda001fa3138d8d6e
 * Flaky: 1/100
 * Description:
 *   This is another example for deadlock caused by recursively
 * acquiring RWLock. There are two lock variables (systemConfigCond and systemConfigMu)
 * involved in this bug, but they are actually the same lock, which can be found from
 * the following code.
 *   There are two goroutine involved in this deadlock. The first goroutine acquires
 * systemConfigMu.Lock() firstly, then tries to acquire systemConfigMu.RLock(). The
 * second goroutine tries to acquire systemConfigMu.Lock(). If the second goroutine
 * interleaves in between the two lock operations of the first goroutine, deadlock will happen.
 */

package cockroach16167

import (
	"sync"
	"testing"
)

type PreparedStatements struct {
	session *Session
}

func (ps PreparedStatements) New(e *Executor) {
	e.Prepare(ps.session)
}

type Session struct {
	PreparedStatements PreparedStatements
}

func (s *Session) resetForBatch(e *Executor) {
	e.getDatabaseCache()
}

type Executor struct {
	systemConfigCond *sync.Cond
	systemConfigMu   sync.RWMutex
}

func (e *Executor) Start() {
	e.updateSystemConfig()
}

func (e *Executor) execParsed(session *Session) {
	e.systemConfigCond.L.Lock() // Same as e.systemConfigMu.RLock()
	defer e.systemConfigCond.L.Unlock()
	runTxnAttempt(e, session)
}

func (e *Executor) execStmtsInCurrentTxn(session *Session) {
	e.execStmtInOpenTxn(session)
}

func (e *Executor) execStmtInOpenTxn(session *Session) {
	session.PreparedStatements.New(e)
}

func (e *Executor) Prepare(session *Session) {
	session.resetForBatch(e)
}

func (e *Executor) getDatabaseCache() {
	e.systemConfigMu.RLock()
	defer e.systemConfigMu.RUnlock()
}

func (e *Executor) updateSystemConfig() {
	e.systemConfigMu.Lock() // Block here
	defer e.systemConfigMu.Unlock()
}

func runTxnAttempt(e *Executor, session *Session) {
	e.execStmtsInCurrentTxn(session)
}

func NewExectorAndSession() (*Executor, *Session) {
	session := &Session{}
	session.PreparedStatements = PreparedStatements{session}
	e := &Executor{}
	return e, session
}

/// G1 							G2
/// e.Start()
/// e.updateSystemConfig()
/// 							e.execParsed()
/// 							e.systemConfigCond.L.Lock()
/// e.systemConfigMu.Lock()
/// 							e.systemConfigMu.RLock()
/// ----------------------G1,G2 deadlock--------------------
func TestCockroach16167(t *testing.T) {
	e, s := NewExectorAndSession()
	e.systemConfigCond = sync.NewCond(e.systemConfigMu.RLocker())
	go e.Start()    // G1
	e.execParsed(s) // G2
}
