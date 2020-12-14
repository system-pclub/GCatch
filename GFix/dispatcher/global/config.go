package global

//This file contains some key parameters.
const Print_err_log = false

const Max_in_edges = 100

const Kill_FP_avoid_multi_report = false // This should be set true before used by a company: give only 1 bug report for a channel/cond/...

const Pointer_consider_reflection = false

const CHA_CallGraph_Prune_SDK = true
const CHA_CallGraph_Prune_Soundly = true

const C3_min_times_of_usage = 4
const C3_Ratio float32 = 0.75
const C3_kill_FP_1 bool = true
const C3_kill_FP_2 bool = true // see if there is an Unlock that will definitely be executed after Lock
const C3_FP_ratio float32 = 0.5
const C3_FP_layer int = 3
var C3_Exclude_slice = []string{ "init", "close", "start", "new", "lockfree","shutdown"}

const C5_call_chain_layer int = 8 // Be careful: Algorithm complexity: e^n, n = C5_call_chain_layer
const C5_call_chain_max_fn int = 100000
const C5_kill_FP_1 bool = true
const C5_kill_FP_2 bool = true
const C5_kill_FP_3 bool = false // This can reduce a lot of FP but will also produce FN
const C5_kill_FP_4 bool = false // If callee is in builtin packages, skip
var C5_black_list_pkg = []string{}

const C5A_kill_FP_1 bool = false

const C6_kill_FP_1 bool = false
const C6_kill_FP_2 bool = true
var C6_Exclude_slice = []string{ "init", "close", "shutdown","start","new"}

const C6A_kill_FP_1 bool = false // ignore channels and conds not as field
const C6A_call_chain_layer_for_inside_CS int = 5 //CS is critical section
const C6A_call_chain_layer_for_lock_wrapper int = 5
const C6A_max_count_for_lock_wrapper int = 10000
const C6A_call_chain_layer_for_chan_wrapper int = 3
const C6A_max_recursive_count = 5000
const C6A_wrapper_count = 50000000

const C7_report_directly_usage bool = true

const C7A_max_layer = 3
const C7A_max_call_chain_length = 3
const C7A_max_recursive_count = 50000
const C7A_kill_FP_ignoreTest = false

const C7B_max_call_path_length = 7
const C7B_max_layer = 7
const C7B_max_recursive_count = 5000
