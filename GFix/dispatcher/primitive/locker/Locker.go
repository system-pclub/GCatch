package locker

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/sync_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/output"
	"github.com/system-pclub/GCatch/GFix/dispatcher/search"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/callgraph"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/types"
	"strings"
)

type Locker struct{
	Name string
	Type string
	Locks []*Lock_op
	Unlocks []*Unlock_op
	Pkg string // Don't use *ssa.Package here! It's not reliable

	Status string
}

type Lock_op struct {
	Name string
	Inst ssa.Instruction
	Is_RLock bool
	Is_defer bool

	Wrappers []*Wrapper //This field is specially designed for C6A

	Parent *Locker
}

type Unlock_op struct {
	Name string
	Inst ssa.Instruction
	Is_RUnlock bool
	Is_defer bool

	Parent *Locker
}

// Wrapper records a function that "contains" a Lock_op. "Contains" means the function directly uses this Lock_op, or
// its callee (or callee's callee) uses this Lock_op. The maximum layer is C6A_call_chain_layer_for_lock_wrapper
type Wrapper struct {
	Fn     *ssa.Function // When compare two Wrapper, can't directly compare Fn or inst, because pointer will change during each compilation
	Fn_str	string
	Inst   ssa.Instruction
	Callee *Wrapper // if callee is nil, then inst is the Lock_op itself, else inst is calling to another Wrapper
	Op     *Lock_op // the wrapped operation
}

const Unknown = "Unknown"
const Edited = "Edited"
const M_Lock = "M_Lock"
const M_Unlock = "M_Unlock"
const RW_Lock = "RW_Lock"
const RW_RLock = "RW_RLock"
const RW_Unlock = "RW_Unlock"
const RW_RUnlock = "RW_RUnlock"

func (l *Locker) Modify_status(str string) {
	l.Status = str
}

func Scan_lock_inst_return_value_comment(inst ssa.Instruction) (v ssa.Value, comment string) {
	v = nil
	comment = ""
	if inst.Parent().Pkg == nil {
		return
	}

	inst_call,ok := inst.(ssa.CallInstruction)
	if !ok {
		return
	}

	args := inst_call.Common().Args
	switch {
	case sync_check.Is_mutex_lock(inst):
		if inst_call.Common().IsInvoke() == false {
			//if args[0].Type().String() != "*sync.Mutex" {
			//	fmt.Println("Warning: a Mutex Lock op's argument has type:",args[0].Type().String())
			//}
			if len(args) == 0 {
				output.Print_inst_and_location(inst)
				fmt.Println("Warning: a  Mutex Lock op (static call) has no arguments")
				return
			}
			return args[0],M_Lock
		} else {
			//if inst_call.Common().Value.Type().String() !="sync.Locker" {
			//	output.Print_inst_and_location(inst)
			//	fmt.Println("Warning: a Mutex Lock op's argument has type:",inst_call.Common().Value.Type().String())
			//}
			return inst_call.Common().Value,M_Lock
		}
	case sync_check.Is_mutex_unlock(inst):
		if inst_call.Common().IsInvoke() == false {
			if len(args) == 0 {
				output.Print_inst_and_location(inst)
				fmt.Println("Warning: a  Mutex Unlock op (static call) has no arguments")
				return
			}
			return args[0],M_Unlock
		} else {
			return inst_call.Common().Value,M_Unlock
		}
	case sync_check.Is_rwmutex_lock(inst):
		if inst_call.Common().IsInvoke() == false {
			if len(args) == 0 {
				output.Print_inst_and_location(inst)
				fmt.Println("Warning: a  RWMutex Lock op (static call) has no arguments")
				return
			}
			return args[0],RW_Lock
		} else {
			return inst_call.Common().Value,RW_Lock
		}
	case sync_check.Is_rwmutex_unlock(inst):
		if inst_call.Common().IsInvoke() == false {
			if len(args) == 0 {
				output.Print_inst_and_location(inst)
				fmt.Println("Warning: a  RWMutex Unlock op (static call) has no arguments")
				return
			}
			return args[0],RW_Unlock
		} else {
			return inst_call.Common().Value,RW_Unlock
		}
	case sync_check.Is_rwmutex_rlock(inst):
		if inst_call.Common().IsInvoke() == false {
			if len(args) == 0 {
				output.Print_inst_and_location(inst)
				fmt.Println("Warning: a  RWMutex RLock op (static call) has no arguments")
				return
			}
			return args[0],RW_RLock
		} else {
			return inst_call.Common().Value,RW_RLock
		}
	case sync_check.Is_rwmutex_runlock(inst):
		if inst_call.Common().IsInvoke() == false {
			if len(args) == 0 {
				output.Print_inst_and_location(inst)
				fmt.Println("Warning: a  RWMutex RUnlock op (static call) has no arguments")
				return
			}
			return args[0],RW_RUnlock
		} else {
			return inst_call.Common().Value,RW_RUnlock
		}
	default:
		return
	}
}

