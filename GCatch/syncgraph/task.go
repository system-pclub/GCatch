package syncgraph

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/output"
	"github.com/system-pclub/GCatch/GCatch/path"
	"github.com/system-pclub/GCatch/GCatch/tools/go/callgraph"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
)

type Task struct {
	VecTaskPrimitive                []*TaskPrimitive // VecTaskPrimitive means these primitives satify: all its operations are in the graph
	MapValue2TaskPrimitive          map[interface{}]*TaskPrimitive
	MapLCARoot2Op                   map[*ssa.Function][]*ChainsToReachOp
	BoolFinished                    bool
	BoolGiveupIfCallgraphInaccurate bool
}

type TaskPrimitive struct {
	Primitive interface{}
	Ops       map[interface{}]*ChainsToReachOp
	Finished  bool
	Task      *Task
}

// Operation of a primitive, context-sensitive
// E.g. func send() { ch <- 1}
//		func main() {
//			send() // Line 3
//			send() // Line 4
//		}
// There are 2 operations, called by two paths: main() ---Line3---> send(); main() ---Line4---> send()
type ChainsToReachOp struct {
	Op                     interface{}
	Inst                   ssa.Instruction
	Chains                 []*path.EdgeChain
	VecBoolIsChainFinished []bool
	Finished               bool
}

func newTask(boolGiveupIfCallgraphInaccurate bool) *Task {
	return &Task{
		VecTaskPrimitive:                nil,
		MapValue2TaskPrimitive:          make(map[interface{}]*TaskPrimitive),
		MapLCARoot2Op:                   make(map[*ssa.Function][]*ChainsToReachOp),
		BoolFinished:                    false,
		BoolGiveupIfCallgraphInaccurate: boolGiveupIfCallgraphInaccurate,
	}
}

// After this function, a new primitive is added to the task, but TaskPrimitive.Ops are not completed. Need to run Step2CompletePrims after adding all primitives
func (t *Task) Step1AddPrim(newP interface{}, vecChannel []*instinfo.Channel, vecLocker []*instinfo.Locker) {
	newTPrimitive := &TaskPrimitive{
		Primitive: newP,
		Ops:       make(map[interface{}]*ChainsToReachOp),
		Finished:  false,
		Task:      t,
	}
	t.VecTaskPrimitive = append(t.VecTaskPrimitive, newTPrimitive)
	t.MapValue2TaskPrimitive[newP] = newTPrimitive
	thisPrimCh, ok := newP.(*instinfo.Channel)
	if ok && thisPrimCh.MakeInst == nil {
		return
	} else {
		thisPrimLocker, ok2 := newP.(*instinfo.Locker)
		if ok2 && thisPrimLocker.Value == nil {
			return
		}
	}
	// VERI
	// skip dependency check, but append all ops
	for _, otherPrimCh := range vecChannel {
		if otherPrimCh == newP {
			continue
		}
		if otherPrimCh.MakeInst != nil {
			otherTPrimitive := &TaskPrimitive{
				Primitive: otherPrimCh,
				Ops:       make(map[interface{}]*ChainsToReachOp),
				Finished:  false,
				Task:      t,
			}
			t.VecTaskPrimitive = append(t.VecTaskPrimitive, otherTPrimitive)
			t.MapValue2TaskPrimitive[otherPrimCh] = otherTPrimitive
		}
	}

	for _, otherPrimLocker := range vecLocker {
		if otherPrimLocker == newP {
			continue
		}
		if otherPrimLocker.Value != nil {
			otherTPrimitive := &TaskPrimitive{
				Primitive: otherPrimLocker,
				Ops:       make(map[interface{}]*ChainsToReachOp),
				Finished:  false,
				Task:      t,
			}
			t.VecTaskPrimitive = append(t.VecTaskPrimitive, otherTPrimitive)
			t.MapValue2TaskPrimitive[otherPrimLocker] = otherTPrimitive
		}
	}

	//if dPrim, ok := DependMap[newP]; ok {
	//	for _, otherPrim := range dPrim.Circular_depend {
	//		if otherPrimCh, ok := otherPrim.Primitive.(*instinfo.Channel); ok {
	//			if otherPrimCh.MakeInst != nil {
	//				if otherPrimCh.MakeInst.Parent() == thisPrimCh.MakeInst.Parent() {
	//					otherTPrimitive := &TaskPrimitive{
	//						Primitive: otherPrimCh,
	//						Ops:       make(map[interface{}]*ChainsToReachOp),
	//						Finished:  false,
	//						Task:      t,
	//					}
	//					t.VecTaskPrimitive = append(t.VecTaskPrimitive, otherTPrimitive)
	//					t.MapValue2TaskPrimitive[otherPrimCh] = otherTPrimitive
	//				}
	//			}
	//		}
	//	}
	//}
}

var countInaccurateCall, countMaxLayer int

