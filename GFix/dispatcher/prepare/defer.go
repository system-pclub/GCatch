package prepare

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/output"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa/ssautil"
)

func Gen_defer_map() map[*ssa.RunDefers][]*ssa.Defer {
	result := make(map[*ssa.RunDefers][]*ssa.Defer)
	for fn,_ := range ssautil.AllFunctions(global.Prog) {
		gen_defer_map_in_fn(fn,result)
	}

	// Delete rundefer that has no defer (these are added by go's ssa, but can be deleted)
	// Reverse the order of defers, because go will call defers that is executed late
	for rundefer, defers := range result {
		if len(defers) == 0 {
			delete(result,rundefer)
		}
		reverse := []*ssa.Defer{}
		for i := len(defers) - 1; i >= 0; i-- {
			reverse = append(reverse,defers[i])
		}
		result[rundefer] = reverse
		defers = nil
	}
	
	//print_defer_map(result)

	return result
}

func print_defer_map(defer_map map[*ssa.RunDefers][]*ssa.Defer) {
	count := 0
	for rundefer, defers := range defer_map {
		fmt.Println("--------NO.",count)
		fmt.Println("----Location of rundefer")
		output.Print_inst_only_location(rundefer)
		fmt.Println("----Location of defer")
		for _,a_defer := range defers {
			output.Print_inst_and_location(a_defer)
			if a_defer.Call.IsInvoke() {
				fmt.Println("\tDefering:",a_defer.Call.Method.String()," of interface ",a_defer.Call.Value.Type().String())
			} else {
				fmt.Println("\tDefering:",a_defer.Call.Value.String())
			}
		}
		count++

		if count%10 == 0 {
			output.Wait_for_input()
		}
	}
}

var map1_gen map[ssa.Instruction] *ssa.Defer
// No need for map2_kill in this task
var map3_before map[ssa.Instruction] []*ssa.Defer
var map4_after map[ssa.Instruction] []*ssa.Defer

func gen_defer_map_in_fn(fn *ssa.Function, result map[*ssa.RunDefers] []*ssa.Defer) {
	map1_gen = make(map[ssa.Instruction] *ssa.Defer)
	map3_before = make(map[ssa.Instruction] []*ssa.Defer)
	map4_after = make(map[ssa.Instruction] []*ssa.Defer)

	prepare_map1(fn)
	if len(map1_gen) == 0 {
		return
	}

	worklist := []ssa.Instruction{find_head_inst(fn)}

	for len(worklist) > 0 {
		inst := worklist[0]

		prev_insts := previous_insts(inst)
		before_defers := previous_defers(prev_insts)
		map3_before[inst] = before_defers

		gen_defer := map1_gen[inst]
		old_map4_defers,is_key_exist := map4_after[inst] // is_key_exist means whether inst is a key in map4_after
		map4_after[inst] = calculate_after(gen_defer,before_defers)

		worklist = delete_inst(worklist,inst)
		if is_key_exist == false || is_map4_changed(old_map4_defers,map4_after[inst]) { // if map4 doesn't have this key
		// before calculate_after or map4 is changed during calculate_after, then we append the next inst to worklist
			worklist = append_next_inst_in_order(worklist,inst)
		}
	}

	for inst,defers := range map3_before {
		if inst_rundefer,ok := inst.(*ssa.RunDefers); ok {
			result[inst_rundefer] = defers
		}
	}

	map1_gen = nil
	map3_before = nil
	map4_after = nil
}

func prepare_map1(fn *ssa.Function) {
	for _,bb := range fn.Blocks {
		for _,inst := range bb.Instrs {
			inst_as_defer, ok := inst.(*ssa.Defer)
			if ok {
				map1_gen[inst] = inst_as_defer
			}
		}
	}
}

func find_head_inst(fn *ssa.Function) ssa.Instruction {
	for _, bb := range fn.Blocks {
		for _, inst := range bb.Instrs {
			return inst
		}
	}
	return nil
}

func previous_insts(inst ssa.Instruction) []ssa.Instruction {
	result := []ssa.Instruction{}
	if inst == find_head_inst(inst.Parent()) { //case1: inst is the head of whole function
		return result
	}

	bb := inst.Block()
	if inst == bb.Instrs[0] { //case2: inst is the head of a bb (but not the head of whole function)
		for _, pred_bb := range bb.Preds {
			result = append(result, last_inst_bb(pred_bb))
		}
		return result
	}

	//case3: inst is not the head of a bb
	for index,_ := range bb.Instrs {
		if inst == bb.Instrs[index] {
			result = append(result, bb.Instrs[index-1])
			return result
		}
	}

	return result
}

func last_inst_bb(bb *ssa.BasicBlock) ssa.Instruction {
	insts := bb.Instrs
	if len(insts) == 0 {
		return nil
	}

	last_inst := insts[len(insts) - 1]
	return last_inst
}

