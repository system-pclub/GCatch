package forgetunlock

import (
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"github.com/system-pclub/GCatch/GCatch/util"
	"go/token"
	"strconv"
	"strings"
)


type stDecSort struct {
	Name string
}

type stDecFn struct {
	Name string
	Inputs [] string
	Outputs [] string
}

type stDefineSort struct {
	Name string
	Input string
	Output string
	HasInputOutput bool
}

type stDecConst struct {
	Name string
	Type string
	Value ssa.Value
}

type Assert string

type StSMTSet struct {
	DecSorts [] stDecSort
	DecFns  [] stDecFn
	DefineSorts  [] stDefineSort
	DecConsts [] stDecConst
	Asserts [] Assert
	FinalConst stDecConst
	FinalAssert Assert
	Todo [] ssa.Value
	Status string
	Index int
}


var (
	SMTBool = "Bool"
	SMTEQL = "="
	SMTNot = "not"
	DefaultSorts = []string{ "byte", "int", "int64", "bool", "string"}

)

var (
	CountSMTSet = 0
)

func isDecConstInSlice(new stDecConst, slice [] stDecConst) bool {
	for _,elem := range slice {
		if new.Name == elem.Name {
			return true
		}
	}
	return false
}


func isDecSortInSlice(new stDecSort, slice [] stDecSort) bool {
	for _,elem := range slice {
		if new.Name == elem.Name {
			return true
		}
	}
	return false
}

func isDefInSlice(new stDefineSort, slice [] stDefineSort) bool {
	for _,elem := range slice {
		if new.Name == elem.Name {
			return true
		}
	}
	return false
}

func isAssertInSlice(assert Assert, slice []Assert) bool {

	for _, elem := range slice {
		if assert == elem {
			return true
		}
	}
	return false
}

func isStringInSlice(str string, slice []string) bool {
	for _,elem := range slice {
		if str == elem {
			return true
		}
	}
	return false
}

func appendConstOnLeft(oldConsts [] stDecConst, newConst stDecConst) (result [] stDecConst) {
	result = append([]stDecConst{},newConst)
	for _, oldConst := range oldConsts {
		result = append(result, oldConst)
	}
	return
}

func appendDecSortsOnLeft(oldDecs [] stDecSort, newDec stDecSort) (result []stDecSort) {
	result = append([]stDecSort{}, newDec)
	for _, oldDec := range oldDecs {
		result = append(result, oldDec)
	}
	return
}

func appendDefineSortsOnLeft(oldDefs [] stDefineSort, newDef stDefineSort) (result [] stDefineSort) {
	result = append([]stDefineSort{}, newDef)
	for _, oldDef := range oldDefs {
		result = append(result,oldDef)
	}
	return
}


func appendAssertOnLeft(oldAsserts [] Assert, newAssert Assert) (result [] Assert) {
	result = append([]Assert{},newAssert)

	for _,oldAssert := range oldAsserts {
		result = append(result, oldAssert)
	}
	return
}

func getNilName(constNil *ssa.Const) string {
	typeName := constNil.Type().String()
	typeName = removeParentheses(typeName)
	return typeName + "_nil"
}

func removeParentheses(str string) string {
	str = strings.ReplaceAll(str,"(","")
	str =strings.ReplaceAll(str,")","")
	return str
}

func stripSliceType(str string) (core string, layers int) {
	layers = strings.Count(str,"[]")
	core = strings.ReplaceAll(str,"[]","")
	return
}

func getValueType(v ssa.Value) string {
	if _,ok := v.(*ssa.Alloc); ok {
		return "Alloc"
	}
	if _,ok := v.(*ssa.BinOp); ok {
		return "BinOp"
	}
	if _,ok := v.(*ssa.Builtin); ok {
		return "Builtin"
	}
	if _,ok := v.(*ssa.Call); ok {
		return "Call"
	}
	if _,ok := v.(*ssa.ChangeInterface); ok {
		return "ChangeInterface"
	}
	if _,ok := v.(*ssa.ChangeType); ok {
		return "ChangeType"
	}
	if _,ok := v.(*ssa.Const); ok {
		return "Const"
	}
	if _,ok := v.(*ssa.Convert); ok {
		return "Convert"
	}
	if _,ok := v.(*ssa.Extract); ok {
		return "Extract"
	}
	if _,ok := v.(*ssa.Field); ok {
		return "Field"
	}
	if _,ok := v.(*ssa.FieldAddr); ok {
		return "FieldAddr"
	}
	if _,ok := v.(*ssa.FreeVar); ok {
		return "FreeVar"
	}
	if _,ok := v.(*ssa.Function); ok {
		return "Function"
	}
	if _,ok := v.(*ssa.Global); ok {
		return "Global"
	}
	if _,ok := v.(*ssa.Index); ok {
		return "Index"
	}
	if _,ok := v.(*ssa.IndexAddr); ok {
		return "IndexAddr"
	}
	if _,ok := v.(*ssa.Lookup); ok {
		return "Lookup"
	}
	if _,ok := v.(*ssa.MakeChan); ok {
		return "MakeChan"
	}

	if _,ok := v.(*ssa.MakeClosure); ok {
		return "MakeClosure"
	}
	if _,ok := v.(*ssa.MakeInterface); ok {
		return "MakeInterface"
	}
	if _,ok := v.(*ssa.MakeMap); ok {
		return "MakeMap"
	}
	if _,ok := v.(*ssa.MakeSlice); ok {
		return "MakeSlice"
	}
	if _,ok := v.(*ssa.Next); ok {
		return "Next"
	}
	if _,ok := v.(*ssa.Parameter); ok {
		return "Parameter"
	}
	if _,ok := v.(*ssa.Phi); ok {
		return "Phi"
	}
	if _,ok := v.(*ssa.Range); ok {
		return "Range"
	}
	if _,ok := v.(*ssa.Select); ok {
		return "Select"
	}
	if _,ok := v.(*ssa.Slice); ok {
		return "Slice"
	}
	if _,ok := v.(*ssa.TypeAssert); ok {
		return "TypeAssert"
	}
	if _,ok := v.(*ssa.UnOp); ok {
		return "UnOp"
	}

	return "unknown"
}