func Scan_inst_and_record_to_slice(lockers []*Locker, inst ssa.Instruction) (result []*Locker, is_locker_op bool) {
	result = lockers
	is_locker_op = false

	if inst.Parent().Pkg == nil {
		return
	}

	name,lock_type,op_type,is_defer,is_locker := Find_locker_info_inst(inst)

	if is_locker == false {
		return
	}
	is_locker_op = true
	if name == Unknown || op_type == Unknown {
		return
	}

	// Now we know we find an inst calling a Locker

	var edit_lock *Locker
	for _,lock := range result {
		if lock.Name == name && lock.Pkg == inst.Parent().Pkg.Pkg.Path() {
			edit_lock = lock // we find an existing *Lock whose name matches
			edit_lock.Modify_status(Edited)
			break
		}
	}
	if edit_lock == nil { // no existing *Lock matches. Create a new one
		edit_lock = &Locker{
			Name:    name,
			Type:    lock_type,
			Locks:   []*Lock_op{},
			Unlocks: []*Unlock_op{},
			Pkg:     inst.Parent().Pkg.Pkg.Path(),
			Status:  Edited,
		}
		result = append(result,edit_lock)
	}

	is_RLock_or_RUnlock := false
	if op_type == "RLock" || op_type == "RUnlock" {
		is_RLock_or_RUnlock = true
	}

	if op_type == "Lock" || op_type == "RLock" {
		new_lock_op := &Lock_op{
			Name:     name,
			Inst:     inst,
			Is_RLock: is_RLock_or_RUnlock,
			Is_defer: is_defer,
			Parent: edit_lock,
		}
		edit_lock.Locks = append(edit_lock.Locks,new_lock_op)
	} else if op_type == "Unlock" || op_type == "RUnlock" {
		new_unlock_op := &Unlock_op{
			Name:     name,
			Inst:     inst,
			Is_RUnlock: is_RLock_or_RUnlock,
			Is_defer: is_defer,
			Parent: edit_lock,
		}
		edit_lock.Unlocks = append(edit_lock.Unlocks,new_unlock_op)
	}

	return
}

