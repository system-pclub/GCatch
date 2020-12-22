package syncgraph

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/tools/github.com/aclements/go-z3/z3"
	"strconv"
)

// given a slice of *z3.Bool, onlyOneTrue returns a z3.Bool that is equal to "one and only one of the slice is true"
func (z *Z3System) onlyOneTrue(list []z3.Bool) (result z3.Bool) {
	result = z.Z3Ctx.FromBool(false)
	for i, b1 := range list {
		// obtain a bool representing: b1 is true and all others are false
		b1_T_other_false := b1
		for j, b2 := range list {
			if i == j {
				continue
			}
			b1_T_other_false = b1_T_other_false.And(b2.Not())
		}
		result = result.Or(b1_T_other_false)
	}
	return
}

// given a slice of *z3.Bool, noneIsTrue returns a z3.Bool that is equal to "none of the slice is true"
func (z *Z3System) noneIsTrue(list []z3.Bool) (result z3.Bool) {
	result = z.Z3Ctx.FromBool(true)
	for _, b := range list {
		result = result.And(b.Not())
	}
	return
}

func (z *Z3System) anyNodeInListHappenBefore(list []ZNode, target ZNode) (result z3.Bool) {
	result = z.Z3Ctx.FromBool(false)
	for _, any := range list {
		happenBefore := any.TraceOrder().LT(target.TraceOrder())
		result = result.Or(happenBefore)
	}
	return
}

func znode_name(node ZNode) string {
	return "G_" + strconv.Itoa(node.Goroutine().ID) + "_N_" + strconv.Itoa(node.ID())
}

func (z *Z3System) findZNodeMatchNode(n Node) ZNode {
	for _, Zthread := range z.vecZGoroutines {
		for _, Znode := range Zthread.Nodes {
			if Znode.PNode().Node == n {
				return Znode
			}
		}
	}
	return nil
}

// When you want a new name for a z3 const, uniqueName can give you a unique name
func (z *Z3System) uniqueName() string {
	z.countUniqueName += 1
	return "Unique_" + strconv.Itoa(z.countUniqueName)
}

func (z *Z3System) nowBuffer(op ZNode) z3.Int {
	other_SR := []ZNode{}

	switch concrete := op.(type) {
	case *ZNodeBSend:
		other_SR = concrete.Other_SR
	case *ZNodeBRecv:
		other_SR = concrete.Other_SR
	}

	nowBuffer := Z3Zero
	for _, other := range other_SR {
		prevBuffer := nowBuffer
		nowBuffer = z.Z3Ctx.IntConst(z.uniqueName())

		var cond, then, else_ z3.Bool
		cond = other.TraceOrder().LT(op.TraceOrder())
		if _, isSend := other.(*ZNodeBSend); isSend {
			then = nowBuffer.Eq(prevBuffer.Add(Z3One))
		} else if _, is_recv := other.(*ZNodeBRecv); is_recv { // receive
			then = nowBuffer.Eq(prevBuffer.Sub(Z3One))
		} else {
			fmt.Println("Fatal error in z.nowBuffer: other is not send or recv")
		}
		else_ = nowBuffer.Eq(prevBuffer)

		newIfThenElse := cond.IfThenElse(then, else_).(z3.Bool)
		z.Constraints.SyncOfOp = append(z.Constraints.SyncOfOp, newIfThenElse)
	}

	return nowBuffer
}
