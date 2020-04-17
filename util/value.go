package util

import "strings"
import "github.com/system-pclub/gochecker/tools/go/ssa"

func GetValueName(v ssa.Value) string {
	var strFnName string
	fn := v.Parent()
	if fn == nil {
		strFnName = "Unknown_fn"
	} else {
		strFnName = strings.ReplaceAll(fn.String(),"(","")
		strFnName = strings.ReplaceAll(strFnName,")","")
		strFnName = strings.ReplaceAll(strFnName,"*","Ptr_")
		strFnName = strings.ReplaceAll(strFnName,"/","_Of_")
	}
	result := strFnName + "_" + v.Name()
	return result
}
