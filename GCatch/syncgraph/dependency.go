package syncgraph

import (
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/path"
	"github.com/system-pclub/GCatch/GCatch/tools/go/callgraph"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa/ssautil"
	"github.com/system-pclub/GCatch/GCatch/util"
)

type DPrim struct {
	Primitive interface{} // *instinfo.Channel or *instinfo.Locker
	Out []*DEdge
	In []*DEdge
	Circular_depend []*DEdge
	place string
	depend_on_places []string
}

type DEdge struct {
	Caller *DPrim
	Callee *DPrim
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
				mapChOp, ok := instinfo.MapInst2ChanOp[inst]
				if ok {
					for chOp, boolExist := range mapChOp {
						if boolExist == false {
							continue
						}
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
	for boolDMapUpdated && intCountRecursive < config.MAX_LCA_LAYER {
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

	for _, ch := range vecChan {
		prim, exist := DMap[ch]
		if exist == false {
			prim = &DPrim{
				Primitive: ch,
			}
			DMap[ch] = prim
		}

		vecAllOpInst := []ssa.Instruction{}
		for _, op := range ch.AllOps() {
			vecAllOpInst = append(vecAllOpInst, op.Instr())
		}
		mapLCA2Chains, err := path.FindLCA(util.VecFnForVecInst(vecAllOpInst), config.MAX_LCA_LAYER)
		if err != nil {
			continue
		}

		// TODO: continue this part after I write the code for FindLCA
		_ = mapLCA2Chains
	}


	return
}