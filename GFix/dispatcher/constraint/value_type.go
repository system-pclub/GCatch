package constraint

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/token"
	"strings"
)

func Type_of_value(v ssa.Value) string {
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

// Now can handle == != < <= > >=
func (s *SMT_set) create_assert_body_for_BinOp(binOp *ssa.BinOp) string {

	var v1,v2 string
	if v1_const,ok := binOp.X.(*ssa.Const); ok {
		if v1_const.IsNil() {
			v1 = nil_name(v1_const)
			s.Todo = append(s.Todo,v1_const)
		} else {
			v1 = strip_quote(v1_const.Value.ExactString())
		}
	} else {
		v1 = value_name(binOp.X)
		s.Todo = append(s.Todo,binOp.X)
	}

	if v2_const,ok := binOp.Y.(*ssa.Const); ok {
		if v2_const.IsNil() {
			v2 = nil_name(v2_const)
			s.Todo = append(s.Todo,v2_const)
		} else {
			v2 = strip_quote(v2_const.Value.ExactString())
		}

	} else {
		v2 = value_name(binOp.Y)
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

// Now len() is special. The other calls are the same: declare the value as a new variable
func (s *SMT_set) handle_Call(c *ssa.Call) {
	CallCommon := c.Call

	//see if c is len()
	if CallCommon.IsInvoke() == false {
		callee_builtin,ok := CallCommon.Value.(*ssa.Builtin)
		if ok {
			if callee_builtin.Name() == "len" {
				if len(CallCommon.Args) == 1 {
					// Now we are sure this is like t2 = len(t1)
					argument := CallCommon.Args[0]
					type_in_len :=argument.Type().String()
					sort_name := type2sort(type_in_len)

					// expect type_in_len is like "[][]byte" and then sort_name is "S_S_Byte"
					// Now we see if "S_S_Byte" is already declared. If not, declare it
					s.sort_dec_or_def(sort_name)

					// Now we see if "len_S_S_Byte" is already declared. If not, declare it
					len_name := "len_" + sort_name
					new_dec_fn := Dec_fn {
						len_name,
						[]string{sort_name},
						[]string{"Int"},
					}
					if is_dec_fn_in_slice(new_dec_fn,s.Dec_fns) == false {
						s.Dec_fns = append_left_fn(s.Dec_fns,new_dec_fn)
					}

					// Now we see if t2 is declared
					c_value_name := value_name(c)
					new_dec_const := Dec_const {
						c_value_name,
						"Int",
						c,
					}
					if is_dec_const_in_slice(new_dec_const,s.Dec_consts) == false {
						s.Dec_consts = append_left_const(s.Dec_consts,new_dec_const)
					}

					// Assert t2
					assert_str := "(assert (= " + c_value_name + " (" + len_name + " " + value_name(argument) + ")))"
					new_assert := Assert(assert_str)
					s.Asserts = append_left_assert(s.Asserts,new_assert)

					//Add t1 to s.Todo
					s.Todo = append(s.Todo,argument)
					return
				}
			}
		}
	}

	//Other situations: declare t2 directly
	c_value_name := value_name(c)
	new_dec_const := Dec_const {
		c_value_name,
		"Int",
		c,
	}
	if is_dec_const_in_slice(new_dec_const,s.Dec_consts) == false {
		s.Dec_consts = append_left_const(s.Dec_consts,new_dec_const)
	}
}

func type2sort(type_name string) (sort_name string) {
	//handling slice
	core_type,layers_of_slice := strip_slice_type(type_name)
	for i:=0; i < layers_of_slice; i++ {
		sort_name += "S_"
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
	sort_name += core_sort

	//handling defined structs
	sort_name = strings.ReplaceAll(sort_name,"*","Ptr_")
	sort_name = strings.ReplaceAll(sort_name,"/","_Of_")
	return
}

// sort can be "Byte", can be "S_S_Byte"
func (s *SMT_set) sort_dec_or_def(sort string) {
	core_sort,layers_of_slice := strip_slice_sort(sort)

	// see if the sort has already been declared. Need to go through both default_sorts and s.Dec_sorts
	new_dec_sort := Dec_sort{
		core_sort,
	}
	if is_string_in_slice(core_sort,Default_sorts) == false && is_dec_sort_in_slice(new_dec_sort,s.Dec_sorts) == false  {
		s.Dec_sorts = append_left_dec_sorts(s.Dec_sorts,new_dec_sort)
	}

	if layers_of_slice == 0 {
		return
	}
	// Now we see if "S_S_Byte" is already defined. If not, define it
	new_define := Define_sort{}
	new_define.Has_input_output = true
	new_define.Input = ""
	new_define.Name = sort
	new_define.Output = core_sort
	for i:=0; i < layers_of_slice; i++ {
		new_define.Output = "Array Int " + new_define.Output
	}
	if is_def_in_slice(new_define,s.Define_sorts) == false {
		s.Define_sorts = append_left_define_sorts(s.Define_sorts,new_define)
	}
}

// Now only * is handled
func (s *SMT_set) handle_UnOp(u *ssa.UnOp) {
	switch u.Op {
	default:
		// TODO: need alias analysis for load and store. Now we only declare a new const here without any assert
		u_type := u.Type().String()
		u_sort := type2sort(u_type)
		new_dec_const := Dec_const{
			Name: value_name(u),
			Type: u_sort,
			Value: u,
		}
		s.sort_dec_or_def(u_sort)
		if is_dec_const_in_slice(new_dec_const,s.Dec_consts) == false {
			s.Dec_consts = append_left_const(s.Dec_consts,new_dec_const)
		}

	}
}

func strip_slice_sort(str string) (core string,layers int) {
	layers = strings.Count(str,"S_")
	core = strings.ReplaceAll(str,"S_","")
	return
}

func strip_slice_type(str string) (core string,layers int) {
	layers = strings.Count(str,"[]")
	core = strings.ReplaceAll(str,"[]","")
	return
}

func is_dec_fn_in_slice(new Dec_fn, slice []Dec_fn) bool {
	for _,elem := range slice {
		if new.Name == elem.Name {
			return true
		}
	}
	return false
}

func is_dec_const_in_slice(new Dec_const, slice []Dec_const) bool {
	for _,elem := range slice {
		if new.Name == elem.Name {
			return true
		}
	}
	return false
}

func is_dec_sort_in_slice(new Dec_sort, slice []Dec_sort) bool {
	for _,elem := range slice {
		if new.Name == elem.Name {
			return true
		}
	}
	return false
}

func is_def_in_slice(new Define_sort, slice []Define_sort) bool {
	for _,elem := range slice {
		if new.Name == elem.Name {
			return true
		}
	}
	return false
}

func is_assert_in_slice(assert Assert,slice []Assert) bool {
	for _,elem := range slice {
		if assert == elem {
			return true
		}
	}
	return false
}

func is_string_in_slice(str string,slice []string) bool {
	for _,elem := range slice {
		if str == elem {
			return true
		}
	}
	return false
}