/*
 * Project: moby
 * Issue or PR  : https://github.com/moby/moby/pull/17176
 * Buggy version: d295dc66521e2734390473ec1f1da8a73ad3288a
 * fix commit-id: 2f16895ee94848e2d8ad72bc01968b4c88d84cb8
 * Flaky: 100/100
 * Description:
 *   devices.nrDeletedDevices takes devices.Lock() but does
 * not drop it if there are no deleted devices. This will block
 * other goroutines trying to acquire devices.Lock().
 *   In general reason is that when device deletion is happning,
 * we can try deletion/deactivation in a loop. And that that time
 * we don't want to block rest of the device operations in parallel.
 * So we drop the inner devices lock while continue to hold per
 * device lock
 *   A test is added for this bug, and we need to try whether
 * this bug can be reproduced.
 */
package moby17176

import (
	"errors"
	"sync"
	"testing"
	"time"
)

type DeviceSet struct {
	sync.Mutex
	nrDeletedDevices int
}

func (devices *DeviceSet) cleanupDeletedDevices() error {
	devices.Lock()
	if devices.nrDeletedDevices == 0 {
		/// Missing devices.Unlock()
		return nil
	}
	devices.Unlock()
	return errors.New("Error")
}

func testDevmapperLockReleasedDeviceDeletion() {
	ds := &DeviceSet{
		nrDeletedDevices: 0,
	}
	ds.cleanupDeletedDevices()
	doneChan := make(chan bool)
	go func() {
		ds.Lock()
		defer ds.Unlock()
		doneChan <- true
	}()

	select {
	case <-time.After(time.Millisecond):
	case <-doneChan:
	}
}
func TestMoby17176(t *testing.T) {
	testDevmapperLockReleasedDeviceDeletion()
}
