package channel

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"strconv"
)

func (a_chan *Channel) Debug_print_chan() {

	fmt.Println("------Chan:",a_chan.Name,"\tIn",a_chan.Pkg)
	fmt.Println("---Buffer:",a_chan.Buffer)

	if a_chan.Make_inst != nil {
		m_p := (global.Prog.Fset).Position(a_chan.Make_inst.Pos())
		fmt.Println("---Make:",a_chan.Make_inst,"\tat:",m_p.Filename+":"+strconv.Itoa(m_p.Line))
	}

	fmt.Println("---Send:", len(a_chan.Sends))
	for i,send := range a_chan.Sends {
		p := (global.Prog.Fset).Position(send.Inst.Pos())
		fmt.Println("--",i,":",p.Filename+":"+strconv.Itoa(p.Line),send.Whole_line)
		fmt.Println(" In case:",send.Case_index," Select_blocking:",send.Is_case_blocking)

	}
	fmt.Println("---Recv:", len(a_chan.Recvs))
	for i,recv := range a_chan.Recvs {
		p := (global.Prog.Fset).Position(recv.Inst.Pos())
		fmt.Println("--",i,":",p.Filename+":"+strconv.Itoa(p.Line),recv.Whole_line)
		fmt.Println(" In case:",recv.Case_index," Select_blocking:",recv.Is_case_blocking)
	}
	fmt.Println("---Close:", len(a_chan.Closes))
	for i,close := range a_chan.Closes {
		p := (global.Prog.Fset).Position(close.Inst.Pos())
		fmt.Println("--",i,":",p.Filename+":"+strconv.Itoa(p.Line),close.Whole_line," In defer:",close.Is_defer)
	}
	fmt.Print()
}
