package sync_check

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/token"
	"strings"
)

func Is_make_channel(inst ssa.Instruction) bool {

	if _,ok := inst.(*ssa.MakeChan); ok {
		return true
	} else {
		return false
	}

}
func Is_alloc_channel(inst ssa.Instruction) bool {
	if strings.HasPrefix(inst.String(),"local chan") || strings.HasPrefix(inst.String(),"local <-chan") || strings.HasPrefix(inst.String(),"local chan<-") {
		return true
	}
	return false
}

func Is_send_to_channel(inst ssa.Instruction) bool {
	_, ok := inst.(*ssa.Send)
	return ok
}

func Is_receive_to_channel(inst ssa.Instruction) bool {
	unop, ok := inst.(*ssa.UnOp)
	if ok {
		if unop.Op == token.ARROW {
			return true
		}
	}
	return false
}

func Is_select_to_channel(inst ssa.Instruction) bool {
	selector, ok := inst.(*ssa.Select)
	if ok {
		// if each case in select is related to a channel
		for _, state := range selector.States {
			if state.Chan == nil {
				fmt.Println("There is a state whose channel is nil in inst:",inst.String())
			}
		}
		return true
	}
	return false
}

func Is_chan_close(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	local_flag := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		if call.IsInvoke() {
			return false
		}
		callName := CallName(call)
		if callName == "close" {
			if strings.Contains(call.Value.Type().String(),"chan") && len(call.Args) == 1 {
				local_flag = true
			} else {
				fmt.Println("Warning: a static call to function close has value whose type.String doesn't contain chan\n" +
					"\tinst:",inst.String(),"\tcall.Value.Type:",call.Value.Type().String())
			}
		}
	}


	return local_flag
}
