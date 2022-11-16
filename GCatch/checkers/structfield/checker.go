package structfield

import (
	"fmt"
	"strconv"

	"github.com/system-pclub/GCatch/GCatch/config"
	"golang.org/x/tools/go/ssa"
)

var C3_all_structs []*C3_struct
var C3_reported_fields []*Field_usage

type C3_struct struct {
	Name  string
	Field map[string](map[ssa.Instruction][]string) //map[field_name](map[inst_used_field][]str_alive_mutexs)
}

type Field_usage struct {
	Type          string
	Field_name    string
	Usage         map[ssa.Instruction][]string
	Debug_Usage   map[string][]string
	Nu_locked     int
	Nu_not_locked int
}

type Debug_usage struct {
	Type          string
	Field_name    string
	Usage         map[string][]string
	Nu_locked     int
	Nu_not_locked int
}

func init() {
	C3_reported_fields = *new([]*Field_usage)
}

func C3_cleanup() {
	C3_all_structs = *new([]*C3_struct)
	searched_bb = *new([]*ssa.BasicBlock)

}

func Detect() {

	C3_cleanup()

	vecAllStruct := List_all_struct(config.Prog)

	for _, normal_struct := range vecAllStruct {
		field_map := make(map[string](map[ssa.Instruction][]string))
		for normal_struct_field_name, _ := range normal_struct.Field {
			field_map[normal_struct_field_name] = make(map[ssa.Instruction][]string)
		}
		C3_a_struct := C3_struct{
			Name:  normal_struct.Name,
			Field: field_map,
		}
		C3_all_structs = append(C3_all_structs, &C3_a_struct)
	}

	loop_pkg_C3() //this step is to prepare C3_all_structs; after this step, for every usage of every field in C3_all_structs;
	//we have a []string, containing all alive mutex/rwmutex
	//important part of this function: see loop_BB

	var susp_fields []*Field_usage
	susp_fields = pickup_susp_fields(susp_fields) //from all structs, append fields whose map contains at least one mutex

	for _, susp_field := range susp_fields {
		for _, mutexs := range susp_field.Usage {
			if len(mutexs) > 0 {
				susp_field.Nu_locked++
			} else {
				susp_field.Nu_not_locked++
			}
		}
	}

	var buggy_fields []*Field_usage
	for _, susp_field := range susp_fields {
		protected := float32(susp_field.Nu_locked)
		unprotected := float32(susp_field.Nu_not_locked)
		ratio := protected / (protected + unprotected)
		if ratio >= config.STRUCT_RATIO && ratio < 1 {
			buggy_fields = append(buggy_fields, susp_field)
		}

	}

	for _, buggy_field := range buggy_fields {
		for inst, mutexs := range buggy_field.Usage {
			if len(mutexs) == 0 {
				if is_FP(inst, buggy_field.Type, buggy_field.Field_name) == true {
					buggy_field.Usage[inst] = []string{"FP"}
					buggy_field.Nu_not_locked--
					buggy_field.Nu_locked++

				}
			}
		}
	}

	var refined_buggy_fields []*Field_usage
	for _, buggy_field := range buggy_fields {
		has_unprotected_usage := false
		for _, mutexs := range buggy_field.Usage {
			if len(mutexs) == 0 {
				has_unprotected_usage = true
				break
			}
		}
		if has_unprotected_usage == false {
			continue
		}
		protected := float32(buggy_field.Nu_locked)
		unprotected := float32(buggy_field.Nu_not_locked)
		ratio := protected / (protected + unprotected)
		if ratio >= config.STRUCT_RATIO && ratio < 1 {
			refined_buggy_fields = append(refined_buggy_fields, buggy_field)
		}
	}
	for _, refined_buggy_field := range refined_buggy_fields {
		for inst, mutexs := range refined_buggy_field.Usage {
			inst_position := (config.Prog.Fset).Position(inst.Pos())
			str := inst_position.Filename + ":" + strconv.Itoa(inst_position.Line)
			refined_buggy_field.Debug_Usage[str] = mutexs
		}
	}

outer:
	for _, refined_buggy_field := range refined_buggy_fields {
		for _, reported_field := range C3_reported_fields {
			if reported_field.Type == refined_buggy_field.Type && reported_field.Field_name == refined_buggy_field.Field_name {
				continue outer
			}
		}
		if len(refined_buggy_field.Debug_Usage) < config.STRUCT_MIN_TIME_OF_USAGE {
			continue outer
		}

		C3_reported_fields = append(C3_reported_fields, refined_buggy_field)

		config.BugIndexMu.Lock()
		config.BugIndex++
		fmt.Print("----------Bug[")
		fmt.Print(config.BugIndex)
		config.BugIndexMu.Unlock()
		fmt.Print("]----------\n\tType: Inconsistent Field Protection \tReason: a field in a structure is sometimes protected by Mutex, but sometimes unprotected.\n")
		fmt.Print("\tStructure:", refined_buggy_field.Type, "\tField:", refined_buggy_field.Field_name, "\n")
		fmt.Print("\tWhere it is unprotected:\n")
		for inst_str, mutexs := range refined_buggy_field.Debug_Usage {
			if len(mutexs) == 0 {
				fmt.Print("\t\tInst at ", inst_str, "\n")
			}
		}
		fmt.Print("\tWhere it is protected:\n")
		for inst_str, mutexs := range refined_buggy_field.Debug_Usage {
			if len(mutexs) > 0 {
				fmt.Print("\t\tInst at ", inst_str, "\tProtected by:", mutexs, "\n")
			}
		}

	}

}

