package channel

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/primitive/locker"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
)

type Channel struct {
	Name string
	Make_inst ssa.Instruction
	Pkg string
	Buffer int

	Sends []*Ch_send
	Recvs []*Ch_recv
	Closes []*Ch_close
	Escapes []*Ch_escape

	Status string
}

type Ch_send struct {
	Name string
	Inst ssa.Instruction // this can be *ssa.Send or *ssa.Select
	Case_index int // If Inst is *ssa.Send, Case_index = -1; else Case_index is the index of case of *ssa.Select
	Is_case_blocking bool
	Whole_line string
	Status string

	Locks []*locker.Lock_op // This field is specially designed for C6A
	Wrappers []*Wrapper     // This field is specially designed for C6A

	Parent *Channel
	chan_op // This field only represent that Ch_send can implement the interface Chan_op
}

type Ch_recv struct {
	Name string
	Inst ssa.Instruction // this can be *ssa.UnOp or *ssa.Select
	Case_index int // If Inst is *ssa.UnOp, Case_index = -1; else Case_index is the index of case of *ssa.Select
	Is_case_blocking bool
	Whole_line string
	Status string

	Locks []*locker.Lock_op // This field is specially designed for C6A
	Wrappers []*Wrapper     // This field is specially designed for C6A

	Parent *Channel
	chan_op
}

type Ch_close struct {
	Name string
	Inst ssa.Instruction // this can be *ssa.Call or *ssa.Defer
	Is_defer bool
	Whole_line string
	Status string

	Locks []*locker.Lock_op // This field is specially designed for C6A
	Wrappers []*Wrapper     // This field is specially designed for C6A

	Parent *Channel
	chan_op
}

type Ch_escape struct {
	Name string
	Inst ssa.Instruction
	Whole_line string
	Status string

	Locks []*locker.Lock_op // This field is specially designed for C6A
	Wrappers []*Wrapper     // This field is specially designed for C6A

	Parent *Channel
	chan_op
}

type chan_op struct {}

func (c *chan_op) interface_mark() {

}

type Chan_op interface {
	interface_mark()
}

// Wrapper records a function that "contains" a chan_op. "Contains" means the function directly uses this chan_op, or
// its callee (or callee's callee) uses this chan_op. The maximum layer is C6A_call_chain_layer_for_chan_wrapper
type Wrapper struct {
	Fn *ssa.Function // When compare two Wrapper, can't directly compare fn or inst, because pointer will change during each compilation
	Fn_str string
	Inst ssa.Instruction
	Callee *Wrapper // if callee is nil, then inst is the chan_op itself, else inst is calling to another Wrapper
	Op Chan_op // the wrapped operation
}

const Edited = "Edited"

const Send,Recv,Close,MakeChan = "Send","Recv","Close","MakeChan"

const Dynamic_size = 12344321

func (ch *Channel) Modify_status(str string) {
	ch.Status = str
}
