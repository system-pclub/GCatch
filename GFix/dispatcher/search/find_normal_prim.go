//normal primitives include mutex, rwmutex, waitgroup, once, cond, pool
package search

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/sync_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"strings"
)

func find_normal_prim(inst ssa.Instruction) *ssa.Value {


	global_prim_ptr,ok := find_normal_globally(inst)
	if ok {
		return global_prim_ptr
	}

	local_prim_ptr,ok := find_normal_locally(inst)
	if ok {
		return local_prim_ptr
	}

	return nil
}

func recursive_find_normal_globally(inst ssa.Instruction, prim_type string) (*ssa.Value,bool) {
	temp_str := inst.String()
	_ = temp_str
	prim_ptr_type := "*sync." + prim_type
	all_operands := inst.Operands(nil)
	if all_operands == nil {
		return nil,false
	}

	for _,operand := range all_operands {

		if strings.Contains((*operand).String(),"/") { //This operand is trying to visit a global variable or a global variable's field
			for _,sync_global := range global.All_sync_global {
				if strings.Contains((*operand).String(),sync_global.String()) { //This global variable mentioned above is truly a sync global variable
					var sync_global_as_value ssa.Value
					sync_global_as_value = sync_global

					return &sync_global_as_value,true
				}
			}
			fmt.Println("Warning: an inst's operand.String contain / but is not using a sync global\tinst:",inst,"\tinst.operand",(*operand).Name())
		} else if case_insensitive_contains((*operand).Type().String(),prim_ptr_type) { //This operand's type is sync type, meaning it is possible that it's passing a pointer of a global variable, so we investigate recursively

			operand_as_inst,ok := (*operand).(ssa.Instruction)
			if !ok {
				fmt.Println("Warning: an inst's operand is not an inst","\tinst:",inst,"\tinst.operand.Name",(*operand).Name())
				panic(inst)
			}
			return recursive_find_normal_globally(operand_as_inst,prim_type)
		} else {
			return nil,false
		}
	}

	return nil,false
}


func find_normal_globally(inst ssa.Instruction) (*ssa.Value,bool) {
	inst_sync_type := sync_check.Type_sop(inst)
	prim_type := sync_check.Op_to_prim(inst_sync_type)
	prim_ptr_type := "*sync." + prim_type

	all_operands := inst.Operands(nil)
	if all_operands == nil {
		fmt.Println("Warning: a sync inst has no operand","\tinst:",inst)
		panic(inst)
	}
	for _,operand := range all_operands {
		if case_insensitive_contains((*operand).Type().String(),prim_ptr_type) {
			operand_as_inst,ok := (*operand).(ssa.Instruction)
			if ok {
				return recursive_find_normal_globally(operand_as_inst,prim_type)
			} else {
				for _,sync_global := range global.All_sync_global {
					if strings.Contains((*operand).String(),sync_global.String()) { //This global variable mentioned above is truly a sync global variable
						var sync_global_as_value ssa.Value
						sync_global_as_value = sync_global

						return &sync_global_as_value,true
					}
				}
			}
		}
	}
	return nil, false
}

var all_potential_value []*ssa.Value

func find_normal_locally(inst ssa.Instruction) (*ssa.Value,bool) {
	inst_sync_type := sync_check.Type_sop(inst)
	prim_type := "sync." + sync_check.Op_to_prim(inst_sync_type)

	operands := inst.Operands(nil)
	receiver := operands[1] //TODO: I am making an assumption that operands[1] is the method's receiver

	receiver_inst,ok := (*receiver).(ssa.Instruction)
	if !ok {
		fmt.Println("Warning: receiver can't be converted to inst\tinst:",inst.String(),"\treceiver:",(*receiver).Name()," = ",(*receiver).String())
		return nil,false
	}

	all_potential_value = *new([]*ssa.Value)
	all_potential_value = append(all_potential_value,receiver)
	recursive_find_normal_potential_value(receiver_inst.Block(),receiver_inst)

	///DELETE
	fmt.Println("-----------begin: all potential value")
	for _,a := range all_potential_value {
		fmt.Println((*a).Name(), " = ",(*a).String())
	}
	fmt.Println("-----------end:  all potential value")

	all_refined_potential_value := *new([]*ssa.Value)
	for _,potential_value := range all_potential_value {
		potential_inst,ok := (*potential_value).(ssa.Instruction)
		if !ok {
			continue
		}

		if is_inst_alloc_prim(potential_inst,prim_type)  {
			all_refined_potential_value = append(all_refined_potential_value,potential_value)
		}

	}
	///DELETE
	fmt.Println("-----------begin: all refined potential value")
	for _,a := range all_refined_potential_value {
		fmt.Println((*a).Name(), " = ",(*a).String())
	}
	fmt.Println("-----------end: all refined potential value")

	if len(all_refined_potential_value) == 0 {
		return nil,false
	}


	primitive_value := *new([]*ssa.Value)
	outer:
	for _,potential_value := range all_refined_potential_value {
		potential_inst,ok := (*potential_value).(ssa.Instruction)
		if !ok {panic(potential_value)}

		for _,other_value := range all_refined_potential_value {//check if this poteial_value is the bottomest in a BB
			if *other_value == *potential_value {
				continue
			}
			other_inst,ok := (*other_value).(ssa.Instruction)
			if !ok {panic(other_value)}
			if potential_inst.Block() == other_inst.Block() && other_inst.Pos() > potential_inst.Pos() {
				continue outer
			}
		}

		primitive_value = append(primitive_value,potential_value)
	}

	///DELETE
	fmt.Println("\n---------begin: all primitives")
	for _,primitive := range primitive_value {
		prim_position := (global.Prog.Fset).Position((*primitive).Pos())
		fmt.Println("\tSync primitive.String():",(*primitive).String(),"\tprimitive.File:",prim_position.Filename,"\tprimitive.Line:",prim_position.Line)
	}
	fmt.Println("-----------end: all primitives")

	return nil,false

}

