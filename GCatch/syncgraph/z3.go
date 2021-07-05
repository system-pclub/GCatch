package syncgraph

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/tools/github.com/aclements/go-z3/z3"
)

type ZGoroutine struct {
	PtrPPath     *PPath
	Nodes        []ZNode // Not all nodes in PPath will be reserved
	IsTerminated bool
	BlockAt      int
	ID           int

	Z3Sys *Z3System
}

type ZNode interface {
	ID() int
	SetId(int)
	Goroutine() *ZGoroutine
	PNode() *PNode
	IsBlocking() bool
	TraceOrder() z3.Int
	UpdateOrder(z3.Int)
}

type ZNodeBasic struct {
	Id             int
	ZGoroutine     *ZGoroutine
	PtrPNode       *PNode
	boolIsBlocking bool // If boolIsBlocking, then this Node is not executed, else executed
	Order          z3.Int
}

func (b *ZNodeBasic) ID() int {
	return b.Id
}

func (b *ZNodeBasic) SetId(id int) {
	b.Id = id
}

func (b *ZNodeBasic) Goroutine() *ZGoroutine {
	return b.ZGoroutine
}

func (b *ZNodeBasic) PNode() *PNode {
	return b.PtrPNode
}

func (b *ZNodeBasic) IsBlocking() bool {
	return b.boolIsBlocking
}

func (b *ZNodeBasic) TraceOrder() z3.Int {
	return b.Order
}

func (b *ZNodeBasic) UpdateOrder(i z3.Int) {
	b.Order = i
}

type ZNodeNbSend struct {
	ZNodeBasic
	Closes []*ZNodeClose
	Pairs []*ZSendRecvPair
}

type ZNodeNbRecv struct {
	ZNodeBasic
	Pairs     []*ZSendRecvPair
	Closes    []*ZNodeClose
	FromClose z3.Bool
}

type ZNodeBSend struct {
	Buffer   z3.Int
	Other_SR []ZNode
	ZNodeBasic
}

type ZNodeBRecv struct {
	Buffer     z3.Int
	Other_SR   []ZNode
	Closes     []*ZNodeClose
	From_close z3.Bool
	ZNodeBasic
}

type ZNodeClose struct {
	ZNodeBasic
}

type ZSendRecvPair struct {
	boolUseThisPair z3.Bool
	Send            *ZNodeNbSend
	Recv            *ZNodeNbRecv
}

type Z3System struct {
	vecZGoroutines         []*ZGoroutine
	mapOriginBlockingNodes map[ZNode]struct{} // ZNode that are the first blocking node in its thread
	vecZSendRecvPairs      []*ZSendRecvPair

	Constraints     *AllConstraints
	countUniqueName int

	Warnings []string
	Config   *Z3Cfg
	Z3Ctx    *z3.Context
	Solver   *z3.Solver
}

type AllConstraints struct {
	Order         []z3.Bool
	SyncOfOp      []z3.Bool
	PairInferRule []z3.Bool
	Blocking      []z3.Bool
}

const (
	WARNING_met_chan_var_buf = "WARNING_met_chan_var_buf"
	WARNING_met_nil_chan     = "WARNING_met_nil_chan"
)

var (
	Z3One  z3.Int
	Z3Zero z3.Int
)

type Z3Cfg struct {
	Mode int
}

const (
	ZMode_GL = 0 // In this mode, only nodes of type *Go and SyncOp will be recorded
)

func NewZ3ForGl() *Z3System {
	newZ3Sys := &Z3System{
		vecZGoroutines:         nil,
		Config:                 &Z3Cfg{Mode: ZMode_GL},
		mapOriginBlockingNodes: make(map[ZNode]struct{}),
		Constraints:            &AllConstraints{},
	}
	z3ContextCfg := z3.NewContextConfig()
	newZ3Sys.Z3Ctx = z3.NewContext(z3ContextCfg)

	// some popular z3 const
	Z3Zero = newZ3Sys.Z3Ctx.FromInt(0, newZ3Sys.Z3Ctx.IntSort()).(z3.Int)
	Z3One = newZ3Sys.Z3Ctx.FromInt(1, newZ3Sys.Z3Ctx.IntSort()).(z3.Int)

	return newZ3Sys
}

