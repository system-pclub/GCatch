package pointer

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/syncgraph"
	"github.com/system-pclub/GCatch/GCatch/tools/go/mypointer"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"os"
	"strconv"
	"strings"
)

func mergeAlias(vecinstValue []*instinfo.StOpValue, stPtrResult *mypointer.Result) (result map[mypointer.Label][]*instinfo.StOpValue) {
	result = make(map[mypointer.Label][]*instinfo.StOpValue)
	for _, instValue := range vecinstValue {
		labels := stPtrResult.Queries[instValue.Value].PointsTo().Labels()
		if len(labels) > 1 {
			boolNotSure := false
			strDebugNotSure := ""
			for _, label := range labels {
				if value := label.Value(); value == nil {
					continue
				}
				if parent := label.Value().Parent(); parent == nil {
					continue
				}
				pkg := label.Value().Parent().Pkg
				if pkg == nil {
					continue
				}
				pkgOfPkg := pkg.Pkg
				if pkgOfPkg == nil {
					continue
				}
				if config.IsPathIncluded(pkgOfPkg.Path()) {
					boolNotSure = true
					p := config.Prog.Fset.Position(label.Value().Pos())
					strDebugNotSure = p.Filename + ":" + strconv.Itoa(p.Line)
					break
				}
			}
			if boolNotSure {
				syncgraph.ReportNotSure()
				fmt.Println("Verification reports not sure because of inaccurate pointer analysis in:\n" + strDebugNotSure)

				os.Exit(1)
			}
		}
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
		instinfo.MapInst2ChanOp[chPrim.MakeInst] = append(instinfo.MapInst2ChanOp[chPrim.MakeInst], chPrim.Make)
	}

	for _, send := range chPrim.Sends {
		if send.Inst != nil {
			instinfo.MapInst2ChanOp[send.Inst] = append(instinfo.MapInst2ChanOp[send.Inst], send)
		}
	}
	for _, recv := range chPrim.Recvs {
		if recv.Inst != nil {
			instinfo.MapInst2ChanOp[recv.Inst] = append(instinfo.MapInst2ChanOp[recv.Inst], recv)
		}
	}
	for _, c := range chPrim.Closes {
		if c != nil {
			instinfo.MapInst2ChanOp[c.Inst] = append(instinfo.MapInst2ChanOp[c.Inst], c)
		}
	}

}