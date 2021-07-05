package syncgraph

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/output"
	"os"
)

func ReportNoViolation() {
	fmt.Println("Finished verification: The program has no channel safety or liveness violations")
}

func ReportNotSure() {
	fmt.Println("Finished verification: Can't decide whether the program has channel safety or liveness violations")
}

func ReportViolation() {
	fmt.Println("Finished verification: The program has channel safety or liveness violations")
}

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

		// Check if the program has double close
		vecClose := []*instinfo.ChClose{}
		for _, pPath := range paths {
			for _, pNode := range pPath.Path {
				syncNode, ok := pNode.Node.(SyncOp)
				if !ok {
					continue
				}
				if g.Task.IsPrimATarget(syncNode.Primitive()) {
					if op, ok := syncNode.(*ChanOp);ok {
						if chClose, ok := op.Op.(*instinfo.ChClose); ok {
							vecClose = append(vecClose, chClose)
						}
					}
				}
			}
		}

		for _, aClose := range vecClose {
			for _, bClose := range vecClose {
				if aClose == bClose {
					continue
				}
				if aClose.Parent == bClose.Parent {
					config.BugIndex++
					fmt.Print("----------Bug[")
					fmt.Print(config.BugIndex)
					fmt.Print("]----------\n\tType: Channel Safety \tReason: Double close.\n")
					ReportViolation()
					os.Exit(1)
				}
			}
		}

		// List all blocking op of target channel on any path
		pathId2AllBlockPos := make(map[int][]blockingPos)

		const emptyPNodeId = -2
		for i, pPath := range paths {
			for j, pNode := range pPath.Path {
				syncNode, ok := pNode.Node.(SyncOp)
				if !ok {
					continue
				}
				if g.Task.IsPrimATarget(syncNode.Primitive()) {
					if canSyncOpTriggerGl(syncNode) {
						newBlockPos := blockingPos{
							pathId:  i,
							pNodeId: j,
						}
						pathId2AllBlockPos[i] = append(pathId2AllBlockPos[i], newBlockPos)
					}
				}
			}
			emptyBlockPos := blockingPos{
				pathId:  i,
				pNodeId: emptyPNodeId,
			}
			pathId2AllBlockPos[i] = append(pathId2AllBlockPos[i], emptyBlockPos)
		}

		allBlockPosComb := []map[int]blockingPos{}

		mapIndices := make(map[int]int)
		for pathId, _ := range pathId2AllBlockPos {
			mapIndices[pathId] = 0
		}
		for {
			newComb := make(map[int]blockingPos)
			boolCanSync := false
			for pathId, indice := range mapIndices {
				blockPos := pathId2AllBlockPos[pathId][indice]
				if blockPos.pNodeId != emptyPNodeId {
					// check if blockPos can sync with a previous blockPos
					thisSyncNode, ok := paths[pathId].Path[blockPos.pNodeId].Node.(SyncOp)
					if !ok {
						fmt.Println("Panic when enumerate blockPos combination: Node is not SyncOp")
						panic(1)
					}
					for pathId, otherBlockPos := range newComb {
						if otherBlockPos.pNodeId == emptyPNodeId {
							continue
						}
						pPath := paths[pathId]
						otherSyncNode, ok := pPath.Path[otherBlockPos.pNodeId].Node.(SyncOp)
						if !ok {
							fmt.Println("Panic when enumerate blockPos combination: Node is not SyncOp")
							panic(1)
						}
						if thisSyncNode.Primitive() != otherSyncNode.Primitive() {
							continue
						}
						if canSync(thisSyncNode, otherSyncNode) {
							boolCanSync = true
							break
						}
					}
				}
				newComb[pathId] = blockPos
			}
			if boolCanSync == false {
				boolAllEmpty := true
				for _, blockPos := range newComb {
					if blockPos.pNodeId != emptyPNodeId {
						boolAllEmpty = false
						break
					}
				}
				if boolAllEmpty == false {
					allBlockPosComb = append(allBlockPosComb, newComb)
				}
			}

			nextPathId := -1
			for pathId, indice := range mapIndices {
				if indice >= len(pathId2AllBlockPos[pathId])-1 {
					continue
				} else {
					nextPathId = pathId
					break
				}
			}

			if nextPathId == -1 {
				break
			}

			mapIndices[nextPathId] += 1

			for pathId, _ := range mapIndices {
				if pathId == nextPathId {
					break
				}
				mapIndices[pathId] = 0
			}
		}

		// For every blocking op of target channel on any path
		for i := 0; i < len(allBlockPosComb); i++ {
			var blockPosComb map[int]blockingPos
			blockPosComb = allBlockPosComb[i]

			for _, blockPos := range blockPosComb {
				if blockPos.pNodeId != emptyPNodeId {
					inst := paths[blockPos.pathId].Path[blockPos.pNodeId].Node.Instruction()
					str := output.StringIISrc(inst)
					if _, printed := PrintedBlockPosStr[str]; printed {
						return false
					}
				}
			}
			// Make some paths block and other paths exit
			for i, path := range paths {
				blockPos, exist := blockPosComb[i]
				if exist && blockPos.pNodeId != emptyPNodeId {
					path.SetBlockAt(blockPos.pNodeId)
				} else {
					path.SetAllReached()
				}
			}

			// See if Sync-rule is satisfied. Sync-rule: the number of ops of one prim must match, except the blocking one
			flagSyncRuleSatisfied := true

			mapPrim2OpNodes := make(map[interface{}][]*PNode) // A map from target primitive to its reached ops in these paths
			for i := 0; i < len(goroutines); i++ {
				path := paths[i]
				for _, pNode := range path.Path {
					syncNode, ok := pNode.Node.(SyncOp)
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
					flagSyncRuleSatisfied = checkChOpsLegal(prim, nodes)
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

			vecBlockingPos := []blockingPos{}
			for _, blockPos := range blockPosComb {
				if blockPos.pNodeId == emptyPNodeId {
					continue
				}
				vecBlockingPos = append(vecBlockingPos, blockPos)
			}

			z3Sys_block := NewZ3ForGl()
			z3Sat_block := z3Sys_block.Z3Main(paths, vecBlockingPos, false) // only check liveness problem

			z3Sys_panic := NewZ3ForGl()
			z3Sat_panic := z3Sys_panic.Z3Main(paths, vecBlockingPos, true) // check safety problem

			// Report a bug
			if z3Sat_block || z3Sat_panic {
				//z3Sys.PrintAssert()
				config.BugIndexMu.Lock()
				config.BugIndex++
				fmt.Print("----------Bug[")
				fmt.Print(config.BugIndex)
				config.BugIndexMu.Unlock()
				if z3Sat_panic { // panic bugs have priority, because they may be misunderstood by Z3 as blocking bugs
					fmt.Print("]----------\n\tType: Channel Safety \tReason: Send after close or double close.\n")
				} else {
					fmt.Print("]----------\n\tType: BMOC \tReason: One or multiple channel operation is blocked.\n")
				}
				fmt.Println("-----Blocking at:")
				for _, blockPos := range blockPosComb {
					if blockPos.pNodeId != emptyPNodeId {
						inst := paths[blockPos.pathId].Path[blockPos.pNodeId].Node.Instruction()
						str := output.StringIISrc(inst)
						fmt.Print(str)
						PrintedBlockPosStr[str] = struct{}{}
					}
				}

				for _, blockPos := range blockPosComb {
					if blockPos.pNodeId != emptyPNodeId {
						fmt.Println("-----Blocking Path NO.", blockPos.pathId)
						paths[blockPos.pathId].PrintPPath()
					} else {
						fmt.Println("-----Path NO.", blockPos.pathId, "\tEntry func at:", goroutines[blockPos.pathId].EntryFn.String())
						paths[blockPos.pathId].PrintPPath()
					}
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

var PrintedBlockPosStr map[string]struct{} = make(map[string]struct{})