func (z *Z3System) Z3Main(p_paths []*PPath, block_poses []blockingPos, boolCheckPanic bool) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in Z3Main")
		}
	}()

	err := z.Prepare(p_paths, block_poses, boolCheckPanic)
	if err != nil {
		return false
	}
	z.Solver = z3.NewSolver(z.Z3Ctx)
	z.Assert()
	sat, err := z.Solver.Check()
	if err != nil {
		fmt.Println("Solving error:", err)
		return false
	}
	if sat == false {
		//fmt.Println("Unsatisfiable")
		return false
	}

	// found a trace!
	if config.Print_Debug_Info {
		fmt.Println("Satisfiable!")
		model := z.Solver.Model()
		for i, zthread := range z.vecZGoroutines {
			fmt.Println("------Thread ", i)
			for j, znode := range zthread.Nodes {
				fmt.Println("Node ", j, ":", TypeMsgForNode(znode.PNode().Node))
				order := model.Eval(znode.TraceOrder(), true)
				fmt.Println("\tOrder: ", order.String())
			}
		}
	}

	return true
}

// Initialize z.vecZGoroutines; generate ZNodes on each ZGoroutine from p_paths; generate constraints
func (z *Z3System) Prepare(vecPPath []*PPath, vecBlockPos []blockingPos, boolCheckPanic bool) error {

	// init z.vecZGoroutines
	for _, path := range vecPPath {
		Zpath := &ZGoroutine{
			PtrPPath:     path,
			Nodes:        nil,
			IsTerminated: true,
			BlockAt:      -1,
			Z3Sys:        z,
		}
		z.vecZGoroutines = append(z.vecZGoroutines, Zpath)
	}

	// Record which PNode are blocking. Later we will calculate ZGoroutine's BlockAt
	blockingPNodes := make(map[*PNode]struct{})
	for _, blockPos := range vecBlockPos {
		z.vecZGoroutines[blockPos.pathId].IsTerminated = false
		blockingPNodes[vecPPath[blockPos.pathId].Path[blockPos.pNodeId]] = struct{}{}
	}

	// from PNode in PPath, init ZNode. Ignore useless nodes according to z.Config
	// Calculate ZNode.IsBlocking
	var numZNode int64 = 0
	for _, Zthread := range z.vecZGoroutines {
		flagMetBlocking := false
		for _, pNode := range Zthread.PtrPPath.Path {

			var isBlocking, isOriginBlocking bool
			_, isOriginBlocking = blockingPNodes[pNode]
			isBlocking = isOriginBlocking || flagMetBlocking // this node is blocking if it is in the blockingPNodes or
			// it is after a blocking node

			var newZnode ZNode
			newBasic := ZNodeBasic{
				ZGoroutine:     Zthread,
				PtrPNode:       pNode,
				boolIsBlocking: isBlocking,
			}

			switch n := pNode.Node.(type) {
			case *Go:
				newZnode = &newBasic
			case SyncOp:
				if select_case, ok := n.(*SelectCase); ok {
					if select_case.BoolIsDefault {
						continue
					}
				}
				switch prim := n.Primitive().(type) {
				case *instinfo.Channel:
					//check if this channel is nil
					var nilChan *instinfo.Channel
					nilChan = nil
					if prim == nilChan {
						z.Warnings = append(z.Warnings, WARNING_met_nil_chan)
						continue
					}

					// check if this is a special channel
					if prim == &instinfo.ChanNotDepend || prim == &instinfo.ChanTimer || prim == &instinfo.ChanContext {
						continue
					}

					if prim.Buffer == 0 { // unbuffered channel
						switch n.Operation().(type) {
						case *instinfo.ChSend:
							newZnode = &ZNodeNbSend{
								ZNodeBasic: newBasic,
								Pairs:      nil,
								Closes:     nil,
							}
						case *instinfo.ChRecv:
							newZnode = &ZNodeNbRecv{
								ZNodeBasic: newBasic,
								Pairs:      nil,
								Closes:     nil,
							}
						case *instinfo.ChClose:
							newZnode = &ZNodeClose{
								newBasic,
							}
						case *instinfo.ChMake:
							continue
						}
					} else if prim.Buffer != instinfo.DynamicSize { // buffered channel, buffer is constant
						buffer := z.Z3Ctx.FromInt(int64(prim.Buffer), z.Z3Ctx.IntSort()).(z3.Int)
						switch n.Operation().(type) {
						case *instinfo.ChSend:
							newZnode = &ZNodeBSend{
								Buffer:     buffer,
								ZNodeBasic: newBasic,
							}
						case *instinfo.ChRecv:
							newZnode = &ZNodeBRecv{
								Buffer:     buffer,
								ZNodeBasic: newBasic,
							}
						case *instinfo.ChClose:
							newZnode = &ZNodeClose{
								newBasic,
							}
						case *instinfo.ChMake: // TODO: maybe we should check the buffer size again here
							continue
						}
					} else { // buffered channel, buffer is a variable that we can't analyze, we don't gen constraint for it
						z.Warnings = append(z.Warnings, WARNING_met_chan_var_buf)
						continue
					}
				default: // TODO: add other primitives
					if prim == nil {
						continue
					}
					continue
				}
			default:
				continue

			}

			Zthread.Nodes = append(Zthread.Nodes, newZnode)
			if isOriginBlocking {
				z.mapOriginBlockingNodes[newZnode] = struct{}{}
				flagMetBlocking = true
				Zthread.BlockAt = len(Zthread.Nodes) - 1
			}
		}
		numZNode += int64(len(Zthread.Nodes))
	}

	// set id of vecZGoroutines and ZNodes
	for i, Zthread := range z.vecZGoroutines {
		Zthread.ID = i
		for j, znode := range Zthread.Nodes {
			if znode == nil {
				fmt.Println("warning in z.Prepare: Nil ZNode when set id")
				continue
			}
			znode.SetId(j)
		}
	}


	// For a panic bug, we need at least one panic
	ZBoolHasPanic := z.Z3Ctx.FromBool(false)

	// fill Pairs for some ZNodes
	// each blocking sync op need a list of pairing ops
	for _, Zthread := range z.vecZGoroutines {
		for _, znode := range Zthread.Nodes {
			switch concrete := znode.(type) {
			case *ZNodeClose:
				if boolCheckPanic {
					// If there are multiple closes, we already found a panic bug
					if len(concrete.findAllThreadCloses()) > 1 {
						ZBoolHasPanic = z.Z3Ctx.FromBool(true)
					}
				}
			case *ZNodeNbSend:
				// get recv of the same channel on other threads
				vecOtherThreadRecvs := concrete.findOtherThreadRecvs()
				for _, recv := range vecOtherThreadRecvs {
					// see if send already has the pair to recv
					if concrete.hasPairWith(recv) == false {
						// create *Pair for this send and recv
						boolUseThisPair := z.Z3Ctx.BoolConst("Pair_" + znode_name(concrete) + znode_name(recv) + "_use")
						newPair := &ZSendRecvPair{
							boolUseThisPair: boolUseThisPair,
							Send:            concrete,
							Recv:            recv,
						}
						// add Pair to send and recv, and record in z
						concrete.Pairs = append(concrete.Pairs, newPair)
						recv.Pairs = append(recv.Pairs, newPair)
						z.vecZSendRecvPairs = append(z.vecZSendRecvPairs, newPair)
					}
				}
				if boolCheckPanic {
					// get closes of the same channel on any threads. Note that send will be influenced by close on the same thread
					vecAllThreadCloses := concrete.findAllThreadCloses()
					for _, zclose := range vecAllThreadCloses {
						concrete.Closes = append(concrete.Closes, zclose)
					}
				}
			case *ZNodeNbRecv:
				// get sends of the same channel on other threads
				vecOtherThreadSends := concrete.findOtherThreadSends()
				for _, send := range vecOtherThreadSends {
					// see if recv already has the pair to send
					if concrete.hasPairWith(send) == false {
						// create *Pair for this send and recv
						useThisPair := z.Z3Ctx.BoolConst("Pair_" + znode_name(send) + znode_name(concrete) + "_use")
						newPair := &ZSendRecvPair{
							boolUseThisPair: useThisPair,
							Send:            send,
							Recv:            concrete,
						}
						// add Pair to send and recv
						concrete.Pairs = append(concrete.Pairs, newPair)
						send.Pairs = append(send.Pairs, newPair)
						z.vecZSendRecvPairs = append(z.vecZSendRecvPairs, newPair)
					}
				}
				// get closes of the same channel on any threads. Note that receive will be influenced by close on the same thread
				vecAllThreadCloses := concrete.findAllThreadCloses()
				for _, zclose := range vecAllThreadCloses {
					concrete.Closes = append(concrete.Closes, zclose)
				}
			case *ZNodeBSend:
				// get all other send and recv of the same channel on all threads
				concrete.Other_SR = concrete.findAllThreadOtherSendRecv()
			case *ZNodeBRecv:
				// get all other send and recv of the same channel on all threads
				concrete.Other_SR = concrete.findAllThreadOtherSendRecv()
				concrete.Closes = concrete.findAllThreadCloses()
			}
		}
	}

	// generate order constraints
	// Order constraints: 1. order should in [0, numZNode). This is not mandatory, but may make Z3 search space smaller.
	//	However, blocking nodes will have Order numZNode, meaning they are still blocking after all other nodes are executed
	//						2. for node sequence ABCD, Order of B is greater than Order of A
	//						3. The first node should happen after the *Go that creates this thread
	Z3NumZNode := z.Z3Ctx.FromInt(numZNode, z.Z3Ctx.IntSort()).(z3.Int)
	for _, Zthread := range z.vecZGoroutines {
		// for each ZNode, encode the order
		for j, znode := range Zthread.Nodes {
			if znode == nil {
				fmt.Println("warning in z.Prepare: Nil ZNode when generate order constraints")
				continue
			}
			a := z.Z3Ctx.IntConst(znode_name(znode) + "_Order")
			znode.UpdateOrder(a)
			if _, isOriBlocking := z.mapOriginBlockingNodes[znode]; isOriBlocking {
				aEqNumZNode := a.Eq(Z3NumZNode)
				z.Constraints.Order = append(z.Constraints.Order, aEqNumZNode)
			} else {
				aGe0 := a.GE(Z3Zero)
				aLtNumZNode := a.LT(Z3NumZNode)
				z.Constraints.Order = append(z.Constraints.Order, aGe0, aLtNumZNode)
			}

			if j > 0 {
				prevNode := Zthread.Nodes[j-1]
				if prevNode == nil {
					continue
				}
				if prevNode.IsBlocking() {
					continue
				}
				aGtPrev := a.GT(prevNode.TraceOrder())
				z.Constraints.Order = append(z.Constraints.Order, aGtPrev)
			}
		}

		// encode the order of *Go
		if len(Zthread.Nodes) > 0 {
			firstZN := Zthread.Nodes[0]
			creatorN := firstZN.PNode().Node.CallCtx().Goroutine.Creator
			if creatorN != nil { // this thread has creator
				// find the creator ZNode
				creatorZNode := z.findZNodeMatchNode(creatorN)
				if creatorZNode != nil {
					if creatorZNode.IsBlocking() {
						return fmt.Errorf("Illegal combination: a goroutine's creator is not reached")
					}
					spawnOrder := firstZN.TraceOrder().GT(creatorZNode.TraceOrder())
					z.Constraints.Order = append(z.Constraints.Order, spawnOrder)
				} else {
					fmt.Println("Warning in Z3System.Prepare: failed to find the creator ZNode of a thread")
				}
			}
		}
	}

	// generate sync constraints
	// unbuffered channel:
	// 					send.Sync == one and only one of pair.Sync is true
	//					pair.Sync is true ==> (send.Order == recv.Order) && (send.Value == recv.Value)
	//					recv.Sync == (one and ...) XOR recv.from_close
	// 					recv.from_close == (close1.Order < recv.Order || close2.Order < recv.Order ...)
	//					recv.from_close is true ==> none of pair.Sync is true
	//					recv.from_close is true ==> recv.Value == default_value // TODO: value rules are not encoded
	// buffer channel:
	//					For each send of ch, consider all operations that happen earlier than send, then (Num_sends - Num_recvs) should be less than ch.Buffer_size
	//					For each recv of ch, …, then {(Num_sends - Num_recvs) should be larger than 0} or recv.from_close
	// 					recv.from_close == (close1.Order < recv.Order || close2.Order < recv.Order ...)
	for _, zthread := range z.vecZGoroutines {
		for _, znode := range zthread.Nodes {
			if snode, ok := znode.PNode().Node.(SyncOp); ok {
				if znode.PNode().Node.Parent().Task.IsPrimATarget(snode.Primitive()) == false {
					// This is a Sync Op of a primitive not in Task, meaning we don't have full information of this primitive, we should ignore it
					continue
				}
			}

			// TODO: some ZNode belong to primitives that we don't know all information about, and should be ignored
			switch concrete := znode.(type) {
			case *ZNodeNbSend:
				allPairChosen := []z3.Bool{}
				for _, pair := range concrete.Pairs {
					allPairChosen = append(allPairChosen, pair.boolUseThisPair)
				}
				var syncOfThisSend z3.Bool
				if znode.IsBlocking() {
					// a blocking op is not executed and can't sync with anyone
					// it may panic
					syncOfThisSend = z.noneIsTrue(allPairChosen)

					for _, ZClose := range concrete.Closes {
						ZBoolSendAfterClose := ZClose.Order.LE(concrete.Order)
						ZBoolHasPanic = ZBoolHasPanic.Or(ZBoolSendAfterClose)
					}

				} else {
					// sync of send == one and only one of pair.Sync is true
					syncOfThisSend = z.onlyOneTrue(allPairChosen)
				}
				z.Constraints.SyncOfOp = append(z.Constraints.SyncOfOp, syncOfThisSend)

			case *ZNodeNbRecv:
				// sync of recv = XOR(with_send, from_close)
				// with_send: one and only one of pair.Sync is true
				// from_close: none is true && (close1.Order < recv.Order || close2.Order < recv.Order ...)
				allPairChosen := []z3.Bool{}
				for _, pair := range concrete.Pairs {
					allPairChosen = append(allPairChosen, pair.boolUseThisPair)
				}
				var syncOfThisRecv z3.Bool

				if znode.IsBlocking() { // a blocking op is not executed, and can't sync with anyone
					syncOfThisRecv = z.noneIsTrue(allPairChosen)
				} else {
					syncWithSend := z.onlyOneTrue(allPairChosen)

					var allCloses []ZNode
					for _, zclose := range concrete.Closes {
						allCloses = append(allCloses, zclose)
					}
					anyCloseAlreadyHappen := z.anyNodeInListHappenBefore(allCloses, concrete)
					concrete.FromClose = anyCloseAlreadyHappen

					syncOfThisRecv = syncWithSend.Xor(concrete.FromClose)

					noneIsTrue := z.noneIsTrue(allPairChosen)
					fromCloseImpliesNoneIsTrue := anyCloseAlreadyHappen.Implies(noneIsTrue)
					z.Constraints.SyncOfOp = append(z.Constraints.SyncOfOp, fromCloseImpliesNoneIsTrue)
				}
				z.Constraints.SyncOfOp = append(z.Constraints.SyncOfOp, syncOfThisRecv)

			case *ZNodeBSend:
				if znode.IsBlocking() {
					continue
				} else {
					// consider all operations that happen earlier than send, then (Num_sends - Num_recvs) should be less than ch.Buffer_size
					nowBuffer := z.nowBuffer(concrete)

					lessThanBufferSize := nowBuffer.LT(concrete.Buffer)
					z.Constraints.SyncOfOp = append(z.Constraints.SyncOfOp, lessThanBufferSize)
				}

			case *ZNodeBRecv:
				if znode.IsBlocking() {
					continue
				} else {
					// consider all operations that happen earlier than recv,
					// 	then [(Num_sends - Num_recvs) should be less than ch.Buffer_size] or [recv.FromClose == true]
					nowBuffer := z.nowBuffer(concrete)
					largerThanZero := nowBuffer.GT(Z3Zero)

					var allCloses []ZNode
					for _, zclose := range concrete.Closes {
						allCloses = append(allCloses, zclose)
					}
					anyCloseAlreadyHappen := z.anyNodeInListHappenBefore(allCloses, concrete)
					concrete.From_close = anyCloseAlreadyHappen

					syncOfThisRecv := largerThanZero.Or(anyCloseAlreadyHappen)
					z.Constraints.SyncOfOp = append(z.Constraints.SyncOfOp, syncOfThisRecv)
				}
			}
		}
	}
	if boolCheckPanic { // requires a panic
		z.Constraints.SyncOfOp = append(z.Constraints.SyncOfOp, ZBoolHasPanic)
	}


	for _, SRpair := range z.vecZSendRecvPairs {
		// pair.Sync is true ==> (send.Order == recv.Order) && (send.Value == recv.Value)
		infer := SRpair.boolUseThisPair.Implies(SRpair.Send.Order.Eq(SRpair.Recv.Order))
		z.Constraints.PairInferRule = append(z.Constraints.PairInferRule, infer)
	}

	// Blocking constraints
	for bNode, _ := range z.mapOriginBlockingNodes {
		switch concrete := bNode.(type) {
		case *ZNodeNbSend:
			// Nothing
		case *ZNodeNbRecv:
			// Close must haven't happen
			var allCloses []ZNode
			for _, zclose := range concrete.Closes {
				allCloses = append(allCloses, zclose)
			}
			anyCloseAlreadyHappen := z.anyNodeInListHappenBefore(allCloses, concrete)
			allCloseNotHappen := anyCloseAlreadyHappen.Not()
			z.Constraints.Blocking = append(z.Constraints.Blocking, allCloseNotHappen)
		case *ZNodeBSend:
			// Buffer must be full
			nowBuffer := z.nowBuffer(concrete)
			bufferFull := nowBuffer.Eq(concrete.Buffer)
			z.Constraints.Blocking = append(z.Constraints.Blocking, bufferFull)
		case *ZNodeBRecv:
			// Buffer must be zero
			nowBuffer := z.nowBuffer(concrete)
			bufferEmpty := nowBuffer.Eq(Z3Zero)
			z.Constraints.Blocking = append(z.Constraints.Blocking, bufferEmpty)

			// Close must haven't happen
			var allCloses []ZNode
			for _, zclose := range concrete.Closes {
				allCloses = append(allCloses, zclose)
			}
			anyCloseAlreadyHappen := z.anyNodeInListHappenBefore(allCloses, concrete)
			allCloseNotHappen := anyCloseAlreadyHappen.Not()
			z.Constraints.Blocking = append(z.Constraints.Blocking, allCloseNotHappen)
		}
	}

	// no need to generate blocking constraints (two blocking ops can't unlock each other)
	return nil
}

// Assert puts assertions stored in z.Contraints into the solver
func (z *Z3System) Assert() {
	for _, orderC := range z.Constraints.Order {
		z.Solver.Assert(orderC)
	}
	for _, opSyncC := range z.Constraints.SyncOfOp {
		z.Solver.Assert(opSyncC)
	}
	for _, pairInferC := range z.Constraints.PairInferRule {
		z.Solver.Assert(pairInferC)
	}
	for _, blockingC := range z.Constraints.Blocking {
		z.Solver.Assert(blockingC)
	}
}

func (z *Z3System) PrintAssert() {
	for _, orderC := range z.Constraints.Order {
		fmt.Println(orderC.String())
	}
	for _, opSyncC := range z.Constraints.SyncOfOp {
		fmt.Println(opSyncC.String())
	}
	for _, pairInferC := range z.Constraints.PairInferRule {
		fmt.Println(pairInferC.String())
	}
	for _, blockingC := range z.Constraints.Blocking {
		fmt.Println(blockingC.String())
	}
}
