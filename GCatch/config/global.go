package config

import (
	"github.com/system-pclub/GCatch/GCatch/tools/go/callgraph"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"sync"
)


//-path=/home/song/work/go-workspace/code/src/github.com/etcd-io/etcd -include=github.com/etcd-io/etcd

var StrEntrancePath string // github.com/etcd-io/etcd
var StrGOPATH string // /home/song/work/go-workspace/code
var StrAbsolutePath string // /home/song/work/go-workspace/code/src/
var StrRelativePath string // github.com/etcd-io/etcd
var MapPrintMod map[string]bool // a map indicates which information will be printed

var BoolDisableFnPointer bool


var MapExcludePaths map[string]bool // a map indicates which package names should be ignored

var Prog *ssa.Program
var Pkgs []*ssa.Package


var BugIndex int
var BugIndexMu sync.Mutex

var VecPathStats [] PathStat

var CallGraph * callgraph.Graph
var Inst2Defers map[ssa.Instruction][]*ssa.Defer
var Defer2Insts map[*ssa.Defer][]ssa.Instruction
var Inst2CallSite map[ssa.CallInstruction]map[*callgraph.Edge]bool
