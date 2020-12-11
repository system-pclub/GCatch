package syncgraph

// Define the types and their methods used to build a graph that contains all the information we need
// to generate Z3 constraints, including CFG, callgraph, alias information

import (
	"fmt"
	"github.com/system-pclub/GCatch/config"
	"github.com/system-pclub/GCatch/instinfo"
	"github.com/system-pclub/GCatch/path"
	"github.com/system-pclub/GCatch/tools/go/callgraph"
	"github.com/system-pclub/GCatch/tools/go/ssa"
)

type Node interface{
	Context() *CallCtx
	Instruction() ssa.Instruction
	Parent() *SyncGraph
	In() []*NodeEdge
	Out() []*NodeEdge
	CallCtx() *CallCtx
	InAdd(*NodeEdge)
	OutAdd(*NodeEdge)
	SetId(int)
	GetId() int
}

type node struct {
	Ctx *CallCtx
	Instr ssa.Instruction
	In_ []*NodeEdge
	Out_ []*NodeEdge

	ID int
}

func (n *node) Instruction() ssa.Instruction {
	return n.Instr
}

func (n *node) Context() *CallCtx {
	return n.Ctx
}

func (n *node) Parent() *SyncGraph {
	return n.Ctx.Graph
}

func (n *node) In() []*NodeEdge {
	return n.In_
}

func (n *node) Out() []*NodeEdge {
	return n.Out_
}

func (n *node) CallCtx() *CallCtx {
	return n.Ctx
}

func (n *node) OutAdd(e *NodeEdge) {
	n.Out_ = append(n.Out_, e)
}

func (n *node) InAdd(e *NodeEdge) {
	n.In_ = append(n.In_, e)
}

func (n *node) SetId(id int) {
	n.ID = id
}

func (n *node) GetId() int {
	return n.ID
}

var intFakeNodeId int
// Returns a unique Node that doesn't have any information
func Fake_Node() Node {
	intFakeNodeId++
	return &node{
		Ctx:   nil,
		Instr: nil,
		In_:   nil,
		Out_:  nil,
		ID:    intFakeNodeId,
	}
}

type Jump struct {
	Inst            *ssa.Jump
	Next            Node
	BoolIsNextexist bool
	BoolIsBackedge  bool

	node
}

type If struct {
	Inst               *ssa.If
	Cond               ssa.Value
	Then               Node
	Else               Node
	BoolIsThenBackedge bool
	BoolIsElseBackedge bool

	node
}

type Call struct {
	Inst      ssa.CallInstruction
	Calling   map[*callgraph.Edge]Node
	NextLocal Node

	node
}

const NotInteresting = 0
const MaxRecursive = 1
const EndDefer = 2

type End struct {
	Inst ssa.Instruction
	Reason int

	node
}

type Overwrite struct {
	Inst *ssa.MakeChan
	Chan *instinfo.Channel

	node
}

type Return struct {
	Inst                 ssa.Instruction
	BoolIsEndOfGoroutine bool
	Caller               Node

	node
}

type Kill struct {
	Inst        ssa.Instruction // can be *ssa.Panic or *ssa.Call (callee is t.Fatal/Fatalf/...)
	BoolIsPanic bool
	BoolIsFatal bool
	Next        Node

	node
}

type Go struct {
	Inst                *ssa.Go
	MapCreateGoroutines map[*callgraph.Edge]*Goroutine
	MapCreateNodes      map[*callgraph.Edge]Node
	NextLocal           Node

	node
}

type ChanMake struct {
	Inst    *ssa.MakeChan
	Channel *instinfo.Channel
	MakeOp  instinfo.ChanOp
	Next    Node

	syncNode
}

func (c *ChanMake) Operation() interface{} {
	return c.MakeOp
}

type Select struct {
	Inst           *ssa.Select
	Cases          map[int]*SelectCase
	BoolHasDefault bool
	DefaultCase    *SelectCase

	node
}

type SyncOp interface {
	Primitive() interface{} // *channel.Channel or *locker.Locker or *cond.Cond or *waitgroup.WG
	MapAliasOp() map[SyncOp]bool
	MapSyncOp() map[SyncOp]bool
	Operation() interface{}
}

type syncNode struct {
	Prim                  interface{}
	BoolIsAllAliasInGraph bool
	AliasOp               map[SyncOp]bool
	SyncOp                map[SyncOp]bool
	node
}

func (a *syncNode) MapAliasOp() map[SyncOp]bool {
	return a.AliasOp
}
func (a *syncNode) MapSyncOp() map[SyncOp]bool {
	return a.SyncOp
}
func (a *syncNode) Primitive() interface{} {
	return a.Prim
}

