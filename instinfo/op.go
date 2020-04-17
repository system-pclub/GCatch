package instinfo

import (
	"github.com/system-pclub/gochecker/tools/go/ssa"
	"go/types"
)

func GetCallName(call *ssa.CallCommon) string {

	if call.IsInvoke() {
		return call.String()
	}

	switch v := call.Value.(type) {
	case *ssa.Function:
		fn, ok := v.Object().(*types.Func)
		if !ok {
			return ""
		}
		return fn.FullName()
	case *ssa.Builtin:
		return v.Name()
	}
	return ""
}
