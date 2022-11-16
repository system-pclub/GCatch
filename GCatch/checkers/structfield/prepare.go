package structfield

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/util"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

type MyStruct struct {
	Name  string
	Field map[string]string
}

var searched_bb []*ssa.BasicBlock

var map1_gen map[ssa.Instruction]string //example: "mypkg.mytype.mu_mutex","mypkg.mytype.rwmu_rwmutexR","mypkg.mytype.rwmu_rwmutexW"
var map2_kill map[ssa.Instruction]string
var map3_before map[ssa.Instruction][]string //example: ["mypkg.mytype.mu_mutex","mypkg.mytype.rwmu_rwmutexR","mypkg.mytype.rwmu_rwmutexW"]
var map4_after map[ssa.Instruction][]string

func loop_pkg_C3() {

	vecAllMethods := List_all_methods()

	for _, pkg := range config.Prog.AllPackages() { //loop all packages
		if pkg == nil {
			continue
		}

		//Skip builtin packages, vendor packages. Test functions are automatically skipped. Include packages in "include"
		if config.IsPathIncluded(pkg.Pkg.Path()) {
		} else {
			continue
		}

		for mem_name, mem := range pkg.Members { //loop through all members; the member may be a func or a type; if it is type, loop through all its methods

			//check if this member is a type
			mem_as_type := pkg.Type(mem_name)
			if mem_as_type != nil {
				//This member is a type

				for _, method := range vecAllMethods {
					if method == nil || method.Pkg != pkg {
						continue
					}
					method_prefix_1 := "(*" + pkg.Pkg.Path() + "." + mem_name + ")."
					method_prefix_2 := "(" + pkg.Pkg.Path() + "." + mem_name + ")."
					ptr_to_type := types.NewPointer(mem.Type().Underlying())

					if strings.Contains(method.String(), method_prefix_1) {
						//this function is a method of mem_as_type, and it is in pkg
						inside_method(method, ptr_to_type)
					} else if strings.Contains(method.String(), method_prefix_2) {
						//this function is a method of mem_as_type, and it is in pkg
						inside_method(method, mem.Type())
					}
				}
			}

		} // end of member loop
	} //end of package loop
}

func inside_method(fn *ssa.Function, type_receiver types.Type) {

	if fn.Signature.Recv() == nil {
		//fmt.Println("\tfn.Signature.Recv() == nil") //This is an anonymous function inside our interested method
		return
	}

	receiver_name := fn.Params[0].Name()

	if fn.Blocks == nil { //meaning this is external function. You will see a lot of them if you use Ssa_build_packages
		return
	}

	loop_BB(*fn, receiver_name) // loop through all BB in fn

}

func loop_BB(fn ssa.Function, receiver_name string) {

	//fmt.Println("------------Func:",fn.RelString(fn.Pkg.Pkg))

	for _, bb := range fn.Blocks { // loop all BBs

		insts := bb.Instrs

		for _, inst := range insts { // loop all instructions

			if is_inst_interesting(inst, receiver_name) == false {
				continue
			}

			ptr_struct, field_name := find_struct_ptr(inst)
			if ptr_struct == nil {
				continue
			}
			list_alive_mutexs := Alive_mutexs(inst)
			ptr_struct.Field[field_name][inst] = list_alive_mutexs

		} // end of instruction loop
	} //end of BB loop
}

func is_inst_interesting(target_inst ssa.Instruction, receiver_name string) bool {

	if instinfo.IsDefer(target_inst) { // if target_inst is a defer, it may not be possible to determine its previous locks, because there may be deferred lock/unlock
		return false //we give up this target
	}
	inst_as_FA, ok := target_inst.(*ssa.FieldAddr)
	if !ok {
		return false
	}
	field_name := inst_as_FA.X.Type().Underlying().(*types.Pointer).Elem().Underlying().(*types.Struct).Field(inst_as_FA.Field).Name()

	if inspect_stmt(target_inst, receiver_name, field_name) == false {
		return false
	}

	return true
}

