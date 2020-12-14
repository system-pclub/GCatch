package constraint

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/path"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/token"
)

// This is my own solver that can compare two ssa.Value. It can only consider part of types of ssa.Value and only works when two values are almost the same
func Is_Value_the_same(a,b ssa.Value) bool {
	if a.String() == b.String() {
		return true
	}
	
	a_type := Type_of_value(a)
	b_type := Type_of_value(b)
	if a_type != b_type {
		return false
	}

	//Type().String() is the type defined by package types
	if a.Type().String() != b.Type().String() {
		return false
	}

	switch a_type {
	case "BinOp":
		return is_BinOp_the_same(*a.(*ssa.BinOp),*b.(*ssa.BinOp))
	case "Const":
		return is_Const_the_same(*a.(*ssa.Const),*b.(*ssa.Const))
	case "Call":
		return is_Call_the_same(*a.(*ssa.Call),*b.(*ssa.Call))
	case "UnOp":
		return is_UnOp_the_same(*a.(*ssa.UnOp),*b.(*ssa.UnOp))
	case "FieldAddr":
		return is_FieldAddr_the_same(*a.(*ssa.FieldAddr),*b.(*ssa.FieldAddr))
	}

	return false
}

func is_FieldAddr_the_same(a,b ssa.FieldAddr) bool {
	return a.Field == b.Field && Is_Value_the_same(a.X,b.X)
}

func is_UnOp_the_same(a,b ssa.UnOp) bool {
	return a.Op == b.Op && Is_Value_the_same(a.X,b.X) && a.CommaOk == b.CommaOk
}

func is_Call_the_same(a,b ssa.Call) bool {
	if len(a.Call.Args) != len(b.Call.Args) {
		return false
	}

	is_Value_the_same := true
	for i,_ := range (a.Call.Args) {
		if Is_Value_the_same(a.Call.Args[i],b.Call.Args[i]) == false {
			is_Value_the_same = false
		}
	}

	is_func_the_same := false

	if a.Call.IsInvoke() && b.Call.IsInvoke() {
		is_func_the_same = ( a.Call.Value.String() == b.Call.Value.String() ) && ( a.Call.Method == b.Call.Method)
	} else if a.Call.IsInvoke() == false && b.Call.IsInvoke() == false {
		is_func_the_same = a.Call.Value.String() == b.Call.Value.String()
	}

	return is_func_the_same && is_Value_the_same
}

func is_Const_the_same(a,b ssa.Const) bool {
	return a.String() == b.String()
}

func is_BinOp_the_same(a,b ssa.BinOp) bool { //TODO: only can deal with some of the operations
	if (a.Op == token.EQL && b.Op == token.EQL) || (a.Op == token.NEQ && b.Op == token.NEQ) {
		return (Is_Value_the_same(a.X,b.X) && Is_Value_the_same(a.Y,b.Y) ) ||
			(Is_Value_the_same(a.X,b.Y) && Is_Value_the_same(a.Y,b.X) )
	}

	if (a.Op == token.LSS && b.Op == token.LSS) || (a.Op == token.GTR && b.Op == token.GTR) {
		return Is_Value_the_same(a.X,b.X) && Is_Value_the_same(a.Y,b.Y)
	}

	if (a.Op == token.LSS && b.Op == token.GTR) || (a.Op == token.GTR && b.Op == token.LSS) {
		return Is_Value_the_same(a.X,b.Y) && Is_Value_the_same(a.Y,b.X)
	}

	if (a.Op == token.LEQ && b.Op == token.LEQ) || (a.Op == token.GEQ && b.Op == token.GEQ) {
		return Is_Value_the_same(a.X,b.X) && Is_Value_the_same(a.Y,b.Y)
	}

	if (a.Op == token.LEQ && b.Op == token.GEQ) || (a.Op == token.GEQ && b.Op == token.LEQ) {
		return Is_Value_the_same(a.X,b.Y) && Is_Value_the_same(a.Y,b.X)
	}

	return false
}

func compare_if(a,b *ssa.If) bool {
	return Is_Value_the_same(a.Cond,b.Cond)
}

//TODO: Now it can only deal with circumstances like: if(cond1) then bb1 endif; ... no return ... ; if(cond1) then bb2 endif;
func Is_two_bb_definitely_both_happen(a,b ssa.BasicBlock) bool {

	if a.Parent().String() != b.Parent().String() {
		return false
	}
	if len(a.Preds) != 1 || len(b.Preds) != 1 {
		return false
	}

	a_pred := a.Preds[0]
	b_pred := a.Preds[0]

	err := path.Post_dominates_prepare(*a.Parent())
	defer path.Post_dominates_clean()
	if err != nil {
		return false
	}
	if path.Post_dominates(a_pred,b_pred) || path.Post_dominates(b_pred,a_pred) {
		a_pred_if := find_if_of_bb(a_pred)
		b_pred_if := find_if_of_bb(b_pred)
		return compare_if(a_pred_if,b_pred_if)
	}

	return false
}

func find_if_of_bb(bb *ssa.BasicBlock) *ssa.If {
	if len(bb.Instrs) == 0 {
		return nil
	}
	inst_last := bb.Instrs[len(bb.Instrs) - 1]
	inst_if,ok := inst_last.(*ssa.If)
	if !ok {
		return nil
	}

	return inst_if
}