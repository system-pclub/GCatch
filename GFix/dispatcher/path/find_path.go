package path

import (
	"errors"
	"github.com/system-pclub/GCatch/GFix/dispatcher/mycallgraph"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"strconv"
)

var all_chains map[int][]*ssa.BasicBlock
var count_recursive int

const Max_recursive int = 10000
const Start_length int = 10000

type index_pair struct {
	index1 int
	index2 int
	index2_minus_1 int
}

type Cond struct {
	Cond ssa.Value
	Flag bool
}

// Find_path_by_call_chain can return a slice of basic blocks from bb1 to bb2, depending on a call-chain from bb1 to bb2.
// This slice is in order and minimum. bb1 and bb2's parents are the head and tail of call-chain
func Find_path_by_call_chain (bb1,bb2 *ssa.BasicBlock, call_chain []*mycallgraph.Call_node) ( []*ssa.BasicBlock,  error) {

	var err error
	if len(call_chain) == 0 {
		err = errors.New("The call_chain is empty")
		return nil,err
	}
	if bb1.Parent().String() != call_chain[0].Fn.String() {
		err = errors.New("The first bb is not in the head of call_chain")
		return nil,err
	} else if bb2.Parent().String() != call_chain[len(call_chain) - 1].Fn.String() {
		err = errors.New("The second bb is not in the tail of call_chain")
		return nil,err
	}

	if len(call_chain) == 1 {
		return Find_shortest_path_locally(bb1,bb2)
	}

	// first we get the inst between each node of call_chain
	call_inst_list := []ssa.Instruction{}
	for i,this_node := range call_chain {
		if i == len(call_chain) - 1 {
			break
		}
		next_node := call_chain[i+1]
		var call_inst ssa.Instruction
		for inst,node_list := range this_node.Call_map {
			for _,node := range node_list {
				if node == next_node {
					call_inst = inst
				}
			}
		}
		if call_inst == nil {
			err = errors.New("Fail to find the call_inst from " + this_node.Fn.String() + " to " + next_node.Fn.String())
			return nil,err
		}
		call_inst_list = append(call_inst_list,call_inst)
	}

	path := []*ssa.BasicBlock{}
	for i,node := range call_chain {
		if i == 0 { // head node
			new_path,err := Find_shortest_path_locally(bb1,call_inst_list[i].Block())
			if err != nil {
				return nil,err
			}
			for _,new_bb := range new_path {
				path = append(path,new_bb)
			}

		} else if i == len(call_chain) - 1 { // tail node
			if len(node.Fn.Blocks) == 0 {
				err = errors.New("Empty fn:" + node.Fn.String())
				return nil,err
			}
			new_path,err := Find_shortest_path_locally(node.Fn.Blocks[0],bb2)
			if err != nil {
				return nil,err
			}
			for _,new_bb := range new_path {
				path = append(path,new_bb)
			}
		} else {
			if len(node.Fn.Blocks) == 0 {
				err = errors.New("Empty fn:" + node.Fn.String())
				return nil,err
			}
			new_path,err := Find_shortest_path_locally(node.Fn.Blocks[0],call_inst_list[i].Block())
			if err != nil {
				return nil,err
			}
			for _,new_bb := range new_path {
				path = append(path,new_bb)
			}
		}
	}

	return path,nil
}

