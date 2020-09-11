package syncgraph

import (
	"github.com/system-pclub/GCatch/config"
	"github.com/system-pclub/GCatch/instinfo"
	"github.com/system-pclub/GCatch/tools/go/ssa"
	"github.com/system-pclub/GCatch/tools/go/ssa/ssautil"
)

type DPrim struct {
	Primitive interface{} // *instinfo.Channel/ *locker.Locker
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

func GenDMap(vecChan []*instinfo.Channel) (DMap map[interface{}]*DPrim) {
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

	// TODO: Update the Fn2Prims map recursively: A calls B, then A inherits B's primitives


	return
}