func inspect_stmt(inst ssa.Instruction, receiver_name string, field_name string) bool {
	inst_position := (config.Prog.Fset).Position(inst.Pos())
	filename := inst_position.Filename
	line := inst_position.Line
	flag_see_above_line := false
	if line < 1 {
		index_in_bb := -1
		insts_in_bb := inst.Block().Instrs
		for index, bb_inst := range insts_in_bb {
			if bb_inst == inst {
				index_in_bb = index
			}
		}
		for i := index_in_bb + 1; i < len(insts_in_bb); i++ { //If we can find an bb_inst that is after inst and has line number, we will use bb_inst's line number
			// and filename, and we will also inspect the line before this line
			position := (config.Prog.Fset).Position(insts_in_bb[i].Pos())
			if position.Line > 0 {
				line = position.Line
				filename = position.Filename
				flag_see_above_line = true
				break
			}
		}

		if line < 1 {
			return false
		}
	}
	str_same_line, err := util.ReadFileLine(filename, line)
	if err != nil {
		fmt.Println("Error: during read file:", filename, "\tline:", line, "\tfor inst:", inst)
		panic(err)
	}

	if strings.Contains(str_same_line, ".Lock()") || strings.Contains(str_same_line, ".Unlock()") ||
		strings.Contains(str_same_line, ".RLock()") || strings.Contains(str_same_line, ".RUnlock()") {
		return false
	}

	if is_field_calling(str_same_line, field_name) == true {
		return false
	}

	if strings.Contains(str_same_line, "make(") || strings.Contains(str_same_line, "new(") {
		return false
	}

	if strings.Contains(str_same_line, receiver_name+".") {
		return true
	}

	if flag_see_above_line == true {
		line = line - 1
		str_same_line, err := util.ReadFileLine(filename, line)
		if err != nil {
			fmt.Println("Error: during read file:", filename, "\tline:", line, "\tfor inst:", inst)
			panic(err)
		}

		if strings.Contains(str_same_line, "make(") || strings.Contains(str_same_line, "new(") {
			return false
		}

		if strings.Contains(str_same_line, receiver_name+".") {
			return true
		}
	}

	return false
}

func find_struct_ptr(target_inst ssa.Instruction) (*C3_struct, string) {
	var target_struct_ptr *C3_struct

	target_FA, _ := target_inst.(*ssa.FieldAddr)
	target_struct_ptr_name := target_FA.X.Type().String()
	_, ok := target_FA.X.Type().Underlying().(*types.Pointer).Elem().Underlying().(*types.Struct).Field(target_FA.Field).Type().(*types.Chan)
	if ok {
		return nil, ""
	}

	target_field_name := target_FA.X.Type().Underlying().(*types.Pointer).Elem().Underlying().(*types.Struct).Field(target_FA.Field).Name() // see https://godoc.org/golang.org/x/tools/go/ssa#FieldAddr

outer:
	for _, stru := range C3_all_structs {
		if "*"+stru.Name == target_struct_ptr_name {
			for field_name, _ := range stru.Field {
				if field_name == target_field_name {
					target_struct_ptr = stru

					break outer
				}
			}
		}

	}

	return target_struct_ptr, target_field_name
}

func is_field_calling(str_same_line string, field_name string) bool {
	index_field := strings.Index(str_same_line, field_name)
	if index_field == -1 {
		return false
	}
	index_call := strings.Index(str_same_line[index_field:], "(")
	if index_call == -1 {
		return false
	}
	for i := index_field; i < index_call; i++ {
		if str_same_line[i] == ' ' {
			return false
		}
	}
	return true

}

