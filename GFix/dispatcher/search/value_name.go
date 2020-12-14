package search

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/prepare"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/token"
	"go/types"
	"strings"
)

func Read_position(p token.Position) string {
	filename := p.Filename
	line := p.Line
	if line < 1 || filename == "" {
		return ""
	}

	str_same_line, err := prepare.Read_file_line(filename, line)
	if err != nil {
		return ""
	}

	str_same_line = strings.ReplaceAll(str_same_line,"\n","")

	return str_same_line
}

func Read_inst_line(inst ssa.Instruction) string {
	inst_position := (global.Prog.Fset).Position(inst.Pos())
	filename := inst_position.Filename
	line := inst_position.Line
	if line < 1 || filename == "" {
		return ""
	}

	str_same_line, err := prepare.Read_file_line(filename, line)
	if err != nil {
		return ""
	}

	str_same_line = strings.ReplaceAll(str_same_line,"\n","")

	return str_same_line
}

func Cond_name_and_line(inst ssa.Instruction) (string,string) {
	inst_position := (global.Prog.Fset).Position(inst.Pos())
	inst_str := inst.String()
	_ = inst_str
	filename := inst_position.Filename
	line := inst_position.Line
	if line < 1 {
		//fmt.Println("Error: line < 1")
		return "",""
	}

	str_same_line,err := prepare.Read_file_line(filename,line)
	ori_same_line := str_same_line
	if err != nil {
		//fmt.Println("Error: during read file:",filename,"\tline:",line,"\tfor inst:",inst)
		return "",""
	}

	index_comment := strings.Index(str_same_line,"//")
	if index_comment > -1 {
		str_same_line = str_same_line[:index_comment]
	}

	index_signal := strings.LastIndex(str_same_line,".Signal()")
	index_broadc := strings.LastIndex(str_same_line,".Broadcast()")
	index_wait := strings.LastIndex(str_same_line,".Wait()")
	index_last_dot := -1
	if index_signal >= 1 {
		index_last_dot = index_signal
	} else if index_broadc >= 1 {
		index_last_dot = index_broadc
	} else if index_wait >= 1 {
		index_last_dot = index_wait
	} else {
		return "",""
	}
	cond_name := str_same_line[:index_last_dot]
	if strings.Contains(cond_name,"defer") {
		str_split := strings.Split(cond_name," ")
		cond_name = str_split[len(str_split) - 1]
	}
	cond_name = remove_prefix(cond_name)

	return cond_name,ori_same_line
}

func Mutex_name(inst ssa.Instruction) string {
	inst_position := (global.Prog.Fset).Position(inst.Pos())
	inst_str := inst.String()
	_ = inst_str
	filename := inst_position.Filename
	line := inst_position.Line
	if line < 1 {
		//fmt.Println("Error: line < 1")
		return ""
	}

	str_same_line,err := prepare.Read_file_line(filename,line)
	if err != nil {
		//fmt.Println("Error: during read file:",filename,"\tline:",line,"\tfor inst:",inst)
		return ""
	}

	index_comment := strings.Index(str_same_line,"//")
	if index_comment > -1 {
		str_same_line = str_same_line[:index_comment]
	}

	index_dot := -1
	index_Lock := strings.LastIndex(str_same_line,".Lock")
	index_RLock := strings.LastIndex(str_same_line,".RLock")
	index_Unlock := strings.LastIndex(str_same_line,".Unlock")
	index_RUnlock := strings.LastIndex(str_same_line,".RUnlock")
	switch {
	case index_Lock > 0:
		index_dot = index_Lock
	case index_RLock > 0:
		index_dot = index_RLock
	case index_Unlock > 0:
		index_dot = index_Unlock
	case index_RUnlock > 0:
		index_dot = index_RUnlock
	}
	if index_dot < 1 {
		//fmt.Println("Error: calculating last dot. str_same_line:",str_same_line,"\tfor inst:",inst)
		return ""
	}
	mutex_name := str_same_line[:index_dot]
	str_split := strings.Split(mutex_name," ")
	mutex_name = str_split[len(str_split) - 1]
	mutex_name = strings.TrimSpace(mutex_name)
	mutex_name = remove_prefix(mutex_name)


	return mutex_name
}

