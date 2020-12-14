package constraint

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/path"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"strconv"
	"strings"
)

type SMT_set struct {
	Dec_sorts []Dec_sort
	Dec_fns  []Dec_fn
	Define_sorts  []Define_sort
	Dec_consts []Dec_const
	Asserts []Assert
	Final_const Dec_const
	Final_Assert Assert
	Todo []ssa.Value
	Status string
	Index int
}

type Dec_sort struct {
	Name string
}

type Define_sort struct {
	Name string
	Input string
	Output string
	Has_input_output bool
}

type Dec_fn struct {
	Name string
	Inputs []string
	Outputs []string
}

type Dec_const struct {
	Name string
	Type string
	Value ssa.Value
}

type Assert string

var (
	SMT_Bool = "Bool"
	SMT_EQL = "="
	SMT_not = "not"
	Default_sorts = []string{ "byte", "int", "int64", "bool", "string"}

)

var (
	count_SMT_set = 0

)

func empty_SMT_set() (result *SMT_set) {
	result = &SMT_set{}
	result.Dec_sorts = []Dec_sort{}
	result.Dec_fns = []Dec_fn{}
	result.Define_sorts = []Define_sort{}
	result.Dec_consts = []Dec_const{}
	result.Asserts = []Assert{}
	result.Final_Assert = Assert("")
	result.Final_const = Dec_const{}
	result.Todo = []ssa.Value{}
	result.Status = "Empty"
	result.Index = count_SMT_set
	count_SMT_set++

	return
}

func Conds2SMT_set(conds []path.Cond) (result *SMT_set) {
	result = empty_SMT_set()
	if len(conds) == 0 {
		result.Status = "Unhealthy"
		return
	} else {
		result.Status = "Generating"
	}



	for i,_ := range conds{
		cond := conds[len(conds) - 1 - i] // need to do this from tail to head
		const_name_cond := value_name(cond.Cond)

		//generate declare-const
		new_const := Dec_const{
			Name: const_name_cond,
			Type: SMT_Bool,
			Value: cond.Cond,
		}
		if is_dec_const_in_slice(new_const,result.Dec_consts){
			continue
		}
		result.Dec_consts = append_left_const(result.Dec_consts,new_const)

		//generate assert
		var new_assert Assert
		cond_BinOp,ok := cond.Cond.(*ssa.BinOp)
		if ok {
			new_assert_body := result.create_assert_body_for_BinOp(cond_BinOp)
			new_assert = Assert("(assert (= " + const_name_cond + " " + new_assert_body + " ))")
		} else { // Now all conds I have seen are of type *ssa.BinOp

			err_str := ";Unhandled assert for:" + cond.Cond.String()
			cond_inst,ok := cond.Cond.(ssa.Instruction)
			if ok {
				err_str += " = " + cond_inst.String()
			} else {
				err_str += " in func " + cond.Cond.Parent().String()
			}
			new_assert = Assert(err_str)
		}

		result.Asserts = append_left_assert(result.Asserts,new_assert)

		//handle Todo_list, declaring all values we need
		result.do_Todo()
		if result.Status == "Unhealthy" {
			return
		}

	}

	//declare common sorts like Byte
	result.declare_common_sorts()

	//generate final bool const and its assert, which is linking
	result.Final_const = Dec_const{
		Name: "Final"+strconv.Itoa(result.Index),
		Type: "Bool",
		Value: nil,
	}
	result.Dec_consts = append(result.Dec_consts,result.Final_const)
	var final_assert_body string
	for _,cond := range conds {
		if cond.Flag == true {
			final_assert_body += value_name(cond.Cond) + " "
		} else {
			final_assert_body += "(not " + value_name(cond.Cond) + ") "
		}
	}
	result.Final_Assert = Assert("(assert (= " + result.Final_const.Name + " (and true " + final_assert_body + ")))")
	result.Asserts = append(result.Asserts,result.Final_Assert)
	result.Status = "Finished"

	return
}