func List_all_struct(prog *ssa.Program) []*MyStruct {
	all_structs := *new([]*MyStruct)

	for _, pkg := range prog.AllPackages() { //loop all packages
		if pkg == nil {
			continue
		}

		if config.IsPathIncluded(pkg.Pkg.Path()) == false {
			continue
		}

		for mem_name, _ := range pkg.Members { //loop through all members; the member may be a func, a type, etc

			//check if this member is a type
			mem_as_type := pkg.Type(mem_name)
			if mem_as_type != nil {

				//check if this member is in our interested path
				if config.IsPathIncluded(mem_as_type.String()) == false {
					continue
				}

				struct_name := mem_as_type.String()
				fields_str := mem_as_type.Object().Type().Underlying().String()

				if !strings.HasPrefix(fields_str, "struct{") || fields_str == "struct{}" {
					continue
				}
				fields_str = strings.Replace(fields_str, "struct{", "", 1)
				fields_str = fields_str[:len(fields_str)-1] //delete the last char, which is "}"
				fields := strings.Split(fields_str, "; ")
				if len(fields) == 0 {
					continue
				} else {
					str := strings.ReplaceAll(fields[0], " ", "")
					if len(str) == 0 {
						continue
					}
				}

				struct_field := make(map[string]string)
				for _, field := range fields {
					field_element := strings.Split(field, " ")
					var field_name, field_type string
					if len(field_element) == 1 { //this is an anonymous field
						field_name = field_element[0]
						last_dot_index := strings.LastIndex(field_name, ".") // from "*github.com/coreos/etcd/mvcc.store", we only want "store"
						field_name = field_name[last_dot_index+1:]
						//fmt.Println("Anonymous field:",field_element[0],"\trefined:",field_name,"\tstruct.Name:",struct_name)

						field_type = field_element[0]
						if field_element[0] == "chan" && len(field_element) > 1 {
							field_type = "chan " + field_element[1]
						}
					} else {
						field_name = field_element[0]
						field_type = field_element[1]
						if field_element[1] == "chan" && len(field_element) > 2 {
							field_type = "chan " + field_element[2]
						}
					}
					struct_field[field_name] = field_type
				}
				new_struct_ptr := &MyStruct{
					Name:  struct_name,
					Field: struct_field,
				}

				all_structs = append(all_structs, new_struct_ptr)

			}
		}
	}

	return all_structs
}

func List_all_methods() []*ssa.Function {
	methodset := *new([]*ssa.Function)

	fns_in_prog := ssautil.AllFunctions(config.Prog)
	for fn_in_prog, _ := range fns_in_prog { // a cumbersome loop, looping through all functions in the program
		method_prefix := ")."
		var str string
		if fn_in_prog.Pkg == nil {
			str = fn_in_prog.String()

		} else {
			if config.IsPathIncluded(fn_in_prog.Pkg.Pkg.Path()) == false {
				continue
			}
			str = fn_in_prog.RelString(fn_in_prog.Pkg.Pkg)
		}
		if strings.Contains(str, method_prefix) {
			//this function is a method of mem_as_type, and it is in pkg
			methodset = append(methodset, fn_in_prog)
		}
	}

	var result []*ssa.Function = *new([]*ssa.Function)
	for _, method := range methodset {
		if method.Pkg != nil && method.Synthetic == "" {
			result = append(result, method)
		}
	}

	return result
}

