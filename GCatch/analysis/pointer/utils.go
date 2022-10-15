package pointer

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/tools/go/mypointer"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"go/token"
	"strconv"
	"strings"
)

func mergeAlias(vecinstValue []*instinfo.SyncOpInfo, stPtrResult *mypointer.Result) (result map[mypointer.Label][]*instinfo.SyncOpInfo) {
	result = make(map[mypointer.Label][]*instinfo.SyncOpInfo)
	for _, instValue := range vecinstValue {
		labels := stPtrResult.Queries[instValue.Value].PointsTo().Labels()
		// This condition is incorrect, because sometimes there are repeated fields from the PointsTo result.
		// e.g., kubernetes10182 in gobench/goker
		if len(labels) > 1 {
			boolNotSure := false
			locLabel := ""

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
				locLabel = getFileAndLocString(label.Value())
				//fmt.Printf("%+v %+v %s ", getFileAndLocString(instValue.Value), label, locLabel)
				//fmt.Println(pkgOfPkg.Path())
				if config.IsPathIncluded(pkgOfPkg.Path()) {
					boolNotSure = true
					//break
				}
			}
			if boolNotSure {
				fmt.Println("Verification result is inaccurate because of possible inaccurate pointer analysis in:\n" + locLabel)
				//syncgraph.ReportNotSure()
				//os.Exit(1)
			}
		}
		for _, label := range labels {
			_, ok := result[*label]
			if ok {
				result[*label] = append(result[*label], instValue)
			} else {
				result[*label] = []*instinfo.SyncOpInfo{instValue}
			}
		}
	}

	return
}

func getFileAndLocString(label ssa.Value) string {
	return PosToFileAndLocString(label.Pos())
}

func PosToFileAndLocString(pos token.Pos) string {
	p := config.Prog.Fset.Position(pos)
	strDebugNotSure := p.Filename + ":" + strconv.Itoa(p.Line) + ":" + strconv.Itoa(p.Column)
	return strDebugNotSure
}

func boolIsInContext(v ssa.Value) bool {
	if v == nil || v.Parent() == nil || v.Parent().Pkg == nil || v.Parent().Pkg.Pkg == nil {
		return false
	}
	strPkg := v.Parent().Pkg.Pkg.Path()
	if strPkg == "context" || strings.Contains(strPkg, "golang.org/x/net/context") { // some people still
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