func createEmptySMTSet() (result * StSMTSet) {
	result = & StSMTSet{}
	result.DecSorts = []stDecSort{}
	result.DecFns = [] stDecFn {}
	result.DefineSorts = []stDefineSort{}
	result.DecConsts = [] stDecConst{}
	result.Asserts = []Assert{}
	result.FinalAssert = Assert("")
	result.FinalConst = stDecConst{}
	result.Todo = [] ssa.Value{}
	result.Status = "Empty"
	result.Index = CountSMTSet
	CountSMTSet++

	return
}

func (s * StSMTSet) isValueDeclared(value ssa.Value) bool {
	name := util.GetValueName(value)
	for _, elem := range s.DecConsts {
		if name == elem.Name {
			return true
		}
	}

	return false
}

func (s * StSMTSet) createAssertBodyBinOp(binOp * ssa.BinOp) string {

	var v1, v2 string

	if v1Const, ok := binOp.X.(*ssa.Const); ok {
		if v1Const.IsNil() {
			v1 = getNilName(v1Const)
			s.Todo = append(s.Todo, v1Const)
		} else {
			v1 = removeParentheses(v1Const.Value.ExactString())
		}
	} else {
		v1 = util.GetValueName(binOp.X)
		s.Todo = append(s.Todo, binOp.X)
	}

	if v2Const, ok := binOp.Y.(*ssa.Const); ok {
		if v2Const.IsNil() {
			v2 = getNilName(v2Const)
			s.Todo = append(s.Todo,v2Const)
		} else {
			v2 = removeParentheses(v2Const.Value.ExactString())
		}

	} else {
		v2 = util.GetValueName(binOp.Y)
		s.Todo = append(s.Todo,binOp.Y)
	}

	var body string
	switch binOp.Op {
	case token.EQL:
		body = "(= " + v1 + " " + v2 + ")"
	case token.NEQ:
		body = "(not (= " + v1 + " " + v2 + "))"
	case token.GTR:
		body = "(> " + v1 + " " + v2 + ")"
	case token.GEQ:
		body = "(not (< " + v1 + " " + v2 + "))"
	case token.LSS:
		body = "(< " + v1 + " " + v2 + ")"
	case token.LEQ:
		body = "(not (> " + v1 + " " + v2 + "))"
	default:
		return ""
	}

	return body
}

func (s * StSMTSet) declareCommonSorts(){
	// declare Byte for byte
	decSort := stDecSort {
		"Byte",
	}
	s.DecSorts = appendDecSortsOnLeft(s.DecSorts, decSort)

	return
}

func type2Sort(strTypeName string) (strSortName string) {
	//handling slice
	core_type, layers_of_slice := stripSliceType(strTypeName)

	for i:=0; i < layers_of_slice; i++ {
		strSortName += "S_"
	}
	var core_sort string

	switch core_type {
	case "byte":
		core_sort = "Byte"
	case "int":
		core_sort = "Int"
	case "int64":
		core_sort = "Int"
	case "string":
		core_sort = "String"
	case "bool":
		core_sort = "Bool"
	default:
		core_sort = "Sort_" + core_type
	}
	strSortName += core_sort

	//handling defined structs
	strSortName = strings.ReplaceAll(strSortName,"*","Ptr_")
	strSortName = strings.ReplaceAll(strSortName,"/","_Of_")
	return
}