func Alive_mutexs(target_inst ssa.Instruction) []string {

	//There are 4 maps for every inst in target_inst.Parent(); map1_gen[inst] lists the mutex/rwmutex generated by this inst;
	//map2_kill[inst] lists the mutex/rwmutex killed by this inst; map3_before[inst] lists mutexes/rwmutexes that haven't been unlocked before this inst
	//map4_after[inst] lists mutexes/rwmutexes that haven't been unlocked after this inst
	map1_gen = make(map[ssa.Instruction]string)
	map2_kill = make(map[ssa.Instruction]string)
	map3_before = make(map[ssa.Instruction][]string)
	map4_after = make(map[ssa.Instruction][]string)

	//fill map1 and map2
	prepare_map1_map2(target_inst, false)
	flag_empty_map1 := true
	for _, lock_name := range map1_gen {
		if lock_name != "" {
			flag_empty_map1 = false
			break
		}
	}
	if flag_empty_map1 == true {
		return []string{}
	}

	//fill map3 and map4, let
	prepare_map3_map4(target_inst)
	ptr_head_inst := find_head_inst(target_inst)
	todo := *new([]*ssa.Instruction)
	todo = append(todo, ptr_head_inst)

	for len(todo) > 0 {
		inst := *todo[0]

		var previous_insts []*ssa.Instruction
		previous_insts = calc_previous_insts(inst)
		var before []string
		if len(previous_insts) == 0 { //This is the first time this loop is invoked
			before = []string{}

		} else if len(previous_insts) == 1 {
			before = map4_after[*previous_insts[0]]
		} else { //Union the map4_after of all previous inst
			before = union_prev_inst(previous_insts)
		}
		map3_before[inst] = before

		gen := map1_gen[inst]
		kill := map2_kill[inst]
		after := calc_after(before, gen, kill)

		todo = delete_todo(todo, inst)
		if string_slice_equal(map4_after[inst], after) == false {
			map4_after[inst] = after
			todo = append_in_order_todo(todo, inst) //The order is: if the following inst is the beginning of a BB, append it to the end of todo_list
			//		if the following inst is not the beginning of a BB, append it to the beginning of todo_list
		}

	}

	return map3_before[target_inst]

}

// Alive_unlock is an inverse version of alive_mutexs. It's aimed to find functions that contains a mutex that is only Unlocked, not Locked
// this function has limitations: it can only return one mutex that is only Unlocked; it will fail when the function is like "mu.Unlock(); mu.Lock(); mu.Unlock()"
func Alive_unlock(target_inst ssa.Instruction) string {

	map1_gen = make(map[ssa.Instruction]string)
	map2_kill = make(map[ssa.Instruction]string)
	map3_before = make(map[ssa.Instruction][]string)
	map4_after = make(map[ssa.Instruction][]string)

	//fill map1 and map2
	prepare_map1_map2(target_inst, false)
	flag_empty_map1 := true
	for _, lock_name := range map1_gen {
		if lock_name != "" {
			flag_empty_map1 = false
			break
		}
	}
	if flag_empty_map1 == true {
		return ""
	}

	//fill map3 and map4
	prepare_map3_map4(target_inst)
	ptr_head_inst := find_head_inst(target_inst)
	todo := *new([]*ssa.Instruction)
	todo = append(todo, ptr_head_inst)

	for len(todo) > 0 {
		inst := *todo[0]

		var previous_insts []*ssa.Instruction
		previous_insts = calc_previous_insts(inst)
		var before []string
		if len(previous_insts) == 0 { //This is the first time this loop is invoked
			before = []string{}

		} else if len(previous_insts) == 1 {
			before = map4_after[*previous_insts[0]]
		} else { //Union the map4_after of all previous inst
			before = union_prev_inst(previous_insts)
		}
		map3_before[inst] = before

		gen := map1_gen[inst]
		kill := map2_kill[inst]
		after := calc_after(before, gen, kill)
		seperate_unlock_mutex, find_seperate_unlock := find_separate_unlock(before, kill)
		if find_seperate_unlock == true {
			return seperate_unlock_mutex
		}

		todo = delete_todo(todo, inst)
		if string_slice_equal(map4_after[inst], after) == false {
			map4_after[inst] = after
			todo = append_in_order_todo(todo, inst) //The order is: if the following inst is the beginning of a BB, append it to the end of todo_list
			//		if the following inst is not the beginning of a BB, append it to the beginning of todo_list
		}

	}

	return ""

}

func union_prev_inst(previous_insts []*ssa.Instruction) []string {
	var result []string
	for _, previous_inst := range previous_insts {
		for _, str := range map4_after[*previous_inst] {
			if is_str_in_slice(str, result) == false && str != "init" { //"init" is the initial value in map4_after, see func prepare_map3_map4()
				result = append(result, str)
			}
		}
	}
	return result
}

func is_str_in_slice(str string, slice []string) bool {
	for _, slice_str := range slice {
		if slice_str == str {
			return true
		}
	}
	return false
}

