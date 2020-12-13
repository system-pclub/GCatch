package pointer

import (
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/tools/go/mypointer"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"strings"
)

func mergeAlias(vecinstValue []*instinfo.StOpValue, stPtrResult *mypointer.Result) (result map[mypointer.Label][]*instinfo.StOpValue) {
	result = make(map[mypointer.Label][]*instinfo.StOpValue)
	for _, instValue := range vecinstValue {
		labels := stPtrResult.Queries[instValue.Value].PointsTo().Labels()
		for _, label := range labels {
			_, ok := result[*label]
			if ok {
				result[*label] = append(result[*label], instValue)
			} else {
				result[*label] = []*instinfo.StOpValue{instValue}
			}
		}
	}

	return
}

func boolIsInContext(v ssa.Value) bool {
	if v == nil || v.Parent() == nil || v.Parent().Pkg == nil || v.Parent().Pkg.Pkg == nil {
		return false
	}
	strPkg := v.Parent().Pkg.Pkg.Path()
	if strPkg == "context" || strings.Contains(strPkg,"golang.org/x/net/context") { // some people still
		// import golang.org/x/net/context
		return true
	}
	return false
}

func boolIsInTime(v ssa.Value) bool {
	if v == nil || v.Parent() == nil || v.Parent().Pkg == nil || v.Parent().Pkg.Pkg == nil {
		return false
	}
	strPkg := v.Parent().Pkg.Pkg.Path()
	if strPkg == "time" {
		return true
	}
	return false
}

func recordChInstToMap(chPrim *instinfo.Channel) {
	if chPrim.MakeInst != nil {
		instinfo.MapInst2ChanOp[chPrim.MakeInst] = []instinfo.ChanOp{chPrim.Make}
	}

	for _, send := range chPrim.Sends {
		if send.Inst != nil {
			instinfo.MapInst2ChanOp[send.Inst] = []instinfo.ChanOp{send}
		}
	}
	for _, recv := range chPrim.Recvs {
		if recv.Inst != nil {
			instinfo.MapInst2ChanOp[recv.Inst] = []instinfo.ChanOp{recv}
		}
	}
	for _, c := range chPrim.Closes {
		if c != nil {
			instinfo.MapInst2ChanOp[c.Inst] = []instinfo.ChanOp{c}
		}
	}

}