// After adding all primitives that we want, complete everything in each TaskPrimitive.Ops
func (t *Task) Step2CompletePrims() error {

	// Ignore order, list all insts for op in Target_prims
	vecOpInsts := []ssa.Instruction{}
	for _, tPrim := range t.VecTaskPrimitive {
		switch prim := tPrim.Primitive.(type) {
		case *instinfo.Channel:
			for _, op := range prim.AllOps() {
				vecOpInsts = append(vecOpInsts, op.Instr())
			}
		case *instinfo.Locker:
			for _, op := range prim.AllOps() {
				vecOpInsts = append(vecOpInsts, op.Instr())
			}
		}
	}

	LCA2paths, err := path.FindLCA(fnsForInstsNoDupli(vecOpInsts), t.BoolGiveupIfCallgraphInaccurate, true, 20)
	if err != nil {
		//if err == path.ErrInaccurateCallgraph {
		//	fmt.Println("Task: Give up LCA because callgraph is inaccurate. Count:", countInaccurateCall)
		//	countInaccurateCall++
		//}
		if err == path.LcaErrReachedMax {
			if config.Print_Debug_Info {
				fmt.Println("!!!!")
				fmt.Println("Task: Give up LCA because max layer (", config.MAX_LCA_LAYER, ") is reached. Count:", countMaxLayer)
			}

			countMaxLayer++
		}
		return err
	}
	if err == path.LcaErrNilNode {
		return err
	}

	for _, tPrim := range t.VecTaskPrimitive {
		tPrim.CompleteOps(LCA2paths)
	}

	return nil
}

func (t *Task) Update() {
	if t.BoolFinished {
		return
	}

	for _, prim := range t.VecTaskPrimitive {
		if prim.Finished {
			continue
		}
		for _, chainsToReachOp := range prim.Ops {
			if chainsToReachOp.Finished {
				continue
			}
			boolAllChainFinished := true
			for _, boolFinished := range chainsToReachOp.VecBoolIsChainFinished {
				if boolFinished == false {
					boolAllChainFinished = false
					break
				}
			}
			chainsToReachOp.Finished = boolAllChainFinished
		}
	}

	for _, prim := range t.VecTaskPrimitive {
		if prim.Finished {
			continue
		}
		boolAllOpFinished := true
		for _, chainsToReachOp := range prim.Ops {
			if chainsToReachOp.Finished == false {
				boolAllOpFinished = false
				break
			}
		}
		prim.Finished = boolAllOpFinished
	}

	boolAllPrimFinished := true
	for _, prim := range t.VecTaskPrimitive {
		if prim.Finished == false {
			boolAllPrimFinished = false
			break
		}
	}
	t.BoolFinished = boolAllPrimFinished
}

func (t *Task) WantedList() (result []*path.EdgeChain) {
	if t.BoolFinished {
		return nil
	}

	for _, prim := range t.VecTaskPrimitive {
		if prim.Finished {
			continue
		}
		for _, chainsToReachOp := range prim.Ops {
			if chainsToReachOp.Finished {
				continue
			}
			for i, chain := range chainsToReachOp.Chains {
				if chainsToReachOp.VecBoolIsChainFinished[i] {
					continue
				}

				// if this chain is not in result, append it
				boolInResult := false
				for _, edgeChainResult := range result {
					if chain.Equal(edgeChainResult) {
						boolInResult = true
					}
				}
				if boolInResult == false {
					result = append(result, chain)
				}
			}
		}
	}

	return result
}

// A primitive is a target when it is in Task.VecTaskPrimitive
func (t *Task) IsPrimATarget(prim interface{}) bool {
	for _, tPrim := range t.VecTaskPrimitive {
		if tPrim.Primitive == prim {
			return true
		}
	}
	return false
}

func removeVisitedChains(vecWanted, vecVisited []*path.EdgeChain) (result []*path.EdgeChain) {
	for _, wanted := range vecWanted {
		boolVisited := false
		for _, visited := range vecVisited {
			if wanted.Equal(visited) {
				boolVisited = true
				break
			}
		}
		if boolVisited == false {
			result = append(result, wanted)
		}
	}
	return
}

// Generates tp.Ops, and find all callchains from head to an operation
func (tp *TaskPrimitive) CompleteOps(LCA2paths map[*ssa.Function][]*path.EdgeChain) {
	switch p := tp.Primitive.(type) {
	case *instinfo.Channel:
		for _, op := range p.AllOps() {
			new_op := &ChainsToReachOp{
				Op:                     op,
				Inst:                   op.Instr(),
				Chains:                 nil,
				VecBoolIsChainFinished: nil,
				Finished:               false,
			}
			tp.Ops[op] = new_op
		}
	case *instinfo.Locker:
		for _, op := range p.AllOps() {
			new_op := &ChainsToReachOp{
				Op:                     op,
				Inst:                   op.Instr(),
				Chains:                 nil,
				VecBoolIsChainFinished: nil,
				Finished:               false,
			}
			tp.Ops[op] = new_op
		}
	}

	for _, chainsToReachOp := range tp.Ops {
		op_fn := chainsToReachOp.Inst.Parent()

		for lca, paths := range LCA2paths {
			boolFound := false
			for _, onePath := range paths {
				var end *callgraph.Node
				if len(onePath.Chain) == 0 {
					end = onePath.Start
				} else {
					end = onePath.Chain[len(onePath.Chain)-1].Callee
				}
				if end.Func == op_fn {
					boolFound = true
					chainsToReachOp.Chains = append(chainsToReachOp.Chains, onePath)
					chainsToReachOp.VecBoolIsChainFinished = append(chainsToReachOp.VecBoolIsChainFinished, false)
				}
			}
			if boolFound {
				tp.Task.MapLCARoot2Op[lca] = append(tp.Task.MapLCARoot2Op[lca], chainsToReachOp)
			}
		}

		if len(chainsToReachOp.Chains) == 0 {
			fmt.Println("Warning in CompleteOps: can't find any chain for op:")
			output.PrintIISrc(chainsToReachOp.Inst)
		}
	}
}

func removeFromWorklist(old []*Unfinish, remove *Unfinish) (result []*Unfinish) {
	for _, o := range old {
		if o != remove {
			result = append(result, o)
		}
	}
	return
}