func append_in_order_todo(todo []*ssa.Instruction, inst ssa.Instruction) (result []*ssa.Instruction) {
	//if the next_inst is the head of a BB, append it to the end of todo_list
	//if the next_inst is not the head of a BB, append it to the beginning of todo_list

	next_insts := calc_next_insts(inst)
	head_of_bb_insts := []*ssa.Instruction{}
	not_head_of_bb_insts := []*ssa.Instruction{}
	for _, next_inst := range next_insts {
		if is_inst_in_slice(*next_inst, todo) == true {
			continue
		}
		if bb := (*next_inst).Block(); bb.Instrs[0] == *next_inst { //next_inst is the head of a bb
			head_of_bb_insts = append(head_of_bb_insts, next_inst)
		} else {
			not_head_of_bb_insts = append(not_head_of_bb_insts, next_inst)
		}
	}

	for _, not_head_inst := range not_head_of_bb_insts {
		result = append(result, not_head_inst)
	}
	for _, old_inst := range todo {
		result = append(result, old_inst)
	}
	for _, head_inst := range head_of_bb_insts {
		result = append(result, head_inst)
	}

	return
}

func is_inst_in_slice(inst ssa.Instruction, slice []*ssa.Instruction) bool {
	for _, slice_inst := range slice {
		if *slice_inst == inst {
			return true
		}
	}
	return false
}

func calc_next_insts(inst ssa.Instruction) []*ssa.Instruction {
	if Is_inst_end_of_fn(inst) { //case1: inst is the end of whole function
		return []*ssa.Instruction{}
	}

	for _, bb := range inst.Parent().Blocks { //case2: inst is the end of a bb (but not the end of whole function)
		if inst == *last_inst_bb(*bb) {
			var result []*ssa.Instruction
			for _, succ_bb := range bb.Succs {
				result = append(result, first_inst_bb(*succ_bb))
			}
			return result
		}
	}

	for _, bb := range inst.Parent().Blocks { //case3: inst is not the end of a bb
		for index, _ := range bb.Instrs {
			if inst == bb.Instrs[index] {
				var result []*ssa.Instruction
				result = append(result, &bb.Instrs[index+1])
				return result
			}
		}
	}

	fmt.Println("Error when calculating previous insts for inst:", inst)
	panic(inst)

}

func Is_inst_end_of_fn(inst ssa.Instruction) bool {
	fn_parent := *inst.Parent()
	all_bbs := fn_parent.Blocks
	if len(all_bbs) == 0 {
		return false
	}

	for _, bb := range all_bbs {
		if len(bb.Succs) == 0 {
			last_inst := last_inst_bb(*bb)
			if inst == *last_inst {
				return true
			}
		}
	}

	return false
}

func string_slice_equal(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for _, str1 := range slice1 {
		flag_found := false
		for _, str2 := range slice2 {
			if str1 == str2 {
				flag_found = true
			}
		}
		if flag_found == false {
			return false
		}
	}
	return true
}

func delete_todo(todo []*ssa.Instruction, inst ssa.Instruction) []*ssa.Instruction {
	var result []*ssa.Instruction
	for _, todo_inst := range todo {
		if *todo_inst != inst {
			result = append(result, todo_inst)
		}
	}

	return result
}

func calc_after(before []string, gen string, kill string) []string {

	var after []string
	after = before
	//before + gen
	if gen != "" {
		after = add_mutex_to_mutexs(after, gen)
	}

	if kill != "" {
		after = remove_mutex_from_mutexs(after, kill)
	}

	return after

}

func find_separate_unlock(before []string, kill string) (string, bool) {

	var has_separate_unlock bool = false
	if kill != "" {
		var find_kill bool = false
		for _, str := range before {
			if str == kill {
				find_kill = true
				break
			}
		}
		if find_kill == false {
			has_separate_unlock = true
			return kill, has_separate_unlock
		}
	}

	return "", has_separate_unlock

}

func remove_mutex_from_mutexs(target_slice []string, delete string) []string {
	var result []string
	for _, str := range target_slice {
		if str != string(delete) {
			result = append(result, str)
		}
	}

	return result
}