func recursive_find_normal_potential_value(bb *ssa.BasicBlock,receiver ssa.Instruction) {
	index_receiver_in_bb := -1
	for index,inst := range bb.Instrs {
		if inst == receiver {
			index_receiver_in_bb = index
		}
	}

	var bottom_index int
	if index_receiver_in_bb == -1 {
		bottom_index = len(bb.Instrs) - 1
	} else {
		bottom_index = index_receiver_in_bb
	}
	if bottom_index < 0 {
		return
	}

	bb_insts := bb.Instrs
	for index := bottom_index; index >= 0; index -- {
		inst := bb_insts[index]
		left,right,left_value_name,right_value_name := split_inst(inst)
		_,_ = left,right

		outer:
		for _,value := range all_potential_value {
			if strings.Contains(left_value_name,(*value).Name()) || strings.Contains(right_value_name,(*value).Name()) {//TODO: add alias

				operands := inst.Operands(nil)
				for _,operand := range operands {
					if strings.Contains(right_value_name,(*operand).Name()) {
						for _,other := range all_potential_value {
							if *other == *operand {
								break outer
							}
						}
						all_potential_value = append(all_potential_value,operand)
						break outer
					}
				}
			}
		}
	}

	for _,previous_bb := range bb.Preds {
		if previous_bb.Parent() != bb.Parent() {
			continue
		}
		recursive_find_normal_potential_value(previous_bb,receiver)
	}

}

func is_inst_alloc_prim(inst ssa.Instruction, prim_type string) bool {

	inst_as_Alloc,ok := inst.(*ssa.Alloc)
	if !ok {
		return false
	} else {
		inst_type := inst_as_Alloc.Type().String()
		prim_type = "*" + prim_type //Alloc.Type() has a "*" before the type you see in inst.String()
		if case_insensitive_equal(inst_type,prim_type) {
			return true
		} else {
			return false
		}
	}

}
//
//func is_inst_alloc_sync_struct(inst ssa.Instruction) bool {
//
//	inst_as_Alloc,ok := inst.(*ssa.Alloc)
//	if !ok {
//		return false
//	} else {
//		inst_type := inst_as_Alloc.Type().String()
//		for _,
//	}
//
//}

func split_inst(inst ssa.Instruction) (string,string,string,string) {
	inst_as_value,ok := inst.(ssa.Value)
	var left,right string
	var left_value_name,right_value_name string
	if ok {
		left = inst_as_value.Name()
		right = inst.String()
		left_value_name = strip_string(inst_as_value.Name())
		right_value_name = strip_string(inst.String())
	} else {
		inst_as_store,ok := inst.(*ssa.Store)
		if ok {
			str_split := strings.Split(inst.String()," = ")
			if len(str_split) != 2 {
				return "","","",""
			}
			left = str_split[0]
			right = str_split[1]
			left_value_name = strip_string(inst_as_store.Addr.Name())
			right_value_name = strip_string(inst_as_store.Val.Name())
		} else {
			return "","","",""
		}
	}
	return left,right,left_value_name,right_value_name
}

func strip_string(str string) string {
	str = strings.ReplaceAll(str,"*","")
	str = strings.ReplaceAll(str,"&","")
	return str
}