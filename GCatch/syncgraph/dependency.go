package syncgraph

import (
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/path"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

type DPrim struct {
	Primitive        interface{} // *instinfo.Channel or *instinfo.Locker
	Out              []*DEdge
	In               []*DEdge
	Circular_depend  []*DPrim
	place            string
	depend_on_places []string
}

type DEdge struct {
	Caller *DPrim
	Callee *DPrim
}

func (prim *DPrim) AddOutEdge(prim2 *DPrim) {
	for _, outEdge := range prim.Out {
		if outEdge.Callee == prim2 {
			return
		}
	}
	newEdge := &DEdge{
		Caller: prim,
		Callee: prim2,
	}
	prim.Out = append(prim.Out, newEdge)
	prim2.In = append(prim2.In, newEdge)
}

func (prim *DPrim) AddCircularDepend(prim2 *DPrim) {
	for _, existDepend := range prim.Circular_depend {
		if existDepend == prim2 {
			return
		}
	}
	prim.Circular_depend = append(prim.Circular_depend, prim2)
	for _, existDepend := range prim2.Circular_depend {
		if existDepend == prim {
			return
		}
	}
	prim2.Circular_depend = append(prim2.Circular_depend, prim)
}

func (prim *DPrim) IsCircularDepend(prim2 *DPrim) bool {
	hasIn, hasOut := false, false
	for _, outEdge := range prim.Out {
		if outEdge.Callee == prim2 {
			hasOut = true
			break
		}
	}
	for _, inEdge := range prim.In {
		if inEdge.Caller == prim2 {
			hasIn = true
			break
		}
	}
	return hasIn && hasOut
}

func GenDMap(vecChan []*instinfo.Channel, vecLocker []*instinfo.Locker) (DMap map[interface{}]*DPrim) {
	DMap = make(map[interface{}]*DPrim)

	// Store primitives that have blocking operation in a function
	Fn2Prims := make(map[*ssa.Function]map[*DPrim]struct{})

	for fn, _ := range ssautil.AllFunctions(config.Prog) {
		mapPrim := Fn2Prims[fn]
		if mapPrim == nil {
			mapPrim = make(map[*DPrim]struct{})
			Fn2Prims[fn] = mapPrim
		}
		for _, bb := range fn.Blocks {
			for _, inst := range bb.Instrs {
				vecChOp, ok := instinfo.MapInst2ChanOp[inst]
				if ok {
					for _, chOp := range vecChOp {
						isBlocking := false
						switch chOp.(type) {
						case *instinfo.ChSend, *instinfo.ChRecv:
							isBlocking = true
						}
						if isBlocking == false { // Not interested in non blocking operations
							continue
						}

						prim, exist := DMap[chOp.Prim()]
						if exist == false {
							prim = &DPrim{
								Primitive: chOp.Prim(),
							}
							DMap[chOp.Prim()] = prim
						}

						mapPrim[prim] = struct{}{}
					}
				} else {
					muOp, ok := instinfo.MapInst2LockerOp[inst]
					lockOp, isBlocking := muOp.(*instinfo.LockOp)
					if ok && isBlocking {
						prim, exist := DMap[lockOp.Prim()]
						if exist == false {
							prim = &DPrim{
								Primitive: lockOp.Prim(),
							}
							DMap[lockOp.Prim()] = prim
						}

						mapPrim[prim] = struct{}{}
					}
				}
			}
		}
	}

	// Update the Fn2Prims map: A calls B, then A inherits B's primitives
	// Current implementation is conservative: only when the len(potention_Bs) == 1, will A inherits B's primitives
	intCountRecursive := 0
	boolDMapUpdated := true
	for boolDMapUpdated && intCountRecursive < 7 {
		boolDMapUpdated = false
		intCountRecursive++

		for a, prims_map := range Fn2Prims {
			a_node := config.CallGraph.Nodes[a]
			if a_node == nil {
				continue
			}

			intNumOfOldPrims := len(prims_map)

			callsite2callees := make(map[ssa.CallInstruction][]*callgraph.Node)
			for _, out := range a_node.Out {
				boolCalleeExist := false
				for _, exist_callee := range callsite2callees[out.Site] {
					if exist_callee == out.Callee {
						boolCalleeExist = true
					}
				}
				if boolCalleeExist == false {
					callsite2callees[out.Site] = append(callsite2callees[out.Site], out.Callee)
				}
			}

			for _, callees := range callsite2callees {
				if len(callees) == 1 {
					mapCalleePrims := Fn2Prims[callees[0].Func]
					for prim, _ := range mapCalleePrims {
						prims_map[prim] = struct{}{}
					}
				}
			}

			intNumOfNewPrims := len(prims_map)
			if intNumOfNewPrims != intNumOfOldPrims {
				boolDMapUpdated = true
			}

			callsite2callees = nil
		}
	}

	prim2BlockInst := make(map[*DPrim][]ssa.Instruction)
	prim2UnBlockInst := make(map[*DPrim][]ssa.Instruction)
	for _, ch := range vecChan {
		prim, exist := DMap[ch]
		if exist == false {
			prim = &DPrim{
				Primitive: ch,
			}
			DMap[ch] = prim
		}
		for _, op := range ch.Recvs {
			prim2BlockInst[prim] = append(prim2BlockInst[prim], op.Inst)
			prim2UnBlockInst[prim] = append(prim2UnBlockInst[prim], op.Inst)
		}
		for _, op := range ch.Sends {
			prim2BlockInst[prim] = append(prim2BlockInst[prim], op.Inst)
			prim2UnBlockInst[prim] = append(prim2UnBlockInst[prim], op.Inst)
		}
		for _, op := range ch.Closes {
			prim2UnBlockInst[prim] = append(prim2UnBlockInst[prim], op.Inst)
		}
	}

	for ch1, prim1 := range DMap {
		for ch2, prim2 := range DMap {
			if ch1 == ch2 {
				continue
			}
			vecPrim1UnBlockInst, ok := prim2UnBlockInst[prim1]
			if !ok {
				continue
			}
			vecPrim2BlockInst, ok := prim2BlockInst[prim2]
			if !ok {
				continue
			}
			for _, prim1UnBlockInst := range vecPrim1UnBlockInst {
				for _, prim2BlockInst := range vecPrim2BlockInst {
					aPath := path.PathBetweenInst(prim2BlockInst, prim1UnBlockInst)
					if len(aPath) > 0 {
						prim1.AddOutEdge(prim2)
						if prim1.IsCircularDepend(prim2) {
							prim1.AddCircularDepend(prim2)
						}
					}
				}
			}
		}
	}

	return
}