func add_mutex_to_mutexs(target_slice []string, add string) []string {
	for _, str := range target_slice {
		if str == add {
			return target_slice
		}
	}
	result := append(target_slice, add)
	return result
}

func calc_previous_insts(inst ssa.Instruction) []*ssa.Instruction {
	if is_inst_head_of_fn(inst) { //case1: inst is the head of whole function
		return []*ssa.Instruction{}
	}

	bb := inst.Block()
	if inst == bb.Instrs[0] { //case2: inst is the head of a bb (but not the head of whole function)
		var result []*ssa.Instruction
		for _, pred_bb := range bb.Preds {
			result = append(result, last_inst_bb(*pred_bb))
		}
		return result
	}

	//case3: inst is not the head of a bb
	for index, _ := range bb.Instrs {
		if inst == bb.Instrs[index] {
			var result []*ssa.Instruction
			result = append(result, &bb.Instrs[index-1])
			return result
		}
	}

	fmt.Println("Error when calculating previous insts for inst:", inst)
	return []*ssa.Instruction{}

}

func last_inst_bb(bb ssa.BasicBlock) *ssa.Instruction {
	insts := bb.Instrs
	if len(insts) == 0 {
		return nil
	}

	last_inst := insts[len(insts)-1]
	return &last_inst
}

func first_inst_bb(bb ssa.BasicBlock) *ssa.Instruction {
	insts := bb.Instrs
	if len(insts) == 0 {
		return nil
	}

	first_inst := insts[0]
	return &first_inst
}

func is_inst_head_of_fn(inst ssa.Instruction) bool {
	bb := *inst.Parent().Blocks[0]
	return bb.Instrs[0] == inst
}

// if is_brutal == true, then when we decide whether a callee is Lock/Unlock, we use case_insensive_contains
func prepare_map1_map2(target_inst ssa.Instruction, is_brutal bool) {
	target_position := (config.Prog.Fset).Position(target_inst.Pos())
	_ = target_position
	all_bbs := target_inst.Parent().Blocks
	for _, bb := range all_bbs {
		for _, inst := range bb.Instrs {

			//inst_position := (config.Prog.Fset).Position(inst.Pos())
			var primitive_locked string
			primitive_locked = ""
			map1_gen[inst] = primitive_locked

			var primitive_unlocked string
			primitive_unlocked = ""
			map2_kill[inst] = primitive_unlocked

			if instinfo.IsDefer(inst) {

			} else if instinfo.IsMutexLock(inst) || instinfo.IsRwmutexLock(inst) || instinfo.IsRwmutexRlock(inst) ||
				((is_brutal == false && Is_self_lock(inst)) || (is_brutal == true && Is_self_lock_brutal(inst))) {
				var primitive_locked string
				if instinfo.IsMutexLock(inst) {
					primitive_locked = string(Find_stmt_match(inst) + "_mutex")

				} else if instinfo.IsRwmutexLock(inst) {
					primitive_locked = string(Find_stmt_match(inst) + "_rwmutexW")

				} else if instinfo.IsRwmutexRlock(inst) {
					primitive_locked = string(Find_stmt_match(inst) + "_rwmutexR")
				} else {

					primitive_locked = string(Find_stmt_match(inst) + "_unknown")
				}
				map1_gen[inst] = primitive_locked

			} else if instinfo.IsMutexUnlock(inst) || instinfo.IsRwmutexUnlock(inst) || instinfo.IsRwmutexRunlock(inst) ||
				((is_brutal == false && Is_self_unlock(inst)) || (is_brutal == true && Is_self_unlock_brutal(inst))) {
				primitive_locked = ""
				map1_gen[inst] = primitive_locked

				var primitive_unlocked string
				if instinfo.IsMutexUnlock(inst) {
					primitive_unlocked = string(Find_stmt_match(inst) + "_mutex")

				} else if instinfo.IsRwmutexUnlock(inst) {
					primitive_unlocked = string(Find_stmt_match(inst) + "_rwmutexW")

				} else if instinfo.IsRwmutexRunlock(inst) {
					primitive_unlocked = string(Find_stmt_match(inst) + "_rwmutexR")
				} else {
					primitive_unlocked = string(Find_stmt_match(inst) + "_unknown")
				}

				map2_kill[inst] = primitive_unlocked
			} else if flag_is_go_unlock, primitive_unlocked := is_go_unlock(inst, is_brutal); flag_is_go_unlock == true {
				primitive_locked = ""
				map1_gen[inst] = primitive_locked
				map2_kill[inst] = primitive_unlocked
			}
		}
	}
}