// A strategy to reduce FP: if the unprotected instruction meets one of the following contitions, then it is a FP:
// (1) Its parent function contains a key word in config.C3ExcludeSlice
// (2) When its parent is called, there is a mutex alive
// (3) call_inst is the instruction in another function that calls its parent. Recursively inspect (1) and (2) of all call_inst.
//
//	The layer of inspecting and the ratio is determined by config.STRUCT_FP_LAYER and config.STRUCT_FP_RATIO
func is_FP(target_inst ssa.Instruction, type_name string, field_name string) bool {
	target_position := (config.Prog.Fset).Position(target_inst.Pos())
	_ = target_position
	var exclude_str []string
	exclude_str = []string{"init", "close", "start", "new", "lockfree", "shutdown"}
	count := 0

	return recursive_is_FP(target_inst, exclude_str, count)

}

func recursive_is_FP(target_inst ssa.Instruction, exclude_str []string, count int) bool {
	target_position := (config.Prog.Fset).Position(target_inst.Pos())
	_ = target_position
	if count >= config.STRUCT_FP_LAYER {
		return true
	}

	if is_parent_exclude(target_inst, exclude_str) == true {
		return true
	}

	if len(Alive_mutexs(target_inst)) > 0 {
		return true
	}

	call_insts := list_call_insts(target_inst)
	num_call_insts := len(call_insts)
	if num_call_insts == 0 {
		return false
	}
	num_FP_call_inst := 0
	for _, call_inst := range call_insts {
		call_position := (config.Prog.Fset).Position(call_inst.Pos())
		_ = call_position
		if is_parent_exclude(call_inst, exclude_str) == true {
			num_FP_call_inst++
			continue
		}
		if len(Alive_mutexs(call_inst)) > 0 {
			num_FP_call_inst++
			continue
		}

		if call_inst.Parent().Name() == target_inst.Parent().Name() { // This will cause recursively visit the same function
			continue
		}

		count++
		if recursive_is_FP(call_inst, exclude_str, count) == true {
			num_FP_call_inst++
			continue
		}

	}
	if float32(num_FP_call_inst)/float32(num_call_insts) > config.STRUCT_FP_RATIO {
		return true
	} else {
		return false
	}
}

func list_call_insts(target_inst ssa.Instruction) []ssa.Instruction {
	var call_insts []ssa.Instruction

	call_insts = FP_loop_pkg(call_insts, target_inst)
	return call_insts
}

func is_parent_exclude(target_inst ssa.Instruction, exclude_str []string) bool {
	for _, str := range exclude_str {
		if case_insensitive_contains(target_inst.Parent().Name(), str) {
			return true
		}
	}
	return false

}

func pickup_susp_fields(susp_fields []*Field_usage) []*Field_usage {
	for _, C3_struct := range C3_all_structs {
		for field_name, field_map := range C3_struct.Field {
			flag_susp := false
			for _, mutexs := range field_map {
				if len(mutexs) > 0 {
					flag_susp = true
				}
			}

			if flag_susp {
				susp_field := Field_usage{
					Type:          C3_struct.Name,
					Field_name:    field_name,
					Usage:         field_map,
					Debug_Usage:   make(map[string][]string),
					Nu_locked:     0,
					Nu_not_locked: 0,
				}

				susp_fields = append(susp_fields, &susp_field)
			}
		}
	}
	return susp_fields
}
