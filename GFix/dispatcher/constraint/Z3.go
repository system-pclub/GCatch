package constraint

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/path"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
)



func SMT_of_two_bb_precondition_the_same_locally(bb1,bb2 *ssa.BasicBlock) (*SMT_set,*SMT_set) {
	unhealthy_result := empty_SMT_set()
	unhealthy_result.Status = "Unhealthy"
	local_bbs := bb1.Parent().Blocks
	if len(local_bbs) == 0 { // This should never happen
		fmt.Println("Error in Output_equation_of_two_bb_precondition_the_same: bb1.Parent() has no bb")
		return unhealthy_result,unhealthy_result
	}

	if bb1 == bb2 {
		fmt.Println("true")
		return unhealthy_result,unhealthy_result
	}

	bb1_dominators := list_dominators(bb1)
	bb2_dominators := list_dominators(bb2)
	last_mutual_dominator := find_last_mutual_dominator(bb1_dominators,bb2_dominators)

	flag_bb1_dom_bb2 := false
	if last_mutual_dominator == bb1 {
		flag_bb1_dom_bb2 = true
	}
	flag_bb2_dom_bb1 := false
	if last_mutual_dominator == bb2 {
		flag_bb2_dom_bb1 = true
	}

	full_path1 := []*ssa.BasicBlock{}
	var err error
	if flag_bb1_dom_bb2 == false {
		full_path1,err = path.Find_shortest_path_locally(last_mutual_dominator,bb1)
		if err != nil {
			fmt.Println(err.Error())
			return unhealthy_result,unhealthy_result
		}
	}

	full_path2 := []*ssa.BasicBlock{}
	if flag_bb2_dom_bb1 == false {
		full_path2,err = path.Find_shortest_path_locally(last_mutual_dominator,bb2)
		if err != nil {
			fmt.Println(err.Error())
			return unhealthy_result,unhealthy_result
		}
	}

	useful_path1 := path.Delete_useless_bbs_locally(full_path1)
	useful_path2 := path.Delete_useless_bbs_locally(full_path2)

	conds1 := path.List_cond_of_path(useful_path1,bb1)
	conds2 := path.List_cond_of_path(useful_path2,bb2)
	cond1_SMT := Conds2SMT_set(conds1)
	cond2_SMT := Conds2SMT_set(conds2)

	//output.Print_insts_in_fn(bb1.Parent())
	//fmt.Println("--------")
	//fmt.Println("From last_mutual_dom-",last_mutual_dominator.Comment," to bb-",bb1.Comment)
	//cond1_SMT.Print_body()
	//fmt.Println()
	//
	//fmt.Println("From last_mutual_dom-",last_mutual_dominator.Comment," to bb-",bb2.Comment)
	//cond2_SMT.Print_body()
	//fmt.Println()

	cond1_SMT = Conds2SMT_set(conds1)
	cond2_SMT = Conds2SMT_set(conds2)


	return cond1_SMT,cond2_SMT
}




func list_dominators(target *ssa.BasicBlock) (result []*ssa.BasicBlock) {
	for _,bb := range target.Parent().Blocks {
		if bb.Dominates(target) {
			result = append(result, bb)
		}
	}
	return
}

func find_last_mutual_dominator(doms1,doms2 []*ssa.BasicBlock) *ssa.BasicBlock {
	mutual := []*ssa.BasicBlock{}
	for _,dom1 := range doms1 {
		flag_find_the_same := false
		for _,dom2 := range doms2 {
			if dom1.Index == dom2.Index {
				flag_find_the_same = true
			}
		}
		if flag_find_the_same == true {
			mutual = append(mutual,dom1)
		}
	}

	if length := len(mutual); length == 0 { // This should never happen
		return nil
	} else {
		return mutual[length - 1]
	}
}