func is_go_unlock(inst ssa.Instruction, is_brutal bool) (bool, string) {
	inst_as_go, ok := inst.(*ssa.Go)
	if !ok {
		return false, ""
	}
	callCommon := inst_as_go.Call
	var go_fn *ssa.Function = nil
	if callCommon.IsInvoke() == true { //If this is a call to method, we can't track
		return false, ""
	} else {
		callCommon_fn, ok := callCommon.Value.(*ssa.Function) //callCommon.Value can be *ssa.Function or *ssa.Closure or other types that we don't care
		if ok {
			go_fn = callCommon_fn
		} else {
			closure, ok := callCommon.Value.(*ssa.MakeClosure)
			if ok {
				closure_fn, ok := closure.Fn.(*ssa.Function)
				if ok {
					go_fn = closure_fn
				}
			}
		}
	}
	if go_fn == nil {
		return false, ""
	}

	//Now go_fn is the function being called by go statement
	go_fn.DomPreorder()
	for _, bb := range go_fn.Blocks {
		for _, inst := range bb.Instrs {
			if instinfo.IsMutexUnlock(inst) || instinfo.IsRwmutexUnlock(inst) || instinfo.IsRwmutexRunlock(inst) ||
				((is_brutal == false && Is_self_unlock(inst)) || (is_brutal == true && Is_self_unlock_brutal(inst))) {
				primitive_unlocked := ""
				if instinfo.IsMutexUnlock(inst) {
					primitive_unlocked = string(Find_stmt_match(inst) + "_mutex")

				} else if instinfo.IsRwmutexUnlock(inst) {
					primitive_unlocked = string(Find_stmt_match(inst) + "_rwmutexW")

				} else if instinfo.IsRwmutexRunlock(inst) {
					primitive_unlocked = string(Find_stmt_match(inst) + "_rwmutexR")
				} else {
					primitive_unlocked = string(Find_stmt_match(inst) + "_unknown")
				}
				return true, primitive_unlocked
			}
		}
	}
	return false, ""
}

func Find_stmt_match(inst ssa.Instruction) string {
	inst_position := (config.Prog.Fset).Position(inst.Pos())
	filename := inst_position.Filename
	line := inst_position.Line
	if line < 1 {
		return ""
	}

	str_same_line, err := util.ReadFileLine(filename, line)
	if err != nil {
		fmt.Println("Error: during read file:", filename, "\tline:", line, "\tfor inst:", inst)
		return ""
	}

	if index_comment := strings.Index(str_same_line, "//"); index_comment > -1 {
		str_same_line = str_same_line[:index_comment]
	}
	var mutex_name string
	if index_Lock := strings.Index(str_same_line, "Lock"); index_Lock > -1 {
		mutex_name = str_same_line[:index_Lock]
	} else if index_Unlock := strings.Index(str_same_line, "Unlock"); index_Unlock > -1 {
		mutex_name = str_same_line[:index_Unlock]
	} else if index_RUnlock := strings.Index(str_same_line, "RUnlock"); index_RUnlock > -1 {
		mutex_name = str_same_line[:index_RUnlock]
	} else if index_last_dot := strings.LastIndex(str_same_line, "."); index_last_dot > -1 {
		mutex_name = str_same_line[:index_last_dot]
	}

	if strings.Contains(mutex_name, "defer") {
		str_split := strings.Split(mutex_name, " ")
		mutex_name = str_split[len(str_split)-1]
	}
	mutex_name = strings.TrimSpace(mutex_name)

	return mutex_name

}

