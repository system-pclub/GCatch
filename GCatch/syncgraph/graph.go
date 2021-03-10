package syncgraph

// Define the types and their methods used to build a graph that contains all the information we need
// to generate Z3 constraints, including CFG, callgraph, alias information

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/analysis"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/path"
	"github.com/system-pclub/GCatch/GCatch/tools/go/callgraph"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"go/token"
	"strings"
)

type Node interface {
	Context() *CallCtx
	Instruction() ssa.Instruction
	Parent() *SyncGraph
	In() []*NodeEdge
	Out() []*NodeEdge
	CallCtx() *CallCtx
	InAdd(*NodeEdge)
	OutAdd(*NodeEdge)
	InOverWrite(*NodeEdge)
	OutOverWrite(*NodeEdge)
	SetId(int)
	GetId() int
	SetString(string)
	GetString() string
}

type node struct {
	Ctx    *CallCtx
	Instr  ssa.Instruction
	In_    []*NodeEdge
	Out_   []*NodeEdge
	String string

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

func (n *node) OutOverWrite(e *NodeEdge) {
	n.Out_ = []*NodeEdge{e}
}

func (n *node) InOverWrite(e *NodeEdge) {
	n.In_ = []*NodeEdge{e}
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

func (n *node) SetString(str string) {
	n.String = str
}

func (n *node) GetString() string {
	return n.String
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
	Inst   ssa.Instruction
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
	Op      instinfo.ChanOp
	Next    Node

	syncNode
}

func (a *ChanOp) Operation() interface{} {
	return a.Op
}

type LockerOp struct {
	Locker *instinfo.Locker
	Op     instinfo.LockerOp
	Next   Node

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
	Ctx  *CallCtx
}

type SyncGraph struct {
	// Prepare
	MainGoroutine    *Goroutine   // MainGoroutine is the Goroutine that is both a head and it contains the MakeChan operation
	HeadGoroutines   []*Goroutine // HeadGoroutines are the starting Goroutines that don't have creator in the graph
	Goroutines       []*Goroutine
	MapFirstNodeOfFn map[Node]struct{} // All Nodes that are the first Node of a function in SyncGraph
	Task             *Task

	// Build
	MapInstCtxKey2Node  map[InstCtxKey]Node // Two kinds of Node is not in this map: SelectCase, rundefer's Nodes
	Select2Case         map[*Select][]*SelectCase
	MapInstCtxKey2Defer map[InstCtxKey][]Node
	NodeStatus          map[Node]*Status
	MapPrim2VecSyncOp   map[interface{}][]SyncOp
	Visited             []*path.EdgeChain
	Worklist            []*Unfinish
	MapFnOnOpPath       map[*ssa.Function]struct{} // a map of functions that are on a path to reach a sync operation

	// Enumerate path
	PathCombinations []*pathCombination
	EnumerateCfg     *EnumeConfigure
}

type Status struct {
	Str     string // Str can be In_progress or Done. Only used to decide backedge for local nodes
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
	CallSite  Node
	Graph     *SyncGraph
}

type Unfinish struct {
	UnfinishedFn *ssa.Function
	Unfinished   Node
	IsGo         bool
	Site         *callgraph.Edge // Site.Callee has at least 1 bb, and this bb has at least 1 inst
	Dir          bool            // true if Call/Go (from caller to callee), false if Return (from callee to caller)
	Ctx          *CallCtx        // a new CallCtx used for the Site. This can be directly used
}

func NewGraph(task *Task) *SyncGraph {
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
		MapFnOnOpPath:       make(map[*ssa.Function]struct{}),
		MapFirstNodeOfFn:    make(map[Node]struct{}),
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
		fmt.Println("Fatal error in NewCtx: can't find the callgraph.Node for head function:", head.String())
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

// Update g.MapFnOnOpPath.
func (g *SyncGraph) ComputeFnOnOpPath() {
	for _, prim := range g.Task.VecTaskPrimitive {
		if prim == nil {
			continue
		}
		for _, opChain := range prim.Ops {
			if opChain == nil {
				continue
			}
			for _, chain := range opChain.Chains {
				if chain.Start != nil && chain.Start.Func != nil {
					g.MapFnOnOpPath[chain.Start.Func] = struct{}{}
				}
				for _, edge := range chain.Chain {
					if edge == nil {
						continue
					}
					if edge.Caller == nil || edge.Caller.Func == nil {
						continue
					}
					g.MapFnOnOpPath[edge.Caller.Func] = struct{}{}
					if edge.Callee == nil || edge.Callee.Func == nil {
						continue
					}
					g.MapFnOnOpPath[edge.Callee.Func] = struct{}{}
				}
			}
		}
	}
}

// If BB X post-dominates BB Y, and all BBs on any path from Y to X don't contain important Nodes, and X and Y are in the same layer of loop,
// we can ignore these BBs between X and Y.
// We will link Y and X directly (let the first Node of X be the next Node of the last Node of Y)
// What are important Nodes: Call to functions on MapFnOnOpPath/ operation of dependent primitives/ Return/ Kill/ End
func (g *SyncGraph) OptimizeBB_V1() {
	for headNode, _ := range g.MapFirstNodeOfFn {
		fn := headNode.Instruction().Parent()
		postDom := analysis.NewPostDominator(fn)
		loopAnalysis := analysis.NewLoopAnalysis(fn)

		if strings.Contains(fn.String(), "http2Client") && strings.Contains(fn.String(), "reader") {
			fmt.Print()
		}

		// Enumerate X and Y
		for _, bbY := range fn.Blocks {
			if len(bbY.Instrs) == 0 {
				continue
			}
			// There may be multiple X that satisfy our conditions
			vecBBX := []*ssa.BasicBlock{}
			for _, bbX := range fn.Blocks {
				if bbX != bbY && postDom.Dominate(bbX, bbY) {
					if len(bbX.Instrs) == 0 {
						continue
					}
					// List paths from Y to X
					vecPath := path.EnumeratePathForPostDomBBs(bbY, bbX)
					// Check if all paths don't contain important Nodes
					boolAllPathDoNotContain := true
				pathLoop:
					for _, aPath := range vecPath {
						for _, bb := range aPath {
							if bb == bbX || bb == bbY { // only check bbs between X and Y
								continue
							}
							for _, inst := range bb.Instrs {
								if g.isInstImportant(inst) {
									boolAllPathDoNotContain = false
									break pathLoop
								}
							}
						}
					} // end of pathLoop

					if boolAllPathDoNotContain {
						// Check if BBX and BBY are in the same layer of loop
						if bbY.Index == 12 {
							fmt.Print()
						}
						headersX := loopAnalysis.MapBodyBb2LoopHead[bbX]
						headersY := loopAnalysis.MapBodyBb2LoopHead[bbY]
						if isBBSliceEqual(headersX, headersY) {
							vecBBX = append(vecBBX, bbX)
						} else {
							// It is also allowed if X is the loop header of Y
							if isBBSliceEqual(append(headersX, bbX), headersY) {
								vecBBX = append(vecBBX, bbX)
							}
						}
					}
				}

			} // end of bbX loop

			if len(vecBBX) == 0 {
				continue
			}

			// find a bbX that is far from bbY. This can be not so accurate
			var bbX *ssa.BasicBlock
			largestDistance := -9999
			for _, bb := range vecBBX {
				distance := bb.Index - bbY.Index
				if distance > largestDistance {
					bbX = bb
					largestDistance = distance
				}
			}

			if bbX != nil {

				// Link X and Y
				var lastInstY, firstInstX ssa.Instruction
				lastInstY = bbY.Instrs[len(bbY.Instrs)-1]
				firstInstX = bbX.Instrs[0]
				var lastNodeY, firstNodeX Node
				for _, node := range g.MapInstCtxKey2Node {
					inst := node.Instruction()
					if inst == lastInstY {
						lastNodeY = node
					} else if inst == firstInstX {
						firstNodeX = node
					}
				}
				if lastNodeY == nil || firstNodeX == nil {
					continue
				}
				newNodeEdge := &NodeEdge{
					Prev:       lastNodeY,
					Succ:       firstNodeX,
					IsBackedge: false,
					IsCall:     false,
					IsGo:       false,
					AddValue:   0,
				}
				lastNodeY.OutOverWrite(newNodeEdge)
				//firstNodeX.InOverWrite(newNodeEdge)
			}
		}
	}

}

func (g *SyncGraph) OptimizeBB_V2() {
	for headNode, _ := range g.MapFirstNodeOfFn {
		fn := headNode.Instruction().Parent()

		// Enumerate X and Y
		for _, bbY := range fn.Blocks {
			if len(bbY.Instrs) == 0 {
				continue
			}

			// See if there is a bbX such that: bbY -> bbX and bbZ, bbZ -> bbX
			if len(bbY.Succs) != 2 {
				continue
			}
			var bbX *ssa.BasicBlock
			for i, suc := range bbY.Succs {
				if len(suc.Preds) != 2 {
					continue
				}
				var otherBB *ssa.BasicBlock
				if i == 0 {
					otherBB = bbY.Succs[1]
				} else {
					otherBB = bbY.Succs[0]
				}
				if len(otherBB.Succs) != 1 || otherBB.Succs[0] != suc {
					continue
				}
				boolPredFoundY, boolPredFoundZ := false, false
				for _, pred := range suc.Preds {
					if pred == bbY {
						boolPredFoundY = true
					}
					if pred == otherBB {
						boolPredFoundZ = true
					}
				}
				if boolPredFoundZ && boolPredFoundY {
					bbX = suc
					break
				}
			}

			if bbX != nil {
				// Link X and Y
				var lastInstY, firstInstX ssa.Instruction
				lastInstY = bbY.Instrs[len(bbY.Instrs)-1]
				firstInstX = bbX.Instrs[0]
				var lastNodeY, firstNodeX Node
				for _, node := range g.MapInstCtxKey2Node {
					inst := node.Instruction()
					if inst == lastInstY {
						lastNodeY = node
					} else if inst == firstInstX {
						firstNodeX = node
					}
				}
				if lastNodeY == nil || firstNodeX == nil {
					continue
				}
				newNodeEdge := &NodeEdge{
					Prev:       lastNodeY,
					Succ:       firstNodeX,
					IsBackedge: false,
					IsCall:     false,
					IsGo:       false,
					AddValue:   0,
				}
				lastNodeY.OutOverWrite(newNodeEdge)
				firstNodeX.InOverWrite(newNodeEdge)
			}
		}
	}

}

func (g *SyncGraph) isInstImportant(inst ssa.Instruction) bool {
	switch concrete := inst.(type) {
	case *ssa.MakeChan, *ssa.Return, *ssa.Go, *ssa.Select, *ssa.Send, *ssa.Panic:
		return true
	case *ssa.Call:
		if g.isInstCallImportantFn(concrete) {
			return true
		}
		if isKillThread(concrete) {
			return true
		}
		//if instinfo.IsMutexLock(concrete) || instinfo.IsMutexUnlock(concrete) || instinfo.IsRwmutexLock(concrete) ||
		//	instinfo.IsRwmutexUnlock(concrete) || instinfo.IsRwmutexRlock(concrete) || instinfo.IsRwmutexRunlock(concrete) {
		//	return true
		//}
	case *ssa.RunDefers:
		if vecDefer, ok := config.Inst2Defers[concrete]; ok {
			for _, aDefer := range vecDefer {
				if g.isInstCallImportantFn(aDefer) {
					return true
				}
			}
		}
	case *ssa.UnOp:
		if concrete.Op == token.ARROW {
			return true
		}
	}

	return false
}

func (g *SyncGraph) isInstCallImportantFn(inst ssa.CallInstruction) bool {
	if mapEdge, ok := config.Inst2CallSite[inst]; ok {
		for edge, _ := range mapEdge {
			if edge.Callee.Func == nil {
				continue
			}
			if _, ok := g.MapFnOnOpPath[edge.Callee.Func]; ok {
				return true
			}
		}
	}
	return false
}
