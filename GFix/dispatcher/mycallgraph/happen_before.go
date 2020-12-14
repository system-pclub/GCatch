package mycallgraph

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/instruction_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/output"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
)

func Can_inst_happen_before_inst(inst1, inst2 ssa.Instruction) bool {
	if inst1.Parent().String() != inst2.Parent().String() || inst1 == inst2 {
		return false
	}

	// About defer:
	// case 1: inst1 is not in defer, but inst2 is in defer. Then as long as they can be reached by one path, inst1 will
	// happen before inst2
	if instruction_check.Is_defer(inst2) == true && instruction_check.Is_defer(inst1) == false {
		return can_inst_happen_before_inst(inst1,inst2) || can_inst_happen_before_inst(inst2,inst1)
	}

	// case 2: inst1 is in defer and inst2 is not. Then inst1 can't happen before inst2
	if instruction_check.Is_defer(inst2) == false && instruction_check.Is_defer(inst1) == true {
		return false
	}

	// case3: both in defer. If inst2 can be pushed into stack before inst1, inst1 will be executed before inst2
	if instruction_check.Is_defer(inst2) == true && instruction_check.Is_defer(inst1) == true {
		return can_inst_happen_before_inst(inst2, inst1)
	}

	return can_inst_happen_before_inst(inst1, inst2)
}

func can_inst_happen_before_inst(inst1, inst2 ssa.Instruction) bool {

	bb1, bb2 := inst1.Block(), inst2.Block()
	if inst1.Parent() != inst2.Parent() {
		bb2_index := bb2.Index
		bb2 = nil
		// inst1 and inst2 have different parents, but parents' String() is the same
		// This may be a problem in golang's ssa package. We will find a bb matches bb2 in bb1.Parent().Block()
		for _,bb := range bb1.Parent().Blocks {
			if bb.Index == bb2_index {
				bb2 = bb
			}
		}
		if bb2 == nil {
			fmt.Println("--------Warning---------")
			fmt.Println("Still can't find bb2 in bb1.Parent()")
			output.Print_inst_and_location(inst1)
			output.Print_inst_and_location(inst2)
			fmt.Println("-------End Warning----------")
			return false
		}
	}

	if bb1 == bb2 {
		index1,index2 := -1,-1
		for i,inst := range bb1.Instrs {
			if inst.Pos() == inst1.Pos() {
				index1 = i
			}
			if inst.Pos() == inst2.Pos() {
				index2 = i
			}
		}

		if index1 == -1 || index2 == -1 {
			fmt.Println("--------Warning---------")
			fmt.Println("inst1 and inst2 are in the same bb, but can't find them")
			output.Print_inst_and_location(inst1)
			output.Print_inst_and_location(inst2)
			fmt.Println("-------End Warning----------")
			return false
		}

		if index1 < index2 {
			return true
		} else {
			return false
		}
	}

	counter := 0
	empty_bbs := []*ssa.BasicBlock{}
	is_found := recursive_find_target_in_bb_succs(bb1,bb2,empty_bbs,counter)

	return is_found
}

func recursive_find_target_in_bb_succs(bb,target *ssa.BasicBlock, searched_bbs []*ssa.BasicBlock, counter int) bool {
	for _,searched := range searched_bbs {
		if searched == bb {
			return false
		}
	}
	searched_bbs = append(searched_bbs,bb)
	if bb == target {
		return true
	}
	if counter > global.C6A_max_recursive_count {
		return false
	}

	for _,succ := range bb.Succs {
		counter ++
		if recursive_find_target_in_bb_succs(succ,target,searched_bbs,counter) == true {
			return true
		}
	}
	return false
}