func Delete_useless_bbs(full_path []*ssa.BasicBlock) (result []*ssa.BasicBlock) {
	result = []*ssa.BasicBlock{}
	if len(full_path) == 0 {
		return
	}

	previous_fn_string := "INITFUNCTIONNAME"
	previous_sub_path := []*ssa.BasicBlock{}
	sub_path_slice := [][]*ssa.BasicBlock {}
	for i,bb := range full_path {
		if bb.Parent().String() != previous_fn_string { // this is the beginning of a new function
			if i != 0 { // if-then when this is not the beginning bb
				sub_path_slice = append(sub_path_slice,previous_sub_path)
			}
			previous_sub_path = []*ssa.BasicBlock{bb}
			previous_fn_string = bb.Parent().String()
		} else { 										// still in the old function
			previous_sub_path = append(previous_sub_path,bb)
		}

		if i == len(full_path) - 1 {					// this is the ending bb
			sub_path_slice = append(sub_path_slice,previous_sub_path)
		}
	}

	for _,sub_path := range sub_path_slice {
		new_sub_path := Delete_useless_bbs_locally(sub_path)
		for _,bb := range new_sub_path {
			result = append(result,bb)
		}
	}
	return
}

// make sure that all bbs in full_path are of the same fn
func Delete_useless_bbs_locally(full_path []*ssa.BasicBlock) (result []*ssa.BasicBlock) {
	result = []*ssa.BasicBlock{}
	if len(full_path) == 0 {
		return
	}
	result = full_path

	err := Post_dominates_prepare(*full_path[0].Parent())
	if err != nil {
		return
	}
	defer Post_dominates_clean()

	changed := true
	for changed == true {
		result,changed = try_delete_useless_bbs(result)
	}

	return
}

func try_delete_useless_bbs(ori_path []*ssa.BasicBlock) (result []*ssa.BasicBlock,changed bool) {
	index1 := -1
	index2 := -1
	for i,_ := range ori_path {
		for j := len(ori_path) - 1 - i ; j > i; j-- {
			if Post_dominates(ori_path[i],ori_path[j]) {
				index1 = i
				index2 = j
			}
		}
	}
	if index1 == -1 || index2 == -1 {
		result = ori_path
		changed = false
		return
	} else {
		result = []*ssa.BasicBlock{}
		changed = true
		for i,_ := range ori_path {
			if i < index1 || i >= index2 {
				result = append(result,ori_path[i])
			}
		}
		return
	}
}

// can return paths of basic blocks from bb1 to bb2, which are in the same function
func Find_paths_locally (bb1,bb2 *ssa.BasicBlock) (paths [][]*ssa.BasicBlock,err error) {
	paths = *new([][]*ssa.BasicBlock)
	paths_hashs := []string{}
	err = nil
	if bb1.Parent().String() != bb2.Parent().String() {
		err = errors.New("Two bbs are not of the same function")
		return
	}
	if bb1 == bb2 {
		paths = [][]*ssa.BasicBlock{{bb1}}
		return
	}

	chains,err := List_all_exe_chain(*bb1.Parent(),bb2)
	if err != nil {
		return
	}

	for _,chain := range chains {
		// In this chain, any bb can appear up to 2 times. The location and times of appearance of either of bb1 and bb2 can be very various.
		// First we guarantee both of them exist, and there is at least one bb1 before bb2
		// We record indexes of bb1 and bb2 by the way
		flag_has_bb1 := false
		flag_has_bb2 := false
		flag_has_bb1_before_bb2 := false
		index_bb1 := []int{}
		index_bb2 := []int{}
		for i,bb := range chain {
			if bb == bb1 {
				flag_has_bb1 = true
				index_bb1 = append(index_bb1,i)
			}
			if bb == bb2 {
				flag_has_bb2 = true
				index_bb2 = append(index_bb2,i)
				if flag_has_bb1 {
					flag_has_bb1_before_bb2 = true
				}
			}
		}
		if flag_has_bb1 == false || flag_has_bb2 == false || flag_has_bb1_before_bb2 == false {
			continue
		}

		//pairs contain all sets of index of bb1 and bb2, and bb1 is before bb2
		pairs := []index_pair{}
		for _,index1 := range index_bb1 {
			for _,index2 := range index_bb2 {
				if index2_minus_1 := index2 - index1; index2_minus_1 > 0 {
					new_pair := index_pair{
						index1,
						index2,
						index2_minus_1,
					}
					pairs = append(pairs,new_pair)
				}
			}
		}
		if len(pairs) == 0 { // This should always be false. Write just in case
			continue
		}

		pair_loop:
		for _,pair := range pairs {
			new_path := []*ssa.BasicBlock{}

			for i := pair.index1; i <= pair.index2; i++ {
				new_path = append(new_path, chain[i])
			}

			hash_new_path := hash_chain(new_path)
			for _,old_hash := range paths_hashs {
				if old_hash == hash_new_path {
					continue pair_loop
				}
			}

			paths = append(paths,new_path)
			paths_hashs = append(paths_hashs,hash_new_path)
		}

	}

	if len(paths) == 0 {
		err = errors.New("Can't find any path")
	}
	return
}

