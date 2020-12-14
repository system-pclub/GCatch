package mycallgraph

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"strings"
)

type Call_node struct {
	Fn *ssa.Function
	Layer int
	Call_map map[ssa.Instruction][]*Call_node // we use []*Call_node here, because an inst
											// may call to different callees if interface or closure is involved
	Called_map map[ssa.Instruction]*Call_node
}

var Recursive_count_down int
var shown_nodes []*Call_node

func New_Call_node(fn *ssa.Function,layer int) *Call_node {
	return &Call_node{
		fn,
		layer,
		make(map[ssa.Instruction][]*Call_node),
		make(map[ssa.Instruction] *Call_node),
	}
}

func Initialize(Max_countdown int) {
	Recursive_count_down = Max_countdown
	shown_nodes = []*Call_node{}
}

// (layer_count_down + cn.layer) is a const number
func (cn *Call_node) Fill_call_map_after_init(max_layer int) () {
	Recursive_count_down --
	if cn.Layer > max_layer || Recursive_count_down <= 0 {
		return
	}
	shown_nodes = append(shown_nodes,cn)
	
	fn := cn.Fn

	node,ok := global.Call_graph.Nodes[fn]
	if !ok {
		return
	}
	for _, edge_out := range node.Out {

		// We have a whitelist global.C5_black_list_pkg. Skip any callee whose pkg is in this list
		if callee := edge_out.Callee; callee != nil {
			if callee_fn := edge_out.Callee.Func; callee_fn != nil {
				if callee_pkg := callee_fn.Pkg; callee_pkg != nil {
					if callee_pkg_pkg := callee_pkg.Pkg; callee_pkg_pkg != nil {
						flag_in_white_list := false
						for _,white_path := range global.C5_black_list_pkg {
							if callee_pkg_pkg.Path() == white_path {
								flag_in_white_list = true
							}
						}
						if flag_in_white_list == false && strings.Contains(callee_pkg_pkg.Path(),"/") == false {
							continue
						}
					}
				}
			}
		}

		var new_Call_node *Call_node

		flag_shown,shown_node := is_fn_in_node_slice(edge_out.Callee.Func,shown_nodes)
		if flag_shown {
			new_Call_node = shown_node
		} else {
			new_Call_node = New_Call_node(edge_out.Callee.Func,cn.Layer + 1)
			new_Call_node.Fill_call_map_after_init(max_layer)
		}

		old_nodes,ok := cn.Call_map[edge_out.Site]
		if !ok {
			old_nodes = []*Call_node{}
		}
		new_nodes := append(old_nodes,new_Call_node)
		cn.Call_map[edge_out.Site] = new_nodes
	}

	node,ok = global.Call_graph.Nodes[fn]
	if !ok {
		return
	}
	for _, edge_in := range node.In {

		// We have a whitelist global.C5_black_list_pkg. Skip any callee whose pkg is in this list
		if caller := edge_in.Caller; caller != nil {
			if caller_fn := caller.Func; caller_fn != nil {
				if caller_pkg := caller_fn.Pkg; caller_pkg != nil {
					if caller_pkg_pkg := caller_pkg.Pkg; caller_pkg_pkg != nil {
						flag_in_white_list := false
						for _,white_path := range global.C5_black_list_pkg {
							if caller_pkg_pkg.Path() == white_path {
								flag_in_white_list = true
							}
						}
						if flag_in_white_list == false && strings.Contains(caller_pkg_pkg.Path(),"/") == false {
							continue
						}
					}
				}
			}
		}

		flag_shown,shown_node := is_fn_in_node_slice(edge_in.Caller.Func,shown_nodes) 
		if flag_shown { //This node has shown up before, in progress or completed
			cn.Called_map[edge_in.Site] = shown_node
		} else {
			new_Call_node := New_Call_node(edge_in.Caller.Func,cn.Layer + 1)
			new_Call_node.Fill_call_map_after_init(max_layer)

			cn.Called_map[edge_in.Site] = new_Call_node
		}
	}
	
	return
}

func is_fn_in_node_slice(fn *ssa.Function, slice []*Call_node) (bool,*Call_node) {
	for _,old := range slice {
		if old.Fn == fn {
			return true,old
		}
	}
	return false,nil
}
