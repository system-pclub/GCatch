package instinfo

import "github.com/system-pclub/GCatch/tools/go/ssa"

// This file defines primitive channel and its operations

// Define Channel
type Channel struct {
	Name string
	Make_inst *ssa.MakeChan
	Make *ChMake
	Pkg string
	Buffer int

	Sends []*ChSend
	Recvs []*ChRecv
	Closes []*ChClose

	Status string
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
	Whole_line     string // Used for debug
	Status         string

	ChOp // this can be *ssa.Send or *ssa.Select
}

// 		Define operation ChRecv, a concrete implementation of ChanOp
type ChRecv struct {
	Name string
	Case_index int // If Inst is *ssa.UnOp, CaseIndex = -1; else CaseIndex is the index of case of *ssa.Select
	Is_case_blocking bool
	Whole_line string // Used for debug
	Status string

	ChOp // inst can be *ssa.UnOp or *ssa.Select
}

// 		Define operation ChClose, a concrete implementation of ChanOp
type ChClose struct {
	Name string
	Is_defer bool
	Whole_line string
	Status string // Used for debug

	ChOp // inst can be *ssa.Call or *ssa.Defer
}

// A map from inst to its corresponding LockerOp
var mapInst2ChanOp map[ssa.Instruction]map[ChanOp]bool

func ClearChanOpMap() {
	mapInst2ChanOp = make(map[ssa.Instruction]map[ChanOp]bool)
}

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