type SelectCase struct {
	Channel        *instinfo.Channel
	Op             instinfo.ChanOp
	Index          int // -1 if this is default
	BoolIsDefault  bool
	Next           Node
	BoolIsBackedge bool
	Select         *Select

	syncNode
}

func (a *SelectCase) Operation() interface{} {
	return a.Op
}

// Can be send/receive/close. Note that send and receive here must not be in select
type ChanOp struct {
	Channel *instinfo.Channel
	Op instinfo.ChanOp
	Next Node

	syncNode
}

func (a *ChanOp) Operation() interface{} {
	return a.Op
}

type LockerOp struct {
	Locker *instinfo.Locker
	Op instinfo.LockerOp
	Next Node

	syncNode
}

func (a *LockerOp) Operation() interface{} {
	return a.Op
}


type NormalInst struct {
	Inst ssa.Instruction
	Next Node

	node
}

type InstCtxKey struct { // Key that considers both ssa.Instruction and Ctx
	Inst ssa.Instruction
	Ctx *CallCtx
}

type SyncGraph struct {
	// Prepare
	MainGoroutine  *Goroutine   // MainGoroutine is the Goroutine that is both a head and it contains the MakeChan operation
	HeadGoroutines []*Goroutine // HeadGoroutines are the starting Goroutines that don't have creator in the graph
	Goroutines     []*Goroutine
	Task           *Task

	// Build
	MapInstCtxKey2Node  map[InstCtxKey]Node // Two kinds of Node is not in this map: SelectCase, rundefer's Nodes
	Select2Case         map[*Select][]*SelectCase
	MapInstCtxKey2Defer map[InstCtxKey][]Node
	NodeStatus          map[Node]*Status
	MapPrim2VecSyncOp   map[interface{}][]SyncOp
	Visited             []*path.EdgeChain
	Worklist            []*Unfinish

	// Enumerate path
	PathCombinations []*pathCombination
	EnumerateCfg     *EnumeConfigure
}

type Status struct {
	Str string // Str can be In_progress or Done. Only used to decide backedge for local nodes
	Visited int
}

const In_progress = "In_progress"
const Done = "Done"

type NodeEdge struct {
	Prev       Node
	Succ       Node
	IsBackedge bool
	IsCall     bool
	IsGo       bool
	AddValue   int
}

type Goroutine struct {
	Creator  *Go
	EntryFn  *ssa.Function
	IsMain   bool // If IsMain is true then Creator == nil
	HeadNode Node

	Graph *SyncGraph
}

type CallCtx struct {
	CallChain *path.EdgeChain
	Goroutine *Goroutine
	CallSite Node
	Graph *SyncGraph
}

type Unfinish struct {
	UnfinishedFn *ssa.Function
	Unfinished   Node
	IsGo         bool
	Site         *callgraph.Edge // Site.Callee has at least 1 bb, and this bb has at least 1 inst
	Dir          bool // true if Call/Go (from caller to callee), false if Return (from callee to caller)
	Ctx          *CallCtx // a new CallCtx used for the Site. This can be directly used
}


func NewGraph(task *Task) *SyncGraph{
	newGraph := &SyncGraph{
		MainGoroutine:       nil,
		HeadGoroutines:      []*Goroutine{},
		Goroutines:          []*Goroutine{},
		Task:                task,
		MapInstCtxKey2Node:  make(map[InstCtxKey]Node),
		Select2Case:         make(map[*Select][]*SelectCase),
		NodeStatus:          make(map[Node]*Status),
		MapInstCtxKey2Defer: make(map[InstCtxKey][]Node),
		MapPrim2VecSyncOp:   make(map[interface{}][]SyncOp),
		Worklist:            nil,
		Visited:             []*path.EdgeChain{},
		PathCombinations:    nil,
		EnumerateCfg:        nil,
	}

	return newGraph
}

func (g *SyncGraph) NewGoroutine(headFn *ssa.Function) *Goroutine {
	newGoroutine := &Goroutine{
		Creator:  nil,
		EntryFn:  headFn,
		IsMain:   true,
		HeadNode: nil,
		Graph:    g,
	}
	g.Goroutines = append(g.Goroutines, newGoroutine)
	return newGoroutine
}

func (g *SyncGraph) NewCtx(goroutine *Goroutine, head *ssa.Function) *CallCtx {
	headNode := config.CallGraph.Nodes[head]
	if headNode == nil {
		fmt.Println("Fatal error in NewCtx: can't find the callgraph.Node for head function:",head.String())
	}
	newEdgePath := &path.EdgeChain{
		Chain: nil,
		Start: headNode,
	}
	newCtx := &CallCtx{
		CallChain: newEdgePath,
		Goroutine: goroutine,
		CallSite:  nil,
		Graph:     g,
	}
	return newCtx
}