// Before use this function, need to make sure that target_inst is ssa.Makechan
func Chan_name_make(target_inst ssa.Instruction) (fullname string,partname string) {

	fullname = ""
	partname = ""
	position := (global.Prog.Fset).Position(target_inst.Pos())
	filename := position.Filename
	if position.Line < 1 {
		return
	}

	str_line,err := prepare.Read_file_line(position.Filename,position.Line)
	if err != nil {
		fmt.Println("Error: during read file:",position.Filename,"\tline:",position.Line,"\tfor inst:",target_inst)
		return
	}

	var str_before_equal string = ""
	index_equal := strings.Index(str_line,"=")
	index_colon := strings.Index(str_line,":")
	index_make := strings.Index(str_line,"make(")
	if index_colon > -1 {
		if index_make < index_colon {
			return
		} else {
			str_between_colon_and_make := str_line[index_colon+1:index_make]
			str_between_colon_and_make = strings.ReplaceAll(str_between_colon_and_make, "=","")
			str_between_colon_and_make = strings.ReplaceAll(str_between_colon_and_make, "\t","")
			str_between_colon_and_make = strings.ReplaceAll(str_between_colon_and_make, " ","")
			str_between_colon_and_make = strings.TrimSpace(str_between_colon_and_make)
			if len(str_between_colon_and_make) > 0 {
				return
			}
		}

		str_before_equal = str_line[:index_colon]

		//The following is to deal with "s := &XXX{ \n ch: make(chan int), \n }
		if index_equal == -1 {
			line_target,line_first := position.Line,-1


			for _,inst := range target_inst.Block().Instrs {
				p := (global.Prog.Fset).Position(inst.Pos())
				if p.Line > 0 {
					line_first = p.Line
					break
				}
			}
			if line_first > -1 && line_first < line_target{
				line_equal := ""
				for i := line_target - 1; i >= line_first; i -- {
					line,err := prepare.Read_file_line(filename,i)
					if err != nil {
						continue
					}
					if strings.Contains(line,"=") {
						line_equal = line
						break
					}
				}
				if line_equal != "" {
					index_equal_ := strings.Index(line_equal,"=")
					line_equal = line_equal[:index_equal_]
					line_equal = strings.ReplaceAll(line_equal,":","")
					line_equal =strings.ReplaceAll(line_equal, "\t","")
					line_equal =strings.ReplaceAll(line_equal, " ","")
					line_equal = strings.TrimSpace(line_equal)
					str_before_equal = line_equal + "." +str_before_equal
				}
			}
		}

	} else if index_equal > -1 {
		if index_make < index_equal {
			return
		} else {
			str_between_colon_and_make := str_line[index_equal+1:index_make]
			str_between_colon_and_make = strings.ReplaceAll(str_between_colon_and_make, "\t","")
			str_between_colon_and_make = strings.ReplaceAll(str_between_colon_and_make, " ","")
			str_between_colon_and_make = strings.TrimSpace(str_between_colon_and_make)
			if len(str_between_colon_and_make) > 0 {
				return
			}
		}
		str_before_equal = str_line[:index_equal]
	} else {
		return
	}


	str_before_equal = strings.ReplaceAll(str_before_equal,"{","")
	str_before_equal = strings.ReplaceAll(str_before_equal,"var","")
	fullname = strings.TrimSpace(str_before_equal)
	fullname = strings.ReplaceAll(fullname,"\t","")
	partname = fullname
	index_dot := strings.LastIndex(partname,".")
	if index_dot > -1 {
		partname = partname[index_dot + 1 :]
	}
	return

}

// Before use this function, need to make sure that target_inst is one of chan_send/receive/close
// If the inst is "<- s.ch1", then this function returns "ch1"
func Chan_name_SRC(target_inst ssa.Instruction) (fullname string,partname string) {


	position := (global.Prog.Fset).Position(target_inst.Pos())
	if position.Line < 1 {
		return "cannot_locate_chan","cannot_locate_chan"
	}

	str_line,err := prepare.Read_file_line(position.Filename,position.Line)
	if err != nil {
		fmt.Println("Error: during read file:",position.Filename,"\tline:",position.Line,"\tfor inst:",target_inst)
		return "cannot_locate_chan", "cannot_locate_chan"
	}

	_,ok := target_inst.(*ssa.UnOp) //ssa.UnOp includes: t0 = *x | t2 = <-t1,ok
	if ok {
		index_arrow := strings.Index(str_line,"<-")
		if index_arrow < 0 {
			return "cannot_locate_chan_receive","cannot_locate_chan_receive"
		}
		str_after_arrow := str_line[index_arrow + 2 :]
		str_after_arrow = strings.TrimSpace(str_after_arrow)
		index_space_or_colon := max(strings.Index(str_after_arrow," "),strings.Index(str_after_arrow,":"))
		if index_space_or_colon >= 0 {
			fullname = str_after_arrow[:index_space_or_colon]
		} else {
			fullname = str_after_arrow
		}
		fullname = remove_prefix(fullname)
		partname = fullname
		index_dot := strings.LastIndex(partname,".")
		if index_dot >= 0 {
			partname = partname[index_dot + 1 :]
		}
		return
	}

	_,ok = target_inst.(*ssa.Send) // chan <- 1
	if ok {
		index_arrow := strings.Index(str_line,"<-")
		if index_arrow < 0 {
			return "cannot_locate_chan_send","cannot_locate_chan_send"
		}
		str_before_arrow := str_line[: index_arrow]
		fullname = str_before_arrow
		fullname = remove_prefix(fullname)
		partname = fullname
		index_dot := strings.LastIndex(partname,".")
		if index_dot >= 0 {
			partname = partname[index_dot + 1 :]
		}
		return
	}

	// Then this must be a channel close operation. Channel is an operand (another operand is function close())
	index_leftparent := strings.Index(str_line,"(")
	str_after_leftparent := str_line[index_leftparent + 1:]
	index_rightparent := strings.Index(str_after_leftparent,")")
	if index_rightparent < 0 {
		return "cannot_locate_chan_close", "cannot_locate_chan_close"
	}
	str_between := str_after_leftparent[: index_rightparent]
	fullname = strings.TrimSpace(str_between)
	fullname = remove_prefix(fullname)
	partname = fullname
	index_dot := strings.LastIndex(partname,".")
	if index_dot >= 0 {
		partname = partname[index_dot + 1 :]
	}
	return

}