func Union_two_SMTs(s1,s2 *SMT_set) (result *SMT_set) {
	result = empty_SMT_set()

	result.Dec_sorts = s1.Dec_sorts
	for _,dec_sort := range s2.Dec_sorts {
		if is_dec_sort_in_slice(dec_sort,result.Dec_sorts) == false {
			result.Dec_sorts = append(result.Dec_sorts,dec_sort)
		}
	}

	result.Define_sorts = s1.Define_sorts
	for _,def := range s2.Define_sorts {
		if is_def_in_slice(def,result.Define_sorts) == false {
			result.Define_sorts = append(result.Define_sorts,def)
		}
	}

	result.Dec_fns = s1.Dec_fns
	for _,dec_fn := range s2.Dec_fns {
		if is_dec_fn_in_slice(dec_fn,result.Dec_fns) == false {
			result.Dec_fns = append(result.Dec_fns,dec_fn)
		}
	}

	result.Dec_consts = s1.Dec_consts
	for _,dec_const := range s2.Dec_consts {
		if is_dec_const_in_slice(dec_const,result.Dec_consts) == false {
			result.Dec_consts = append(result.Dec_consts,dec_const)
		}
	}

	result.Asserts = s1.Asserts
	for _,assert := range s2.Asserts {
		if is_assert_in_slice(assert,result.Asserts) == false {
			result.Asserts = append(result.Asserts,assert)
		}
	}

	//See if there are alias in const we declared, if there are, add asserts
	for _,const1 := range result.Dec_consts {
		for _,const2 := range result.Dec_consts{
			if const1.Name == const2.Name || const1.Value == nil || const2.Value == nil {
				continue
			}
			if Is_Value_the_same(const1.Value,const2.Value) {
				new_assert_str := "(assert (= " + const1.Name + " " + const2.Name + "))"
				result.Asserts = append(result.Asserts,Assert(new_assert_str))
			}
		}
	}

	result.Final_const = Dec_const{
		Name: "Final" + strconv.Itoa(result.Index),
		Type: "Bool",
		Value: nil,
	}
	result.Dec_consts = append(result.Dec_consts,result.Final_const)

	final_assert_str := "(assert (= " + result.Final_const.Name +
		" (and (=> " + s1.Final_const.Name + " " + s2.Final_const.Name + ") " +
		"(=> " + s1.Final_const.Name + " " + s2.Final_const.Name + "))))"
	result.Final_Assert = Assert(final_assert_str)
	result.Asserts = append(result.Asserts, result.Final_Assert)

	result.Status = "Finished"
	return
}

// Print_body prints everything in order except the last assert and "check-sat"
func (s *SMT_set) Print_body() {
	dec_sort_strs := s.gen_dec_sort()
	define_sort_strs := s.gen_define_sort()
	dec_fn_strs := s.gen_dec_fn()
	dec_const_strs := s.gen_dec_const()
	assert_strs := s.gen_assert()
	for _,str := range dec_sort_strs {
		fmt.Println(str)
	}
	for _,str := range define_sort_strs {
		fmt.Println(str)
	}
	for _,str := range dec_fn_strs {
		fmt.Println(str)
	}
	for _,str := range dec_const_strs {
		fmt.Println(str)
	}
	for _,str := range assert_strs {
		fmt.Println(str)
	}
	return
}

func (s *SMT_set) Print_tail() {
	//final_assert is checking whether the opposite of final const is possibly true (satisfiable),
	// which is the same as whether the final const is always true (valid)
	final_assert := Assert( "(assert (not " + s.Final_const.Name + "))")
	fmt.Println(final_assert)
	fmt.Println("(check-sat)")
}

func (s *SMT_set) gen_dec_sort() (result []string) {
	for _,dec_sort := range s.Dec_sorts {
		result = append(result,"(declare-sort " + dec_sort.Name + ")")
	}
	return
}

func (s *SMT_set) gen_define_sort() (result []string) {
	for _,def := range s.Define_sorts {
		str := "(define-sort " + def.Name + " (" + def.Input + ") (" + def.Output + "))"
		result = append(result,str)
	}
	return
}

