package output

import (
	"fmt"
	"github.com/system-pclub/gochecker/config"
	"github.com/system-pclub/gochecker/tools/go/ssa"
	"github.com/system-pclub/gochecker/util"
	"go/token"
)

func GetLineNum(II ssa.Instruction) int {
	loc := (config.Prog.Fset).Position(II.Pos())

	if loc.Line > 0 {
		return loc.Line
	}

	iiIndex := util.GetIIndexBB(II) - 1
	bbIndex := II.Block().Index

	for bbIndex >= 0 {
		for iiIndex >= 0 {
			I := II.Parent().Blocks[bbIndex].Instrs[iiIndex]
			loc = (config.Prog.Fset).Position(I.Pos())
			if loc.Line > 0 {
				return loc.Line
			}

			iiIndex --
		}

		bbIndex --
		iiIndex = len(II.Parent().Blocks[bbIndex].Instrs) -1
	}

	/*
	for index >= 0 {
		I := II.Block().Instrs[index]
		loc = (config.Prog.Fset).Position(I.Pos())
		if loc.Line > 0 {
			return loc.Line
		}
		index = index -1
	}
	 */

	return 0
}


func GetLoc(II ssa.Instruction) token.Position {
	loc := (config.Prog.Fset).Position(II.Pos())

	if loc.Line > 0 {
		return loc
	}

	iiIndex := util.GetIIndexBB(II) - 1
	bbIndex := II.Block().Index

	for bbIndex >= 0 {
		for iiIndex >= 0 {
			I := II.Parent().Blocks[bbIndex].Instrs[iiIndex]
			loc = (config.Prog.Fset).Position(I.Pos())
			if loc.Line > 0 {
				return loc
			}

			iiIndex --
		}

		bbIndex --
		iiIndex = len(II.Parent().Blocks[bbIndex].Instrs) -1
	}

	return token.Position {Line: 0}
}

func PrintFnSrc(fn * ssa.Function) {
	for _, bb := range fn.Blocks {
		for _, ii := range bb.Instrs {
			loc := (config.Prog.Fset).Position(ii.Pos())
			if loc.Line > 0 {
				fmt.Print("\tFile:", loc.Filename,"\tLine:", loc.Line)
				fmt.Println()
				return
			}
		}
	}
}

func PrintIISrc(ii ssa.Instruction) {
	loc := GetLoc(ii)

	if loc.Line != 0 {
		fmt.Print("\tFile:", loc.Filename,"\tLine:", loc.Line)
		fmt.Println()
	}

}


func PrintInsts( IIs [] ssa.Instruction) {

	if len(IIs) == 0 {
		return
	}

	firstII := IIs[0]
	loc := (config.Prog.Fset).Position(firstII.Pos())

	if loc.Line > 0 {
		fmt.Print("\tFile:", loc.Filename,"\tLine:")
	} else {
		flag := false

		/*
		index := util.GetIIndexBB(firstII) - 1

		for index >= 0 {
			I := firstII.Block().Instrs[index]
			loc = (config.Prog.Fset).Position(I.Pos())
			if loc.Line > 0 {
				flag = true
				fmt.Print("\tFile:", loc.Filename, "\tLine:")
				break
			}
			index = index -1
		}
		*/

		iiIndex := util.GetIIndexBB(firstII) - 1
		bbIndex := firstII.Block().Index
outer:
		for bbIndex >= 0 {
			for iiIndex >= 0 {
				I := firstII.Parent().Blocks[bbIndex].Instrs[iiIndex]
				loc = (config.Prog.Fset).Position(I.Pos())
				if loc.Line > 0 {
					flag = true
					fmt.Print("\tFile:", loc.Filename, "\tLine:")
					break outer
				}

				iiIndex --
			}

			bbIndex --
			iiIndex = len(firstII.Parent().Blocks[bbIndex].Instrs) -1
		}


		if !flag {
			//None of the insts in the same bb has line number, we will just report the name of the function
			fmt.Println("\tInside function:", firstII.Parent().String(),"\tLine:")
		}
	}

	for _, II := range IIs[:len(IIs)-1] {
		fmt.Print(GetLineNum(II), ", ")
	}

	fmt.Println(GetLineNum(IIs[len(IIs)-1]))
	fmt.Print("\n")
}
