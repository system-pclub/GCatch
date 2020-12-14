package output

import (
	"bufio"
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"os"
)

func Wait_for_input() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("\nPlease press Enter to continue")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(sentence))
	}
}

func Print_inst_only_location(target_inst ssa.Instruction) {
	target_position := (global.Prog.Fset).Position(target_inst.Pos())
	if target_position.Line > 0 {
		fmt.Printf("\tFile: %s:%d\n",target_position.Filename,target_position.Line)
		return
	}

	//if the line == 0, we will find the former inst in the same bb which has line number
	insts := target_inst.Block().Instrs
	for index,_ := range insts {
		inst := insts[len(insts) - index - 1]
		inst_position := (global.Prog.Fset).Position(inst.Pos())
		if inst_position.Line > 0 {
			fmt.Printf("\tFile: %s:%d\n",inst_position.Filename,inst_position.Line)
			return
		}
	}

	//None of the insts in the same bb has line number, we will just report the name of the function
	fmt.Println("\tInside function:",target_inst.Parent().String())
}

func Print_inst_and_location(target_inst ssa.Instruction) {
	Print_inst(target_inst)
	target_position := (global.Prog.Fset).Position(target_inst.Pos())
	if target_position.Line > 0 {
		fmt.Printf("\tFile: %s:%d\n",target_position.Filename,target_position.Line)
		return
	}

	//if the line == 0, we will find the former inst in the same bb which has line number
	insts := target_inst.Block().Instrs
	target_index := -1
	for index,inst := range insts {
		if inst == target_inst {
			target_index = index
			break
		}
	}
	for i:=target_index; i>=0; i-- {
		inst := insts[i]
		inst_position := (global.Prog.Fset).Position(inst.Pos())
		if inst_position.Line > 0 {
			fmt.Printf("\tFile: %s:%d\n",inst_position.Filename,inst_position.Line)
			return
		}
	}

	//None of the insts in the same bb has line number, we will just report the name of the function
	fmt.Println("\tInside function:",target_inst.Parent().String())

}

func Print_insts_and_one_location(target_insts []ssa.Instruction) {

	if len(target_insts) == 0 {
		return
	}

	first_inst := target_insts[0]
	target_position := (global.Prog.Fset).Position(first_inst.Pos())
	if target_position.Line > 0 {
		fmt.Print("\tFile:",target_position.Filename,"\tLine:")
	} else {
		//if the line == 0, we will find the former inst which has line number
		insts := first_inst.Block().Instrs
		flag_find := false
		for index,_ := range insts {
			inst := insts[len(insts) - index - 1]
			inst_position := (global.Prog.Fset).Position(inst.Pos())
			if inst_position.Line > 0 {
				fmt.Print("\tFile:",inst_position.Filename, "\tLine:")
				flag_find = true
				break
			}
		}

		if flag_find == false {
			//None of the insts in the same bb has line number, we will just report the name of the function
			fmt.Println("\tInside function:",first_inst.Parent().String(),"\tLine:")
		}
	}

	for _,target_inst := range target_insts {
		target_position := (global.Prog.Fset).Position(target_inst.Pos())
		if target_position.Line > 0 {
			fmt.Print(target_position.Line,", ")
		} else {
			//if the line == 0, we will find the former inst which has line number
			insts := first_inst.Block().Instrs
			for index,_ := range insts {
				inst := insts[len(insts) - index - 1]
				inst_position := (global.Prog.Fset).Position(inst.Pos())
				if inst_position.Line > 0 {
					fmt.Print(inst_position.Line,", ")
					break
				}
			}
		}
	}

	fmt.Print("\n")



}
