package instinfo

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/output"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"strconv"
)

// This file defines primitive channel and its operations

// Define Channel
type Channel struct {
	Name string
	MakeInst *ssa.MakeChan
	Make *ChMake
	Pkg string
	Buffer int

	Sends []*ChSend
	Recvs []*ChRecv
	Closes []*ChClose

	Status string
}

func (ch *Channel) AllOps() []ChanOp {
	result := []ChanOp{}
	result = append(result, ch.Make)
	for _, send := range ch.Sends {
		result = append(result, send)
	}
	for _, recv := range ch.Recvs {
		result = append(result, recv)
	}
	for _, c := range ch.Closes {
		result = append(result, c)
	}
	return result
}

// Define interface ChanOp, and its implementation ChOp, which is inherited by all concrete implementations
type ChanOp interface {
	Prim() *Channel
	Instr() ssa.Instruction
}

type ChOp struct {
	Parent *Channel
	Inst ssa.Instruction
}

func (op *ChOp) Prim() *Channel {
	return op.Parent
}

func (op *ChOp) Instr() ssa.Instruction {
	return op.Inst
}

// 		Define operation ChMake, a concrete implementation of ChanOp
type ChMake struct {
	ChOp // inst can only be MakeChan
}

// 		Define operation ChSend, a concrete implementation of ChanOp
type ChSend struct {
	Name           string
	CaseIndex      int // If Inst is *ssa.Send, CaseIndex = -1; else CaseIndex is the index of case of *ssa.Select
	IsCaseBlocking bool
	WholeLine     string // Used for debug
	Status         string

	ChOp // this can be *ssa.Send or *ssa.Select
}

// 		Define operation ChRecv, a concrete implementation of ChanOp
type ChRecv struct {
	Name string
	CaseIndex int // If Inst is *ssa.UnOp, CaseIndex = -1; else CaseIndex is the index of case of *ssa.Select
	IsCaseBlocking bool
	WholeLine string // Used for debug
	Status string

	ChOp // inst can be *ssa.UnOp or *ssa.Select
}

// 		Define operation ChClose, a concrete implementation of ChanOp
type ChClose struct {
	Name string
	IsDefer bool
	WholeLine string
	Status string // Used for debug

	ChOp // inst can be *ssa.Call or *ssa.Defer
}

// A map from inst to its corresponding LockerOp
var MapInst2ChanOp map[ssa.Instruction][]ChanOp

func ClearChanOpMap() {
	MapInst2ChanOp = make(map[ssa.Instruction][]ChanOp)
}

const DynamicSize = -999 // If a channel's buffer size can't be computed statically, we give DynamicSize to the Buffer field

// Define some special channels and its operations
var ChanTimer Channel
var ChanContext Channel
var ChanNotDepend Channel

// Add a special send, belongs to a channel that we don't consider in this run, because it is not a dependent primitive
// No sync constraint will be generated for this operation.
// Fields like Case_index need to be updated
func AddNotDependSend(inst ssa.Instruction) *ChSend {
	newSend := &ChSend{
		ChOp:            ChOp{
			Parent: &ChanNotDepend,
			Inst:   inst,
		},
	}
	ChanNotDepend.Sends = append(ChanNotDepend.Sends, newSend)
	return newSend
}

// Add a special receive, belongs to a channel that we don't consider in this run, because it is not a dependent primitive
// No sync constraint will be generated for this operation.
// Fields like Case_index need to be updated
func AddNotDependRecv(inst ssa.Instruction) *ChRecv {
	newRecv := &ChRecv{
		ChOp:            ChOp{
			Parent: &ChanNotDepend,
			Inst:   inst,
		},
	}
	ChanNotDepend.Recvs = append(ChanNotDepend.Recvs, newRecv)
	return newRecv
}

// Add a special close, belongs to a channel that we don't consider in this run, because it is not a dependent primitive
// No sync constraint will be generated for this operation.
func AddNotDependClose(inst ssa.Instruction) *ChClose {
	newClose := &ChClose{
		ChOp:            ChOp{
			Parent: &ChanNotDepend,
			Inst:   inst,
		},
	}
	ChanNotDepend.Closes = append(ChanNotDepend.Closes, newClose)
	return newClose
}

func (ch *Channel) DebugPrintChan() {

	fmt.Println("------Chan:", ch.Name,"\tIn", ch.Pkg)
	fmt.Println("---Buffer:", ch.Buffer)

	if ch.MakeInst != nil {
		m_p := (config.Prog.Fset).Position(ch.MakeInst.Pos())
		fmt.Println("---Make:", ch.MakeInst,"\tat:",m_p.Filename+":"+strconv.Itoa(m_p.Line))
	}

	fmt.Println("---Send:", len(ch.Sends))
	for i, send := range ch.Sends {
		p := (config.Prog.Fset).Position(send.Inst.Pos())
		fmt.Print("--",i,":",p.Filename+":"+strconv.Itoa(p.Line))
		output.PrintIISrc(send.Inst)
		fmt.Println()
		fmt.Println(" In case:",send.CaseIndex," Select_blocking:", send.IsCaseBlocking)

	}
	fmt.Println("---Recv:", len(ch.Recvs))
	for i, recv := range ch.Recvs {
		p := (config.Prog.Fset).Position(recv.Inst.Pos())
		fmt.Println("--",i,":",p.Filename+":"+strconv.Itoa(p.Line))
		output.PrintIISrc(recv.Inst)
		fmt.Println()
		fmt.Println(" In case:",recv.CaseIndex," Select_blocking:", recv.IsCaseBlocking)
	}
	fmt.Println("---Close:", len(ch.Closes))
	for i, aClose := range ch.Closes {
		p := (config.Prog.Fset).Position(aClose.Inst.Pos())
		fmt.Println("--",i,":",p.Filename+":"+strconv.Itoa(p.Line)," In defer:", aClose.IsDefer)
		output.PrintIISrc(aClose.Inst)
		fmt.Println()
	}
	fmt.Print()
}