package forgetunlock

import (
	"fmt"
	"github.com/system-pclub/gochecker/tools/go/ssa"
	"go/token"
	"reflect"
	"sort"
	"strconv"
)

func HandleCall(pCall * ssa.Call) string {
	if pCall.Common().Value.Name() == "len" {
		return "len(" + HandleVarOrInst(pCall.Common().Args[0]) + ")"
	} else {
		return pCall.String()
	}
}



func HandleVarOrInst(pValue ssa.Value) string {
	if pUnOp, ok := pValue.(*ssa.UnOp); ok {
		switch pUnOp.Op {
		case token.MUL:
			return "*(" + HandleVarOrInst(pUnOp.X) + ")" //pUnOp.X.Name()
		}
	} else if pConst, ok := pValue.(*ssa.Const); ok {
		return pConst.Value.String()
	} else if pAlloc, ok := pValue.(*ssa.Alloc); ok {
		return pAlloc.Name()
	} else if pGlobal, ok := pValue.(*ssa.Global); ok {
		return pGlobal.Name()
	} else if pFunc, ok := pValue.(*ssa.Call); ok {
		return HandleCall(pFunc)
	} else if pField, ok := pValue.(* ssa.FieldAddr); ok {
		return "&" + HandleVarOrInst(pField.X) + ".changes [" + strconv.Itoa(pField.Field) + "]"
	}

	fmt.Println(reflect.TypeOf(pValue))

	panic("unhandled cases in HandleVarOrInst")


	return ""
}



func HandleBinOp(pBinOp * ssa.BinOp) string {
	aug0 := HandleVarOrInst(pBinOp.X)
	aug1 := HandleVarOrInst(pBinOp.Y)

	var strOp string

	switch pBinOp.Op {
	case token.EQL:
		strOp = "="
	case token.NEQ:
		strOp = "!="
	case token.GTR:
		strOp = ">"
	case token.GEQ:
		strOp = ">="
	case token.LSS:
		strOp = "<"
	case token.LEQ:
		strOp = "<="
	default:
		panic("default in HandleBinOp")
	}

	if strOp == "=" || strOp == "!=" {
		if pBinOp.X.Name() > pBinOp.Y.Name() {
			tmp := aug0
			aug0 = aug1
			aug1 = tmp
		}
	}

	return "(" + aug0 + " " + strOp  + " "  +  aug1 + ")"
}


func ConvertCondsToContraints(conds [] stCond) string {

	if len(conds) == 0 {
		return ""
	}

	vecSubContraints := [] string {}
	for _, cond := range conds {
		if pBinOp, ok := cond.Cond.(*ssa.BinOp); ok {
			subConstraint := HandleBinOp(pBinOp)
			if !cond.Flag {
				subConstraint = "(!" + subConstraint + ")"
			}

			vecSubContraints = append(vecSubContraints, subConstraint)
		} else {
			panic("not binop in ConvertCondsToContraints")
		}
	}



	sort.Strings(vecSubContraints)
	strResult := vecSubContraints[0]

	for _, strContraint := range vecSubContraints[1:] {
		strResult = strResult + " && " + strContraint
	}


	return strResult
}
