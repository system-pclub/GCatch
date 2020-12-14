package output

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/instruction_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/sync_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/token"
	"go/types"
	"strconv"
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

func File_print_inst_operands(inst ssa.Instruction) { // given an ssa.Instruction, print its operands' names

	line_number := strconv.Itoa((global.Prog.Fset).Position(inst.Pos()).Line)
	inst_value,ok := inst.(ssa.Value)

	if ok {
		inst_type := instruction_check.Type_inst(inst)
		_, err := global.Output_file.WriteString("\nInst:\t" + inst_value.Name() + " = " + inst.String() + "\tType:" + inst_type + "\tLine:" + line_number + "\n")
		if err!= nil {panic(err)}


	} else {
		_, err := global.Output_file.WriteString("\nInst:\t" + inst.String() + "\tType:" + instruction_check.Type_inst(inst) +  "\tLine:" + line_number + "\n")
		if err!= nil {panic(err)}
	}

	all_operands := inst.Operands(nil)
	for _, operand := range all_operands {
		operand_as_value,ok := (*operand).(ssa.Value)
		if !ok {
			return
		}
		_, err := global.Output_file.WriteString("\toperand.Name:" + operand_as_value.Name() + "\toperand.Type:" + operand_as_value.Type().String() + "\toperand.String:" + operand_as_value.String() + "\n")
		if err!= nil {panic(err)}

	}


}

func Print_insts_in_fn(fn *ssa.Function) {
	fmt.Println("--------------Func:",fn.String())
	for _,bb := range fn.Blocks {
		fmt.Println("---bb:",bb.Index," ",bb.Comment)
		for _,inst := range bb.Instrs {
			Print_inst(inst)
		}
	}
}

func Print_all_struct(){
	fmt.Println("-------------Begin: in global.All_struct")
	for _,stru := range global.All_struct {
		fmt.Println("\nName:",stru.Name)
		for field_name,field_type := range stru.Field {
			fmt.Println("\tField_Name:",field_name,"\tField_Type:",field_type)
		}
	}
	fmt.Println("-------------End: in global.All_struct")
}



func Print_inst_operands(inst ssa.Instruction) { // given an ssa.Instruction, print its operands' names

	inst_position := (global.Prog.Fset).Position(inst.Pos())
	line_number := strconv.Itoa(inst_position.Line)
	file_name := inst_position.Filename
	inst_value,ok := inst.(ssa.Value)

	if ok {
		inst_type := instruction_check.Type_inst(inst)
		fmt.Println("Inst:\t" + inst_value.Name() + " = " + inst.String() + "\tType:" + inst_type + "\tFile:" + file_name + "\tLine:" + line_number + "\n")

	} else {
		fmt.Println("Inst:\t" + inst.String() + "\tType:" + instruction_check.Type_inst(inst) + "\tFile:" + file_name +  "\tLine:" + line_number + "\n")
	}

	all_operands := inst.Operands(nil)
	if all_operands == nil {
		return
	}
	for _, operand := range all_operands {
		operand_as_value,ok := (*operand).(ssa.Value)
		if !ok {
			return
		}
		fmt.Println("\toperand.Name:" + operand_as_value.Name() + "\toperand.Type:" + operand_as_value.Type().String() + "\toperand.String:" + operand_as_value.String() + "\n")
	}


}

func Print_inst(inst ssa.Instruction) { // given an ssa.Instruction, print its operands' names

	inst_value,ok := inst.(ssa.Value)
	if ok {
		fmt.Println("Inst:\t" + inst_value.Name() + " = " + inst.String() + "\t" + inst_value.Type().String())
	} else {
		fmt.Println("Inst:\t" + inst.String())
	}
}

