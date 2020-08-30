package pointer

import (
	"fmt"
	"github.com/system-pclub/GCatch/config"
	"github.com/system-pclub/GCatch/instinfo"
	"github.com/system-pclub/GCatch/tools/go/callgraph"
	"github.com/system-pclub/GCatch/tools/go/mypointer"
	"github.com/system-pclub/GCatch/tools/go/ssa"
	"github.com/system-pclub/GCatch/tools/go/ssa/ssautil"
)

// AnalyzeAllSyncOp first finds all sync operations and corresponding values, which will be returned
// It then runs the pointer analysis for each value, and return the result
func AnalyzeAllSyncOp() (*mypointer.Result, []*instinfo.StOpValue) {
	vecStOpValue := []*instinfo.StOpValue{}
	for fn, _ := range ssautil.AllFunctions(config.Prog) {
		if fn == nil {
			continue
		}
		// Note that we scan every available functions here, because we don't know where a chan will be passed to
		for _,bb := range fn.Blocks {
			for _,inst := range bb.Instrs {
				// case 1: traditional
				v, comment := instinfo.ScanInstFindLockerValue(inst)
				if v != nil {
					newStOpValue := &instinfo.StOpValue{
						Inst:    inst,
						Value:   v,
						Comment: comment,
					}
					vecStOpValue = append(vecStOpValue, newStOpValue)
					continue
				}

				// case 2: channel
				chs, comments := instinfo.ScanInstFindChanValue(inst)
				for i, ch := range chs {
					if ch == nil {
						continue
					}
					newStOpValue := &instinfo.StOpValue{
						Inst:    inst,
						Value:   chs[i],
						Comment: comments[i],
					}
					vecStOpValue = append(vecStOpValue, newStOpValue)
				}
			}
		}
	}

	queries := make(map[ssa.Value]struct{})
	for _, stOpValue := range vecStOpValue {
		queries[stOpValue.Value] = struct{}{}
	}
	cfg := &mypointer.Config{
		OLDMains:        nil,
		Prog:            config.Prog,
		Reflection:      false,
		BuildCallGraph:  true,
		Queries:         queries,
		IndirectQueries: nil,
		Log:             nil,
	}
	stPtrResult, err := mypointer.Analyze(cfg, config.CallGraph)
	if err != nil {
		fmt.Println("Error when querying all channel values:\n",err.Error())
		return nil, nil
	}

	// Update config.Callgraph, and create a map from instruction to all its corresponding out edges in CallGraph
	config.CallGraph = stPtrResult.CallGraph

	config.Inst2CallSite = make(map[ssa.CallInstruction]map[*callgraph.Edge]bool)
	for _, node := range config.CallGraph.Nodes {
		for _, out := range node.Out {
			mapCallSites, boolExist := config.Inst2CallSite[out.Site]
			if !boolExist {
				mapCallSites = make(map[*callgraph.Edge]bool)
				config.Inst2CallSite[out.Site] = mapCallSites
			}

			mapCallSites[out] = true
		}
	}

	return stPtrResult, vecStOpValue
}