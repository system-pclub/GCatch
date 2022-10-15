package util

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
)

func GetPackagePath(value ssa.Value) (string, error) {
	parent := value.Parent()
	if parent == nil {
		return "", fmt.Errorf("parent of value %s is nil", value)
	}
	pkg := value.Parent().Pkg
	if pkg == nil {
		return "", fmt.Errorf("package name is empty")
	}
	pkgOfPkg := pkg.Pkg
	if pkgOfPkg == nil {
		return "", fmt.Errorf("package name is empty")
	}
	return pkgOfPkg.Path(), nil
}

func IsInstInVec(inst ssa.Instruction, vec []ssa.Instruction) bool {
	for _, elem := range vec {
		if elem == inst {
			return true
		}
	}
	return false
}

func VecFnForVecInst(vecInst []ssa.Instruction) []*ssa.Function {
	result := []*ssa.Function{}

	mapFn := make(map[*ssa.Function]struct{})
	for _, inst := range vecInst {
		mapFn[inst.Parent()] = struct{}{}
	}

	for fn, _ := range mapFn {
		result = append(result, fn)
	}

	return result
}
