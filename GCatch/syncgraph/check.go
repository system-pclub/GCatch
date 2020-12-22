package syncgraph

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/output"
)

type blockingPos struct {
	pathId  int
	pNodeId int
}

func (g SyncGraph) CheckWithZ3() bool {


	countBlockPoint := 0
	// Main loop: for each pathCombination
	for _, pathCombination := range g.PathCombinations {
		// If the pathCombination satisfies: the MainGoroutine's path is nil, skip this pathCombination

		// Store path's nodes into a slice of PNode. Make sure that each PNode is unique, though PNode.Node may appear multiple times
		goroutines := []*Goroutine{}
		paths := []*PPath{}
		for _, goPath := range pathCombination.go_paths {
			goroutines = append(goroutines, goPath.goroutine)

			vecNewPNode := []*PNode{}
			for i, oldNode := range goPath.path.Path {
				newNode := &PNode{
					Path:     goPath.path,
					Index:    i,
					Node:     oldNode,
					Blocked:  false,
					Executed: false,
				}
				vecNewPNode = append(vecNewPNode, newNode)
			}

			newPath := &PPath{
				Path:      vecNewPNode,
				localPath: goPath.path,
			}
			paths = append(paths, newPath)
		}

		// List all blocking op of target channel on any path
		allBlockPos := []blockingPos{}

		for i, pPath := range paths {
			for j, pNode := range pPath.Path {
				sync_node,ok := pNode.Node.(SyncOp)
				if !ok {
					continue
				}
				if g.Task.IsPrimATarget(sync_node.Primitive()) {
					if canSyncOpTriggerGl(sync_node) {
						new_p_pos := blockingPos{
							pathId:  i,
							pNodeId: j,
						}
						allBlockPos = append(allBlockPos, new_p_pos)
					}
				}
			}
		}

		// For every blocking op of target channel on any path
		for _, blockPos := range allBlockPos {
			// Make it block and other paths exit
			for i,path := range paths {
				if i == blockPos.pathId {
					path.SetBlockAt(blockPos.pNodeId)
				} else {
					path.SetAllReached()
				}
			}

			//// See if Go-rule is satisfied. Go-rule: A. if the goroutine is g.MainGoroutine, its path mustn't be nil
			////										 B. if a goroutine has no nil path, its creation must be on another path in pathCombination
			////										 C. if a Go is reached, the goroutine created must not be nil
			////										 TODO: number of goroutines need to be the same as the number of Go (think of Go in loop)
			//flag_Go_rule_satisfied := true
			//outer:
			//for i := 0; i < len(goroutines); i++ {
			//	goroutine := goroutines[i]
			//	path := paths[i]
			//
			//	// check A
			//	if goroutine == g.MainGoroutine && path.LocalPath.IsEmpty() {
			//		flag_Go_rule_satisfied = false
			//		break outer
			//	}
			//
			//	// check B
			//	if path.LocalPath.IsEmpty() == false && goroutine.Creator != nil { // No nil path, and not main goroutine.
			//		// Need to verify this goroutine is created by another path
			//		flag_found_creator := false
			//		for j, other := range paths {
			//			if i == j {
			//				continue
			//			}
			//			if other.IsNodeIn(goroutine.Creator) {
			//				flag_found_creator = true
			//				break
			//			}
			//		}
			//		if flag_found_creator == false {
			//			flag_Go_rule_satisfied = false
			//			break outer
			//		}
			//	}
			//
			//	// check C
			//	for _,node := range path.Path {
			//		node_go,ok := node.Node.(*Go)
			//		if !ok {
			//			continue
			//		}
			//
			//		flag_found_created_goroutine := false
			//		for j,_ := range goroutines {
			//			other := goroutines[j]
			//			other_path := paths[j]
			//			if other == goroutine {
			//				continue
			//			}
			//			if other.Creator == node_go	{ // this goroutine is created by our Go
			//				flag_found_created_goroutine = true
			//				if other_path.LocalPath.IsEmpty() {
			//					flag_Go_rule_satisfied = false
			//					break outer
			//				}
			//			}
			//		}
			//		if flag_found_created_goroutine == false {
			//			flag_Go_rule_satisfied = false
			//			break
			//		}
			//	}
			//}
			//if flag_Go_rule_satisfied == false {
			//	continue
			//}

			// See if Sync-rule is satisfied. Sync-rule: the number of ops of one prim must match, except the blocking one
			flagSyncRuleSatisfied := true

			mapPrim2OpNodes := make(map[interface{}][]*PNode) // A map from target primitive to its reached ops in these paths
			for i := 0; i < len(goroutines); i++ {
				path := paths[i]
				for _, pNode := range path.Path {
					syncNode,ok := pNode.Node.(SyncOp)
					if !ok {
						continue
					}
					if pNode.Executed == false {
						continue
					}
					prim := syncNode.Primitive()
					if g.Task.IsPrimATarget(prim) == false { // Only consider prim that is in Task.Target
						continue
					}
					mapPrim2OpNodes[prim] = append(mapPrim2OpNodes[prim], pNode)
				}
			}

			for p, vecOpNodes := range mapPrim2OpNodes {
				nodes := []Node{}
				for _, pNode := range vecOpNodes {
					nodes = append(nodes, pNode.Node)
				}

				switch prim := p.(type) {
				case *instinfo.Channel:
					flagSyncRuleSatisfied = checkChOpsLegal(prim,nodes)
				case *instinfo.Locker:
					// Do we really need a rule for Locker?
				}
				if flagSyncRuleSatisfied == false {
					break
				}
			}

			if flagSyncRuleSatisfied == false {
				continue
			}

			vecBlockingPos := []blockingPos{blockPos}
			z3Sys := NewZ3ForGl()
			z3Sat := z3Sys.Z3Main(paths, vecBlockingPos)

			// Report a bug
			if z3Sat {
				z3Sys.PrintAssert()
				fmt.Println("-------Confirmed blocking path ")
				fmt.Println("-----Blocking at:")
				inst := paths[blockPos.pathId].Path[blockPos.pNodeId].Node.Instruction()
				output.PrintIISrc(inst)

				fmt.Println("-----Blocking Path:")
				paths[blockPos.pathId].PrintPPath()

				for i, path := range paths {
					if i == blockPos.pathId {
						continue
					}
					fmt.Println("-----Path NO.",i,"\tEntry func at:",goroutines[i].EntryFn.String())
					path.PrintPPath()
				}
				fmt.Println()

				return true
			}

			countBlockPoint++
		}
	}

	//fmt.Println("=========Total path sets:",countBlockPoint)
	//output.Wait_for_input()
	return false
}
