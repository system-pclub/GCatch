package global

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/callgraph"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/mypointer"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"os"
	"sync"
)

var Root string
var Entrance, Include string
var Exclude []string
var Absolute_root string
var Prog *ssa.Program
var Pkgs []*ssa.Package
var Call_graph *callgraph.Graph
var CHA_Call_graph *callgraph.Graph
var Static_call_graph *callgraph.Graph
var Inst2edge map[ssa.CallInstruction][]*callgraph.Edge
var Bug_index int
var Bug_index_mu sync.Mutex
var Num_scan_pkg_lock_send int
var Last_progress float32
var Worthy_path []Parent_path
var Process_count int
var GOPATH string
var Target_GOPATH string
var PointerAnalysisResult *mypointer.Result

var Wg_B1 sync.WaitGroup

var All_pkg_paths []Path_stat
var All_method []*ssa.Function
var All_struct []*Struct
var All_sync_struct []*Sync_struct
var All_sync_global []*ssa.Global
var Defer_map map[*ssa.RunDefers][]*ssa.Defer

var Output_file *os.File
var Flag_write_inst = false