// Chan_name_value is typically used on ssa.Value that is ssa.SelectState.Chan or ssa.UnOp.X (Op == Arrow) or ssa.Send.Chan
// This function can return name of channel in ssa.Value, which need to be *ch1 (UnOp+Alloc) or *mytype.fieldch1 (UnOp+FieldAddr+Alloc)
func Chan_name_value(target_value ssa.Value) string {
	inst_UnOp,ok := target_value.(*ssa.UnOp)
	if !ok {
		return ""
	}
	if inst_UnOp.Op != token.MUL {
		return ""
	}
	ch_value := inst_UnOp.X
	// ch_value can be: Alloc, FreeVar
	ch_FreeVar,ok := ch_value.(*ssa.FreeVar)
	if ok {
		return ch_FreeVar.Name()
	}

	ch_Alloc,ok := ch_value.(*ssa.Alloc)
	if ok {
		return ch_Alloc.Comment
	} else {
		ch_FieldAddr,ok := ch_value.(*ssa.FieldAddr)
		if ok {
			// The following is in golang's ssa document, type FieldAddr
			field := ch_FieldAddr.X.Type().Underlying().(*types.Pointer).Elem().Underlying().(*types.Struct).Field(ch_FieldAddr.Field)
			field_name := field.Name()
			receiver_name := "Unknown"

			receiver_alloc,ok := ch_FieldAddr.X.(*ssa.Alloc)
			if ok {
				receiver_name = receiver_alloc.Comment
			}

			receiver_UnOp,ok := ch_FieldAddr.X.(*ssa.UnOp)
			if ok {
				if receiver_UnOp.Op == token.MUL {
					receiver_alloc,ok := receiver_UnOp.X.(*ssa.Alloc)
					if ok {
						receiver_name = receiver_alloc.Comment
					}
				}
			}

			return receiver_name + "." + field_name

		}
	}

	return ""
}

// Before use this function, need to make sure that target_inst is *ssa.Select
// If the inst is "<- s.ch1", then this function returns "ch1"
func Chan_name_select(inst *ssa.Select, state *ssa.SelectState) (fullname string, partname string) {

	position := (global.Prog.Fset).Position(state.Pos)
	if position.Line < 1 {
		return "", ""
	}

	str_line,err := prepare.Read_file_line(position.Filename,position.Line)
	if err != nil {
		fmt.Println("Error: during read file:",position.Filename,"\tline:",position.Line,"\tfor case of inst:",inst)
		return "",""
	}

	if state.Dir == types.SendOnly { // e.g: case chan1 <- "abc":
		index_arrow := strings.Index(str_line,"<-")
		if index_arrow < 0 {
			return "", ""
		}
		str_before_arrow := str_line[: index_arrow]
		fullname =str_before_arrow
		fullname = remove_prefix(fullname)
		partname = fullname
		index_dot := strings.LastIndex(partname,".")
		if index_dot >= 0 {
			partname = partname[index_dot + 1 :]
		}
		return
	} else if state.Dir == types.RecvOnly { // e.g: case x := <- chan1:
		index_arrow := strings.Index(str_line,"<-")
		if index_arrow < 0 {
			return "",""
		}
		str_after_arrow := str_line[index_arrow + 2 :]
		str_after_arrow = strings.TrimSpace(str_after_arrow)
		index_space_or_colon := max(strings.Index(str_after_arrow," "),strings.Index(str_after_arrow,":"))
		if index_space_or_colon >= 0 {
			fullname = str_after_arrow[:index_space_or_colon]
		} else {
			fullname = str_after_arrow
		}
		fullname = remove_prefix(fullname)
		partname = fullname
		index_dot := strings.LastIndex(partname,".")
		if index_dot >= 0 {
			partname = partname[index_dot + 1 :]
		}
		return
	} else {
		return "",""
	}

}

func max(x,y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func remove_prefix(str string) string {
	str = strings.ReplaceAll(str,"case ","")
	str = strings.ReplaceAll(str,"\t","")
	str = strings.TrimSpace(str)
	return str
}
