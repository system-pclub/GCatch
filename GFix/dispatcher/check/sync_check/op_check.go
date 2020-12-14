package sync_check

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/types"
)


func CallName(call *ssa.CallCommon) string {

	if call.IsInvoke() {
		return call.String()
	}

	switch v := call.Value.(type) {
	case *ssa.Function:
		fn, ok := v.Object().(*types.Func)
		if !ok {
			return ""
		}
		return fn.FullName()
	case *ssa.Builtin:
		return v.Name()
	}
	return ""
}



func Type_sop(inst ssa.Instruction) string { // mind that type checking for atomic operations is in atomic_check.go

	switch true {

	case Is_atomic(inst):        // Is_atomic will recognize any function call including "atomic." as atomic operations
		return Type_atomic(inst) // So we must run Type_atomic() to check if it really belongs to sync/atomic package


	case Is_send_to_channel(inst):
		return "chan_send"

	case Is_receive_to_channel(inst):
		return "chan_receive"

	case Is_select_to_channel(inst):
		return "chan_select"

	case Is_chan_close(inst):
		return "chan_close"

	case Is_mutex_lock(inst):
		return "mutex_lock"

	case Is_mutex_unlock(inst):
		return "mutex_unlock"

	case Is_rwmutex_lock(inst):
		return "rwmutex_lock"

	case Is_rwmutex_unlock(inst):
		return "rwmutex_unlock"

	case Is_rwmutex_rlock(inst):
		return "rwmutex_rlock"

	case Is_rwmutex_runlock(inst):
		return "rwmutex_runlock"

	case Is_waitgroup_add(inst):
		return "waitgroup_add"

	case Is_waitgroup_done(inst):
		return "waitgroup_done"

	case Is_waitgroup_wait(inst):
		return "waitgroup_wait"

	case Is_once_do(inst):
		return "once_do"

	case Is_cond_broadcast(inst):
		return "cond_broadcast"

	case Is_cond_signal(inst):
		return "cond_signal"

	case Is_cond_wait(inst):
		return "cond_wait"

	case Is_pool_get(inst):
		return "pool_get"

	case Is_pool_put(inst):
		return "pool_put"

	case Is_syncmap_delete(inst):
		return "syncmap_delete"

	case Is_syncmap_load(inst):
		return "syncmap_load"

	case Is_syncmap_loadorstore(inst):
		return "syncmap_loadorstore"

	case Is_syncmap_range(inst):
		return "syncmap_range"

	case Is_syncmap_store(inst):
		return "syncmap_store"
	}



	return "other"
}