func prepare_map3_map4(target_inst ssa.Instruction) {
	all_bbs := target_inst.Parent().Blocks
	for _, bb := range all_bbs {
		for _, inst := range bb.Instrs {

			var s []string = []string{}
			map3_before[inst] = s
			var s2 = []string{"init"}
			map4_after[inst] = s2
		}
	}
}

func find_head_inst(target_inst ssa.Instruction) *ssa.Instruction {
	all_bbs := target_inst.Parent().Blocks
	for _, bb := range all_bbs {
		for _, inst := range bb.Instrs {
			return &inst
		}
	}
	fmt.Println("Error: can't find head_inst for target_inst:", target_inst)
	return nil
}

// If target_inst is a ssa.Call and it is calling a method or function containing string "lock", we believe it is calling a function containing a lock
func Is_self_lock_brutal(target_inst ssa.Instruction) bool {
	inst_as_call, ok := target_inst.(*ssa.Call)
	if !ok {
		return false
	}
	if inst_as_call.Call.IsInvoke() == true {
		if case_insensitive_contains(inst_as_call.Call.Method.Name(), "lock") {
			return true
		}
	} else {
		callee, ok := inst_as_call.Call.Value.(*ssa.Function)
		if ok {
			if case_insensitive_contains(callee.Name(), "lock") {
				return true
			}
		}
	}
	return false
}

func case_insensitive_contains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

// If target_inst is a ssa.Call and it is calling a method naming "*.Lock()", we believe it is calling a function containing a lock
//
//	If it is calling a function that has a Mis-Unlock behavior, we believe it is calling a function containing a lock
func Is_self_lock(target_inst ssa.Instruction) bool {
	inst_as_call, ok := target_inst.(*ssa.Call)
	if !ok {
		return false
	}
	if inst_as_call.Call.IsInvoke() == true {
		if inst_as_call.Call.Method.Name() == "Lock" {
			return true
		}
	} else {
		callee, ok := inst_as_call.Call.Value.(*ssa.Function)
		if !ok {
			return false
		}
		if callee.Name() != "Lock" {
			return false
		}
		inst_return := find_one_return(callee)
		if inst_return == nil {
			return false
		}
		if len(Alive_mutexs(inst_return)) == 0 {
			return false
		} else {
			return true
		}

	}
	return false
}

// If target_inst is a ssa.Call and it is calling a method or function containing "unlock", we believe it is calling a function containing a Unlock
func Is_self_unlock_brutal(target_inst ssa.Instruction) bool {
	inst_as_call, ok := target_inst.(*ssa.Call)
	if !ok {
		return false
	}
	if inst_as_call.Call.IsInvoke() == true {
		if case_insensitive_contains(inst_as_call.Call.Method.Name(), "unlock") {
			return true
		}
	} else {
		callee, ok := inst_as_call.Call.Value.(*ssa.Function)
		if ok {
			if case_insensitive_contains(callee.Name(), "unlock") {
				return true
			}
		}
	}
	return false
}

// If target_inst is a ssa.Call and it is calling something naming "*.Unlock()", we believe it is calling a function containing a Unlock
func Is_self_unlock(target_inst ssa.Instruction) bool {
	inst_as_call, ok := target_inst.(*ssa.Call)
	if !ok {
		return false
	}
	if inst_as_call.Call.IsInvoke() == true {
		if inst_as_call.Call.Method.Name() == "Unlock" {
			return true
		}
	} else {
		callee, ok := inst_as_call.Call.Value.(*ssa.Function)
		if !ok {
			return false
		}
		if callee.Name() == "Unlock" {
			return true
		}

	}
	return false
}

func find_one_return(fn *ssa.Function) ssa.Instruction {
	for _, bb := range fn.Blocks {
		for _, inst := range bb.Instrs {
			_, ok := inst.(*ssa.Return)
			if ok {
				return inst
			}
		}
	}
	return nil
}