func previous_defers(prev_insts []ssa.Instruction) []*ssa.Defer {
	result := []*ssa.Defer{}
	if len(prev_insts) == 0 { //This is the first time this whole worklist loop is invoked
		// Do nothing
	} else if len(prev_insts) == 1 { // Inherit the defers from previous inst
		prev_defers,ok := map4_after[prev_insts[0]]
		if ok { // if !ok, then map4_after is empty, don't store
			result = prev_defers
		}
	} else {					//Union the map4_after of all previous inst
		map_prev_defer := make(map[*ssa.Defer]struct{})
		for _,prev_inst := range prev_insts {
			prev_defers,ok := map4_after[prev_inst]
			if !ok { // if !ok, then map4_after is empty, don't store
				continue
			}
			for _,a_defer := range prev_defers {
				map_prev_defer[a_defer] = struct{}{}
			}
		}
		for a_defer,_ := range map_prev_defer {
			result = append(result,a_defer)
		}
	}
	return result
}

func calculate_after(gen_defer *ssa.Defer, before_defers []*ssa.Defer) []*ssa.Defer {
	flag_in_slice := false
	for _,before_defer := range before_defers {
		if before_defer == gen_defer {
			flag_in_slice = true
			break
		}
	}

	if flag_in_slice || gen_defer == nil {
		return before_defers
	} else {
		return append(before_defers,gen_defer)
	}
}

func delete_inst(todo []ssa.Instruction, inst ssa.Instruction) (result []ssa.Instruction) {

	for _,todo_inst := range todo {
		if todo_inst != inst {
			result = append(result,todo_inst)
		}
	}

	return result
}

func is_map4_changed(old_defers,new_defers []*ssa.Defer) bool {
	if len(old_defers) != len(new_defers) {
		return true
	}

	if len(old_defers) == 0 {
		return false
	}

	// Now two slice are the same. If and only if one element of new_defers is not in old_defers, the map is changed
	old_map := make(map[*ssa.Defer]struct{})
	for _,old := range old_defers {
		old_map[old] = struct{}{}
	}

	for _,a_new := range new_defers {
		if _,is_new_in_old := old_map[a_new]; is_new_in_old == false {
			return true
		}
	}

	return false
}

func append_next_inst_in_order(todo []ssa.Instruction, inst ssa.Instruction) (result []ssa.Instruction) {
	//if the next_inst is the head of a BB, append it to the end of todo_list
	//if the next_inst is not the head of a BB, append it to the beginning of todo_list

	next_insts := calc_next_insts(inst)
	head_of_bb_insts := []ssa.Instruction{}
	not_head_of_bb_insts := []ssa.Instruction{}
	for _,next_inst := range next_insts {
		if is_inst_in_slice(next_inst,todo) == true {
			continue
		}
		if bb := (next_inst).Block(); bb.Instrs[0] == next_inst { //next_inst is the head of a bb
			head_of_bb_insts = append(head_of_bb_insts,next_inst)
		} else {
			not_head_of_bb_insts = append(not_head_of_bb_insts,next_inst)
		}
	}

	for _,not_head_inst := range not_head_of_bb_insts {
		result = append(result,not_head_inst)
	}
	for _,old_inst := range todo {
		result = append(result,old_inst)
	}
	for _,head_inst := range head_of_bb_insts {
		result = append(result,head_inst)
	}

	return
}

func calc_next_insts(inst ssa.Instruction) []ssa.Instruction {
	if Is_inst_end_of_fn(inst) { //case1: inst is the end of whole function. May be return or panic
		return []ssa.Instruction{}
	}

	for _,bb := range inst.Parent().Blocks {	//case2: inst is the end of a bb (but not the end of whole function)
		if inst == last_inst_bb(bb) {
			var result []ssa.Instruction
			for _,succ_bb := range bb.Succs {
				result = append(result,first_inst_bb(succ_bb))
			}
			return result
		}
	}

	for _,bb := range inst.Parent().Blocks{ //case3: inst is not the end of a bb
		for index,_ := range bb.Instrs {
			if inst == bb.Instrs[index] {
				var result []ssa.Instruction
				result = append(result, bb.Instrs[index+1])
				return result
			}
		}
	}

	return nil
}

func Is_inst_end_of_fn(inst ssa.Instruction) bool {
	fn_parent := inst.Parent()
	all_bbs := fn_parent.Blocks
	if len(all_bbs) == 0 {
		return false
	}

	for _,bb := range all_bbs {
		if len(bb.Succs) == 0 {
			last_inst := last_inst_bb(bb)
			if inst == last_inst {
				return true
			}
		}
	}

	return false
}

func first_inst_bb(bb *ssa.BasicBlock) ssa.Instruction {
	insts := bb.Instrs
	if len(insts) == 0 {
		return nil
	}

	first_inst := insts[0]
	return first_inst
}

func is_inst_in_slice(target ssa.Instruction, slice []ssa.Instruction) bool {
	for _,inst := range slice {
		if inst == target {
			return true
		}
	}
	return false
}
