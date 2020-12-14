package search

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/prepare"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/token"
	"strings"
)

func Receiver_name_of_method(fn ssa.Function) string {
	if fn.Signature.Recv() == nil {
		return ""
	}
	all_params := fn.Params
	if len(all_params) == 0 {
		return ""
	}
	receiver := fn.Params[0]
	return receiver.Name()
}

func First_inst_with_solid_position(fn *ssa.Function) ssa.Instruction {
	bbs := fn.Blocks
	if  len(bbs) == 0 {
		return nil
	}

	for _,inst := range bbs[0].Instrs {
		position := (global.Prog.Fset).Position(inst.Pos())
		if position.Line > 0 && position.Filename != "" {
			return inst
		}
	}

	return nil
}

func Lines_using_name_between(position_a token.Position, position_b token.Position, name string, exclude_line int) (result []string) {
	result = []string{}
	if position_a.Line < 1 || position_b.Line < 1 || position_a.Filename != position_b.Filename || position_a.Line >= position_b.Line {
		return
	}


	for line := position_a.Line; line < position_b.Line; line++ {
		if line == exclude_line {
			continue
		}

		str_line,err := prepare.Read_file_line(position_a.Filename,line)
		if err != nil {
			return
		}

		if index_comment := strings.Index(str_line,"//"); index_comment > -1 {
			str_line = str_line[:index_comment]
		}

		flag_useful_line := false
		for _,index_name := range Str_All_Index(str_line,name) {
			 //index of one occurrence of name in this line
			if index_name > 0 {
				before := str_line[index_name - 1]
				if (before >= 'A' && before <= 'Z') || (before >= 'a' && before <= 'z') {
					continue
				}
			}
			if index_name + len(name) < len(str_line) - 1 {
				after := str_line[index_name + len(name)]
				if (after >= 'A' && after <= 'Z') || (after >= 'a' && after <= 'z') {
					continue
				}
			}
			//the char before and the char after name are both not letter
			flag_useful_line = true
		}

		if flag_useful_line == true {
			result = append(result,str_line)
		}

	}

	return
}

func Str_All_Index(str string, sub string) (result []int) {
	origin_str := str
	sub_len := 0
	result = []int{}
	count := 0
	for {
		count ++
		if count > 1000 {
			break
		}
		index := strings.Index(str,sub)
		if index == -1 {
			break
		}

		result = append(result,index + sub_len)

		str = str[index + len(sub):]
		sub_len = len(origin_str) - len(str)
	}

	return
}