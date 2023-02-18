/*
 * Project: cockroach
 * Issue or PR  : https://github.com/cockroachdb/cockroach/pull/6181
 * Buggy version: c0a232b5521565904b851699853bdbd0c670cf1e
 * fix commit-id: d5814e4886a776bf7789b3c51b31f5206480d184
 * Flaky: 57/100
 */
package cockroach6181

import (
	"fmt"
	"sync"
	"testing"
)

type testDescriptorDB struct {
	cache *rangeDescriptorCache
}

func initTestDescriptorDB() *testDescriptorDB {
	return &testDescriptorDB{&rangeDescriptorCache{}}
}

type rangeDescriptorCache struct {
	rangeCacheMu sync.RWMutex
}

func (rdc *rangeDescriptorCache) LookupRangeDescriptor() {
	rdc.rangeCacheMu.RLock()
	fmt.Printf("lookup range descriptor: %s", rdc)
	rdc.rangeCacheMu.RUnlock()
	rdc.rangeCacheMu.Lock()
	rdc.rangeCacheMu.Unlock()
}

func (rdc *rangeDescriptorCache) String() string {
	rdc.rangeCacheMu.RLock()
	defer rdc.rangeCacheMu.RUnlock()
	return rdc.stringLocked()
}

func (rdc *rangeDescriptorCache) stringLocked() string {
	return "something here"
}

func doLookupWithToken(rc *rangeDescriptorCache) {
	rc.LookupRangeDescriptor()
}

func testRangeCacheCoalescedRquests() {
	db := initTestDescriptorDB()
	pauseLookupResumeAndAssert := func() {
		var wg sync.WaitGroup
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func() { // G2,G3,...
				doLookupWithToken(db.cache)
				wg.Done()
			}()
		}
		wg.Wait()
	}
	pauseLookupResumeAndAssert()
}

/// G1 									G2							G3					...
/// testRangeCacheCoalescedRquests()
/// initTestDescriptorDB()
/// pauseLookupResumeAndAssert()
/// return
/// 									doLookupWithToken()
///																 	doLookupWithToken()
///										rc.LookupRangeDescriptor()
///																	rc.LookupRangeDescriptor()
///										rdc.rangeCacheMu.RLock()
///										rdc.String()
///																	rdc.rangeCacheMu.RLock()
///																	fmt.Printf()
///																	rdc.rangeCacheMu.RUnlock()
///																	rdc.rangeCacheMu.Lock()
///										rdc.rangeCacheMu.RLock()
/// -------------------------------------G2,G3,... deadlock--------------------------------------

func TestCockroach6181(t *testing.T) {
	go testRangeCacheCoalescedRquests() // G1
}