func (s * StSMTSet) sortDecOrDef(sort string) {

	coreSort, numLayersOfSlice := stripSliceType(sort)

	// see if the sort has already been declared. Need to go through both default_sorts and s.Dec_sorts
	new_dec_sort := stDecSort{
		coreSort,
	}

	if isStringInSlice(coreSort, DefaultSorts) == false && isDecSortInSlice(new_dec_sort, s.DecSorts) == false  {
		s.DecSorts = appendDecSortsOnLeft(s.DecSorts, new_dec_sort)
	}

	if numLayersOfSlice == 0 {
		return
	}
	// Now we see if "S_S_Byte" is already defined. If not, define it
	newDefine := stDefineSort{ }
	newDefine.HasInputOutput = true
	newDefine.Input = ""
	newDefine.Name = sort
	newDefine.Output = coreSort

	for i:=0; i < numLayersOfSlice; i++ {
		newDefine.Output = "Array Int " + newDefine.Output
	}

	if isDefInSlice(newDefine, s.DefineSorts) == false {
		s.DefineSorts =  appendDefineSortsOnLeft(s.DefineSorts, newDefine)
	}
}

// Now only * is handled
func (s * StSMTSet) handleUnOp(u *ssa.UnOp) {
	switch u.Op {
	default:
		// TODO: need alias analysis for load and store. Now we only declare a new const here without any assert
		u_type := u.Type().String()
		u_sort := type2Sort(u_type)
		new_dec_const := stDecConst{
			Name: getValueType(u),
			Type: u_sort,
			Value: u,
		}
		s.sortDecOrDef(u_sort)
		if isDecConstInSlice(new_dec_const, s.DecConsts) == false {
			s.DecConsts = appendConstOnLeft(s.DecConsts, new_dec_const)
		}

	}
}

func (s * StSMTSet) processTodo() {
	for len(s.Todo) > 0 {
		value := s.Todo[0]
		if s.isValueDeclared(value) == false {
			valueType := getValueType(value)
			switch valueType {
			case "Call":
				panic("Call in processTodo")
				//value_Call,_ := value.(*ssa.Call)
				//s.handle_Call(value_Call)
			case "UnOp":
				//panic("UnOp in processTodo")
				value_UnOp,_ := value.(*ssa.UnOp)
				s.handleUnOp(value_UnOp)
			default:
				s.Status = "Unhealthy"
				return
			}
		}
		new_Todo := [] ssa.Value {}
		for i, v := range s.Todo {
			if i > 0 {
				new_Todo = append(new_Todo, v)
			}
		}
		s.Todo = new_Todo
	}
}


func Conds2SMTSet(vecCond [] stCond) (result * StSMTSet) {
	result = createEmptySMTSet()
	if len(vecCond) == 0 {
		result.Status = "Unhealthy"
		return
	} else {
		result.Status = "Generating"
	}

	for i, _ := range vecCond {
		cond := vecCond[len(vecCond) - 1 - i] // need to do this from tail to head
		strCondName := util.GetValueName(cond.Cond)

		//generate declare-const
		newConst := stDecConst {
			Name: strCondName,
			Type: SMTBool,
			Value: cond.Cond,
		}
		if isDecConstInSlice(newConst, result.DecConsts) {
			continue
		}
		result.DecConsts = appendConstOnLeft(result.DecConsts, newConst)

		//generate assert
		var newAssert Assert
		binOp, ok := cond.Cond.(*ssa.BinOp)

		if ok {
			newAssertBody := result.createAssertBodyBinOp(binOp)
			newAssert = Assert("(assert (= " + strCondName + " " + newAssertBody + " ))")
		} else { // Now all conds I have seen are of type *ssa.BinOp

			strErr := ";Unhandled assert for:" + cond.Cond.String()
			condInst, ok := cond.Cond.(ssa.Instruction)
			if ok {
				strErr += " = " + condInst.String()
			} else {
				strErr += " in func " + cond.Cond.Parent().String()
			}
			newAssert = Assert(strErr)
		}

		result.Asserts = appendAssertOnLeft(result.Asserts,newAssert)

		//handle Todo_list, declaring all values we need
		result.processTodo()
		if result.Status == "Unhealthy" {
			return
		}
	}

	result.declareCommonSorts()

	//generate final bool const and its assert, which is linking
	result.FinalConst = stDecConst {
		Name: "Final" + strconv.Itoa(result.Index),
		Type: "Bool",
		Value: nil,
	}

	result.DecConsts = append(result.DecConsts, result.FinalConst)

	var strFinalAssertBody string
	for _, cond := range vecCond {
		if cond.Flag == true {
			strFinalAssertBody += util.GetValueName(cond.Cond) + " "
		} else {
			strFinalAssertBody += "(not " + util.GetValueName(cond.Cond) + ") "
		}
	}
	result.FinalAssert = Assert("(assert (= " + result.FinalConst.Name + " (and true " + strFinalAssertBody + ")))")
	result.Asserts = append(result.Asserts,result.FinalAssert)
	result.Status = "Finished"

	return result
}
