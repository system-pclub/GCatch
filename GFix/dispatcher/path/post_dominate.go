package path

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
)


type Myerror struct {
	content string
}

func (e *Myerror) Error() string {
	return fmt.Sprint(e.content)
}

func Post_dominates_prepare(fn ssa.Function) error {
	var err error
	all_chains,err = List_all_exe_chain(fn,nil)
	return err
}

func Post_dominates_clean() {
	all_chains = nil
}

// Before using Post_dominates, you should call List_all_exe_chain, and use the result as the second parameter
func Post_dominates(pre,post *ssa.BasicBlock) bool {
	if pre.Parent().String() != post.Parent().String() {
		return false
	}
	
	if len(all_chains) == 0 {
		return false
	}

	for _,chain := range all_chains {

		flag_pre_in_chain := false
		flag_post_exist_after_pre := false
		for _,bb := range chain {
			if bb == pre {
				flag_pre_in_chain = true
			}
			if flag_pre_in_chain == true {
				if bb == post {
					flag_post_exist_after_pre = true
					break
				}
			}
		}
		if flag_pre_in_chain == true && flag_post_exist_after_pre == false {
			return false
		}
	}

	return true
}