// Find_shortest_path_locally can return a slice of basic blocks from bb1 to bb2, which are in the same function. This slice is in order and minimum
func Find_shortest_path_locally(bb1,bb2 *ssa.BasicBlock) (shortest_path []*ssa.BasicBlock, err error) {
	shortest_path = []*ssa.BasicBlock{}
	err = nil
	if bb1.Parent().String() != bb2.Parent().String() {
		err = errors.New("Two bbs are not of the same function")
		return
	}
	if bb1 == bb2 {
		shortest_path = []*ssa.BasicBlock{bb1}
		return
	}

	all_paths,err := Find_paths_locally(bb1,bb2)
	if err != nil {
		return nil,err
	}

	min_length := 9999999999
	for _,path := range all_paths {
		if len(path) < min_length {
			min_length = len(path)
			shortest_path = path
		}
	}

	if len(shortest_path) == 0 {
		err = errors.New("Can't find a shortest path between:"+bb1.String()+"\t"+bb2.String())
	}

	return
}

func List_cond_of_path(path []*ssa.BasicBlock, target_bb *ssa.BasicBlock) (result []Cond) {
	for i,bb := range path {
		if bb == target_bb { // we need target_bb to determine when should we exist. Note that path's last bb may not be target_bb, due to Delete_useless_bb
			break
		}

		insts := bb.Instrs
		if len(insts) == 0 {
			continue
		}
		last_inst := insts[len(insts) - 1]
		inst_if,ok := last_inst.(*ssa.If)
		if !ok {
			continue
		}

		if i+1 >= len(path) {
			continue
		}

		next_bb := path[i+1]
		succs := bb.Succs
		if len(succs) != 2 { // This should never happen
			continue
		}

		var new_cond Cond
		new_cond.Cond = inst_if.Cond
		if next_bb == succs[0] {
			new_cond.Flag = true
			result = append(result,new_cond)
		} else if next_bb == succs[1] {
			new_cond.Flag = false
			result = append(result,new_cond)
		} else {
			// meaning there are bbs (that were deleted) between path[i] and path[i+1], do nothing
		}
	}
	return
}

// TODO: This function produces many redundant chains during calculation, meaning the algorithm is not efficient;
//  Every bb can only appear two times in this chain (see function recursive_append_succ)
func List_all_exe_chain(fn ssa.Function,end_bb *ssa.BasicBlock) (map[int][]*ssa.BasicBlock, error) {
	all_chains = make(map[int][]*ssa.BasicBlock)
	if len(fn.Blocks) > 0 {
		all_chains[0] = []*ssa.BasicBlock{fn.Blocks[0]}
		count_recursive = 0
		recursive_append_succ(0,end_bb)
	}

	if count_recursive > Max_recursive {
		err := errors.New("Reached max recursive number")
		return nil,err
	}


	//store chains whose last_bb has no successor to result
	result := make(map[int][]*ssa.BasicBlock)
	result_hashes := []string{}
	for index,_ := range all_chains {
		length := len(all_chains[index])
		if length > 0 {
			if last_bb := all_chains[index][length-1]; len(last_bb.Succs) == 0 || last_bb == end_bb {
				chain_hash := hash_chain(all_chains[index])
				if is_hash_in_result(chain_hash,result_hashes) == false {
					result[index] = all_chains[index]
					result_hashes = append(result_hashes,chain_hash)
				}
			}
		}
	}

	return result,nil
}