func (s *SMT_set) gen_dec_fn() (result []string) {
	for _,dec_fn := range s.Dec_fns {
		str := "(declare-fun " + dec_fn.Name + " ("
		for _,input := range dec_fn.Inputs {
			str += input
		}
		str += ") ("
		for _,output := range dec_fn.Outputs {
			str += output
		}
		str += "))"
		result = append(result,str)
	}
	return
}

func (s *SMT_set) gen_dec_const() (result []string) {
	for _,dec_const := range s.Dec_consts {
		str := "(declare-const " + dec_const.Name + " " + dec_const.Type + ")"
		result = append(result,str)
	}
	return
}

func (s *SMT_set) gen_assert() (result []string) {
	for _,assert := range s.Asserts {
		result = append(result,string(assert))
	}
	return
}

func (s *SMT_set) declare_common_sorts(){
	// declare Byte for byte
	dec_sort_byte := Dec_sort{
		"Byte",
	}

	s.Dec_sorts = append_left_dec_sorts(s.Dec_sorts,dec_sort_byte)



	return
}

func (s *SMT_set) do_Todo() {
	for len(s.Todo) > 0 {
		value := s.Todo[0]
		if s.is_value_already_declared(value) == false {
			value_struct := Type_of_value(value)
			switch value_struct {
			case "Call":
				value_Call,_ := value.(*ssa.Call)
				s.handle_Call(value_Call)
			case "UnOp":
				value_UnOp,_ := value.(*ssa.UnOp)
				s.handle_UnOp(value_UnOp)
			default:
				s.Status = "Unhealthy"
				return
			}
		}
		new_Todo := []ssa.Value{}
		for i,v := range s.Todo {
			if i > 0 {
				new_Todo = append(new_Todo,v)
			}
		}
		s.Todo = new_Todo
	}
}

func (s *SMT_set) is_value_already_declared(value ssa.Value) bool {
	name := value_name(value)
	for _,elem := range s.Dec_consts {
		if name == elem.Name {
			return true
		}
	}
	return false
}

func nil_name(const_nil *ssa.Const) string {
	type_name := const_nil.Type().String()
	type_name = strip_parentheses(type_name)
	return type_name + "_nil"
}

func strip_quote(str string) string {
	str = strings.ReplaceAll(str,"\"","")
	return str
}

func value_name(v ssa.Value) string {
	var fn_name string
	parent_fn := v.Parent()
	if parent_fn == nil {
		fn_name = "Unknown_fn"
	} else {
		fn_name = strip_parentheses(parent_fn.String())
		fn_name = strings.ReplaceAll(fn_name,"*","Ptr_")
		fn_name = strings.ReplaceAll(fn_name,"/","_Of_")
	}
	result := fn_name + "_" + v.Name()
	return result
}

func append_left_dec_sorts(old_decs []Dec_sort, new_dec Dec_sort) (result []Dec_sort) {
	result = append([]Dec_sort{},new_dec)
	for _,old_dec := range old_decs {
		result = append(result,old_dec)
	}
	return
}

func append_left_define_sorts(old_defs []Define_sort, new_def Define_sort) (result []Define_sort) {
	result = append([]Define_sort{},new_def)
	for _,old_def := range old_defs {
		result = append(result,old_def)
	}
	return
}

func append_left_assert(old_asserts []Assert,new_assert Assert) (result []Assert) {
	result = append([]Assert{},new_assert)
	for _,old_assert := range old_asserts {
		result = append(result,old_assert)
	}
	return
}

func append_left_const(old_consts []Dec_const, new_const Dec_const) (result []Dec_const) {
	result = append([]Dec_const{},new_const)
	for _,old_const := range old_consts {
		result = append(result,old_const)
	}
	return
}

func append_left_fn(old_fns []Dec_fn, new_fn Dec_fn) (result []Dec_fn) {
	result = append([]Dec_fn{},new_fn)
	for _,old_fn := range old_fns {
		result = append(result,old_fn)
	}
	return
}


func strip_parentheses(str string) string {
	str = strings.ReplaceAll(str,"(","")
	str =strings.ReplaceAll(str,")","")
	return str
}

func append_check_sat(strs []string) []string {
	check_sat := "(check-sat)"
	return append(strs,check_sat)
}