func Find_locker_info_inst(inst ssa.Instruction) (name, locker_type, op_type string, is_defer, is_locker bool) {
	name = Unknown
	locker_type = Unknown
	op_type = Unknown
	is_defer = false
	is_locker = true

	var call *ssa.CallCommon

	switch inst_type := inst.(type) {
	case *ssa.Call:
		call = inst_type.Common()
	case *ssa.Defer:
		call = inst_type.Common()
		is_defer = true
	default :
		is_locker = false
		return
	}

	call_name := CallName(call)
	switch call_name {
	case "(*sync.Mutex).Lock":
		locker_type = "Mutex"
		op_type = "Lock"

	case "(*sync.Mutex).Unlock":
		locker_type = "Mutex"
		op_type = "Unlock"

	case "(*sync.RWMutex).Lock":
		locker_type = "RWMutex"
		op_type = "Lock"

	case "(*sync.RWMutex).Unlock":
		locker_type = "RWMutex"
		op_type = "Unlock"

	case "(*sync.RWMutex).RLock":
		locker_type = "RWMutex"
		op_type = "RLock"

	case "(*sync.RWMutex).RUnlock":
		locker_type = "RWMutex"
		op_type = "RUnlock"

	default:
		var callee_name string = ""
		if call.IsInvoke() {
			callee_name = call.Method.Name()
		} else {
			if call_static_fn,ok := call.Value.(*ssa.Function);ok {
				callee_name = call_static_fn.Name()
			}
		}
		if callee_name != "" {
			switch {
			case case_insensitive_equal(callee_name,"Lock"): op_type = "Lock"
			case case_insensitive_equal(callee_name,"Unlock"): op_type = "Unlock"
			case case_insensitive_equal(callee_name,"RLock"): op_type = "RLock"
			case case_insensitive_equal(callee_name,"RUnlock"): op_type = "RUnlock"
			default: is_locker = false; return
			}
		} else {
			is_locker = false; return
		}

	}

	name = search.Mutex_name(inst)
	if name == "" {
		name = Unknown
	}

	//if locker_type == Unknown { // This is not a standard Locker. We use what is before ".Lock()" in the line of code as the name
	//
	//} else { // This is a standard Locker
	//	if len(call.Args) != 1 { // This should never happen for standard Locker
	//		name = Unknown
	//		return
	//	}
	//
	//	locker_v := call.Args[0]
	//
	//	switch locker := locker_v.(type) {
	//	case *ssa.FieldAddr:
	//		// For locker in field: "FIELD_${index}_" + ("PTR_") + "${type}" is the name
	//		name = "FIELD_"
	//		name += strconv.Itoa(locker.Field)
	//
	//		var struct_ *ssa.Alloc
	//		if struct_in_UnOp,ok := locker.X.(*ssa.UnOp);ok {
	//			if struct_Alloc,ok := struct_in_UnOp.X.(*ssa.Alloc);ok && struct_in_UnOp.Op == token.MUL { // locker is a field of *XXX
	//				name += "PTR_"
	//				struct_ = struct_Alloc
	//			}
	//		} else if struct_Alloc_in_FieldAddr,ok := locker.X.(*ssa.Alloc);ok {
	//			struct_ = struct_Alloc_in_FieldAddr
	//		}
	//
	//		if struct_ == nil { // This happens so frequently. Like s.mu.Lock() in an anonymous function
	//			fmt.Println("Fail in Find_locker_info_inst: can't obtain name from FieldAddr in Inst:")
	//			output.Print_inst_and_location(inst)
	//			name = Unknown
	//			return
	//		}
	//
	//		name += struct_.Type().String()
	//
	//	case *ssa.Alloc:
	//		// For global locker: "GLOBAL_${variable_name}" is the name
	//		// For local locker: "LOCAL_${fn.Name()}_${variable_name}" is the name
	//
	//	default:
	//		name = Unknown
	//	}
	//}

	return
}

func Loop_lockers_to_record_wrappers(all_lockers []*Locker) (result []*Locker) {
	result = all_lockers

	for _,locker := range result {
		for _,lock := range locker.Locks {
			parent_wrapper := &Wrapper{
				Fn:     lock.Inst.Parent(),
				Fn_str: lock.Inst.Parent().String(),
				Inst:   lock.Inst,
				Callee: nil,
				Op:     lock,
			}
			init_wrappers := []*Wrapper{parent_wrapper}
			wrappers := recursive_list_wrappers(parent_wrapper,0,init_wrappers)
			lock.Wrappers = wrappers
		}
	}

	return
}

func recursive_list_wrappers(this_wrapper *Wrapper, layer int, old_wrappers []*Wrapper) (result []*Wrapper) {
	result = old_wrappers

	if layer > global.C6A_call_chain_layer_for_lock_wrapper || len(result) > global.C6A_max_count_for_lock_wrapper{
		return
	}

	in_edges := []*callgraph.Edge{}
	this_node,ok := global.Call_graph.Nodes[this_wrapper.Fn]
	if !ok {
		return
	}
	in_edges = this_node.In
	for _,caller_edge := range in_edges {
		if _,is_go := caller_edge.Site.(*ssa.Go);is_go { // meaning the caller is in another goroutine, can't record this caller
			continue
		}
		new_wrapper := &Wrapper{
			Fn:     caller_edge.Caller.Func,
			Fn_str:     caller_edge.Caller.Func.String(),
			Inst:   caller_edge.Site,
			Callee: this_wrapper,
			Op:     this_wrapper.Op,
		}
		if is_wrapper_in_slice(new_wrapper, result) == false {
			result = append(result,new_wrapper)
			result = recursive_list_wrappers(new_wrapper,layer + 1,result)
		}
	}

	return
}

func CallName(call *ssa.CallCommon) string {

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

func case_insensitive_equal(s1, s2 string) bool {
	s1, s2 = strings.ToUpper(s1), strings.ToUpper(s2)
	return s1 == s2
}

func Append_wrappers_if_not_in(old[]*Wrapper, add *Wrapper) (result []*Wrapper) {
	result = old
	if is_wrapper_in_slice(add,result) == false {
		result = append(result,add)
	}
	return
}

func is_wrapper_in_slice(w *Wrapper,s []*Wrapper) bool {
	for _,old := range s {
		if w == old {
			return true
		}
	}
	return false
}