func is_hash_in_result(chain_hash string, result_hashes []string) bool {
	for _,hash := range result_hashes {
		if hash == chain_hash {
			return true
		}
	}
	return false
}

func hash_chain(chain []*ssa.BasicBlock) (result string) {
	for _,bb := range chain {
		result += strconv.Itoa(bb.Index)
		result += ";"
	}
	return
}

//recursive_append_succ aims for the chain indexed by chain_index. It will appends the next bb to the chain, and then call this function again.
// When there are N next bbs, it will create (N-1) new chains,
// and append one or two bb (if next bb has occurred, append two bbs to force it select a different path)
// to each of the N chains, and then call this function for each chain.
// Force return after Max_recursive times.
// If end_bb != nil, return when end_bb is reached
func recursive_append_succ(chain_index int,end_bb *ssa.BasicBlock) {

	count_recursive++
	if count_recursive > Max_recursive {
		return
	}

	chain := all_chains[chain_index]
	last_bb := chain[len(chain)-1]

	if last_bb == end_bb {
		return
	}

	//new_succs := []*ssa.BasicBlock{}
	//for _,suc := range last_bb.Succs {
	//	if is_bb_in_chain(suc,chain) == false {
	//		new_succs = append(new_succs,suc)
	//	}
	//}
	new_succs := last_bb.Succs

	if len(new_succs) == 0 {
		return
	} else {
		first_suc := new_succs[0]

		if index_bb := is_bb_in_chain(first_suc,chain); index_bb == -1 {
			all_chains[chain_index] = append(chain, first_suc)
			recursive_append_succ(chain_index,end_bb)
		} else { //meaning this suc has occurred before
			if index_bb == len(chain) - 1 {
				return
			} else {
				old_suc_of_suc := chain[index_bb + 1]
				var other_suc_of_suc *ssa.BasicBlock
				for _,suc_of_suc := range first_suc.Succs {
					if suc_of_suc != old_suc_of_suc {
						other_suc_of_suc = suc_of_suc
					}
				}
				if other_suc_of_suc == nil {
					return
				} else {
					all_chains[chain_index] = append(chain, first_suc, other_suc_of_suc)
					recursive_append_succ(chain_index,end_bb)
				}
			}
		}


		if len(new_succs) > 1 {
			for i,suc := range new_succs {
				if i == 0 {
					continue
				}

				last_index := len(all_chains)
				all_chains[last_index] = []*ssa.BasicBlock{}
				for _,old_bb := range chain {
					all_chains[last_index] = append(all_chains[last_index],old_bb)
				}

				if index_bb := is_bb_in_chain(suc,chain); index_bb == -1 {
					all_chains[last_index] = append(all_chains[last_index],suc)
					recursive_append_succ(last_index,end_bb)
				} else { //meaning this suc has occurred before
					if index_bb == len(chain) - 1 {
						return
					} else {
						old_suc_of_suc := chain[index_bb + 1]
						var other_suc_of_suc *ssa.BasicBlock
						for _,suc_of_suc := range suc.Succs {
							if suc_of_suc != old_suc_of_suc {
								other_suc_of_suc = suc_of_suc
							}
						}
						if other_suc_of_suc == nil {
							return
						} else {
							all_chains[last_index] = append(all_chains[last_index],suc)
							recursive_append_succ(last_index,end_bb)
						}
					}
				}
			}
		}
	}

}

func is_bb_in_chain(target_bb *ssa.BasicBlock, chain []*ssa.BasicBlock ) int {

	for i:=len(chain)-1; i>=0; i-- {
		if chain[i] == target_bb {
			return i
		}
	}

	return -1
}
