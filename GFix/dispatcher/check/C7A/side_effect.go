package C7A

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/sync_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"reflect"
	"strings"
)

func checkHasSideEffectOnBB(ins ssa.Instruction, checkChannelOperations bool) bool {
	println("[DEBUG] Checking side effect on BB:")
	//printSSAByBB(ins.Block())
	started := false
	for _, insIt := range ins.Block().Instrs {
		if insIt == ins {
			started = true
			fmt.Println("started from: ", ins)
			continue
		}
		if started {
			printSSAIns(insIt)
			if checkSideEffectIns(insIt) {
				return true
			}
			if checkChannelOperations && (sync_check.Is_receive_to_channel(insIt) || sync_check.Is_send_to_channel(insIt)) {
				return true
			}
		}
	}
	println("Finished side effect checking.")
	return false
}

func checkHasSideEffect(entry *ssa.BasicBlock, checkChannelOperations bool) bool {
	finder := NewSuccBasicBlockFinder()
	bbs := finder.Analyze(entry)
	for _, bb := range bbs {
		if bb == entry {
			continue
		}
		fmt.Println("[DEBUG] Checking side effect on BB ", bb.Comment, bb.Index)
		printSSAByBB(bb)
		for _, ins := range bb.Instrs {
			b := checkSideEffectIns(ins)
			if b {
				return true
			}
			if checkChannelOperations && (sync_check.Is_receive_to_channel(ins) || sync_check.Is_send_to_channel(ins)) {
				return true
			}
		}
	}
	return false
}

func checkSideEffectIns(ins ssa.Instruction) bool {
	instCall, ok := ins.(*ssa.Call)
	if ok {
		funcName := sync_check.CallName(instCall.Common())
		if strings.Contains(funcName, "Error") ||
			strings.Contains(funcName, "error") ||
			strings.Contains(funcName, "Log") ||
			strings.Contains(funcName, "(*sync/atomic.Value)") ||
			strings.Contains(funcName, "Leave") {
			println("[DEBUG] function name is: " + funcName + ", which matchs side effect funcs.")
			return true
		} else {
			println("[DEBUG] function name is: " + funcName + ", which doesn't match side effect funcs.")
		}
	} else {
		fmt.Println(reflect.TypeOf(ins))
	}
	return false
}

func checkDeferSideEffects(fn *ssa.Function) bool {
	for _, bb := range fn.Blocks {
		for _, inst := range bb.Instrs {
			inst_as_defer, ok := inst.(*ssa.Defer)
			if ok {
				funcName := sync_check.CallName(&inst_as_defer.Call)
				if strings.Contains(funcName, "Close") {
					return true
				} else {
					println("[DEBUG] funcname is " + funcName + " and is not a side effect function.")
				}
			}
		}
	}
	return false
}