func Print_inst_with_type_and_location(inst ssa.Instruction) { // given an ssa.Instruction, print its operands' names

	inst_position := (global.Prog.Fset).Position(inst.Pos())
	line_number := strconv.Itoa(inst_position.Line)
	file_name := inst_position.Filename
	inst_value,ok := inst.(ssa.Value)

	if ok {
		inst_type := instruction_check.Type_inst(inst)
		fmt.Println("Inst:\t" + inst_value.Name() + " = " + inst.String() + "\tType:" + inst_type + "\tFile:" + file_name + "\tLine:" + line_number + "\n")

	} else {
		fmt.Println("Inst:\t" + inst.String() + "\tType:" + instruction_check.Type_inst(inst) + "\tFile:" + file_name +  "\tLine:" + line_number + "\n")
	}

}

func Print_value_and_location(value ssa.Value) { // given an ssa.Instruction, print its operands' names

	var print_position token.Position
	value_position := (global.Prog.Fset).Position(value.Pos())
	if value_position.IsValid() {
		print_position = value_position
	} else {
		//if the line == 0, we will first try to convert value into inst
		if inst_value,ok := value.(ssa.Instruction); ok {
			// if succeed, find the latter inst in the same bb which has line number
			insts := inst_value.Block().Instrs
			target_index := -1
			for index,inst := range insts {
				if inst == inst_value {
					target_index = index
					break
				}
			}
			for i:=target_index; i<len(insts); i++ {
				inst := insts[i]
				inst_position := (global.Prog.Fset).Position(inst.Pos())
				if inst_position.Line > 0 {
					print_position = inst_position
					break
				}
			}
		}
	}

	if print_position.IsValid() == false {
		value_inst,ok := value.(ssa.Instruction)

		if ok {
			fmt.Print("Value:\t" + value.Name() + " = " + value_inst.String() + "\tType:" + value.Type().String() + "\t Function:" + value.Parent().String()  + "\t BB:" + value_inst.Block().String() + "\n")
		} else {
			fmt.Print("Value:\t" + value.String() + "\tType:" + value.Type().String() + "\t Function:" + value.Parent().String() + "\n")
		}
	} else {
		line_number := strconv.Itoa(print_position.Line)
		file_name := print_position.Filename
		value_inst,ok := value.(ssa.Instruction)

		if ok {
			fmt.Print("Value:\t" + value.Name() + " = " + value_inst.String() + "\tType:" + value.Type().String() + "\t File:" + file_name  + ":" + line_number + "\n")

		} else {
			fmt.Print("Value:\t" + value.String() + "\tType:" + value.Type().String() + "\t File:" + file_name  +  ":" + line_number + "\n")
		}
	}
	return
}

func Print_inst_debugref(inst ssa.Instruction) {

	inst_as_dr, ok := inst.(*ssa.DebugRef)
	if ok {
		fmt.Println("\tThis is a DebugRef inst:", inst_as_dr, "\tExpr:", inst_as_dr.Expr, "\tIsAddr:", inst_as_dr.IsAddr, "\tX:", inst_as_dr.X, "\tX.type:", inst_as_dr.X.Type())
	}
}

func Print_inst_phi(inst ssa.Instruction) {

	inst_as_phi, ok := inst.(*ssa.Phi)
	if ok {
		fmt.Println("\tThis is a Phi inst:", inst_as_phi, "\tComment:", inst_as_phi.Comment, "\tEdge[0]:", inst_as_phi.Edges[0])
	}
}

func Print_inst_call(inst ssa.Instruction) { // given an ssa.Instruction, this function can print its callName if it is a ssa.Call
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	fmt.Println("Inst:", inst, "\tInst.Call:", call_)

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := sync_check.CallName(call)
		fmt.Println("callName:", callName)
	}

}

func Print_inst_alloc(inst ssa.Instruction) { // given an ssa.Instruction, this function can print results of all functions under ssa.Call

	inst_asAlloc, ok := inst.(*ssa.Alloc)
	if ok {
		fmt.Println("\tThis is an alloc Inst:", inst_asAlloc, "\tvariable type:",inst_asAlloc.Type().Underlying().(*types.Pointer).Elem())
	}
}


