package syncgraph

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/output"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"strconv"
	"strings"
)


type tupleGoroutinePath struct {
	goroutine *Goroutine
	path *LocalPath
}

type pathCombination struct { // []'s length is len(g.Goroutines). Every tupleGoroutinePath.Goroutine is unique
	go_paths []*tupleGoroutinePath
}


type EnumeConfigure struct {
	Unfold             int
	IgnoreFn           map[*ssa.Function]struct{} // a map of not interesting
	FlagIgnoreNoSyncFn bool
	FlagIgnoreNormal   bool
}

type LocalPath struct {
	Path                   []Node
	Hash                   string
	mapNodeEdge2IntVisited map[*NodeEdge]int
	//finished bool
}


var intEmptyPathId int
func NewEmptyPath() *LocalPath {
	intEmptyPathId++
	return &LocalPath{
		Path:                   []Node{},
		Hash:                   "Empty_path_NO." + strconv.Itoa(intEmptyPathId),
		mapNodeEdge2IntVisited: nil,
	}
}

func (l *LocalPath) IsEmpty() bool {
	return strings.HasPrefix(l.Hash, "Empty_path_NO.")
}

type PNode struct { // A Node will have n PNode if it shows up n times in path
	Path     *LocalPath
	Index    int
	Node     Node
	Blocked  bool
	Executed bool
}

type PPath struct {
	Path      []*PNode
	localPath *LocalPath
}


func (p PPath) IsNodeIn(node Node) bool {
	for _,t_n := range p.Path {
		if t_n.Node == node {
			return true
		}
	}
	return false
}

func (p PPath) SetAllReached() {
	for _,n := range p.Path {
		n.Blocked = false
		n.Executed = true
	}
}

func (p PPath) SetBlockAt(index int) {
	for i,n := range p.Path {
		if i < index {
			n.Executed = true
			n.Blocked = false
		} else if i == index {
			n.Executed = false
			n.Blocked = true
		} else {
			n.Executed = false
			n.Blocked = false
		}
	}
}

var mapHash2Map map[string]*LocalPath

type tupleCallerCallee struct {
	caller Node
	callee Node
}

func deleteNormalFromPath(oldPath *LocalPath) *LocalPath {
	newSlice := []Node{}
	for _, node := range oldPath.Path {
		if _, boolIsNormal := node.(*NormalInst); boolIsNormal {
			continue
		}
		newSlice = append(newSlice, node)
	}

	newLocalPath := &LocalPath{
		Path:                   newSlice,
		Hash:                   hashOfPath(newSlice),
		mapNodeEdge2IntVisited: copyBackedgeMap(oldPath.mapNodeEdge2IntVisited),
	}
	oldPath = nil
	return newLocalPath
}

func copyLocalPath(old *LocalPath) *LocalPath {
	newLocalPath := &LocalPath{
		Path:                   copyPathSlice(old.Path),
		Hash:                   old.Hash,
		mapNodeEdge2IntVisited: copyBackedgeMap(old.mapNodeEdge2IntVisited),
	}
	return newLocalPath
}

func debugPrintEnumeratedPaths(path_map map[string]*LocalPath) {
	count := 0
	fmt.Println("In total:", len(path_map))
	for _,path := range path_map {
		fmt.Println("-----Path:", count)
		fmt.Println(path.mapNodeEdge2IntVisited)
		count++
		for i,n := range path.Path {
			str := TypeMsgForNode(n)
			if str == "Normal_inst" {
				continue
			}
			fmt.Println(str)
			if i < len(path.Path) - 1 {
				var flag_backedge bool
				next := path.Path[i+1]
				for _,out := range n.Out() {
					if out.Succ == next {
						flag_backedge = out.IsBackedge
						break
					}
				}
				if flag_backedge {
					fmt.Println("--Backedge")
				}
			}
			if i == len(path.Path) - 1 {
				if TypeMsgForNode(n) != "Return" {
					output.WaitForInput()
				}
			}
		}
		output.WaitForInput()
	}
}


func copyBackedgeMap(old map[*NodeEdge]int) map[*NodeEdge]int {
	n := make(map[*NodeEdge]int)
	for key,value := range old {
		n[key] = value
	}
	return n
}

func copyPathMap(old map[string]*LocalPath) map[string]*LocalPath {
	n := make(map[string]*LocalPath)
	for key,value := range old {
		n[key] = value
	}
	return n
}

func copyPathSlice(old []Node) []Node {
	copy_path := []Node{}
	for _,n := range old {
		copy_path = append(copy_path, n)
	}
	return copy_path
}

func copyIntSlice(old []int) []int {
	n := []int{}
	for _, o := range old {
		n = append(n, o)
	}
	return n
}

func hashOfPath(node []Node) string {
	var buffer bytes.Buffer
	for _,n := range node {
		buffer.WriteString(strconv.Itoa(n.GetId())+" ")
	}
	byte_key := buffer.Bytes()
	hash := sha256.Sum256(byte_key)
	return string(hash[:])
}

func removeFromPathWorklist(old_worklist []*LocalPath, remove *LocalPath) []*LocalPath {
	result := []*LocalPath{}
	for _,o := range old_worklist {
		if o == remove {
			continue
		}
		result = append(result, o)
	}
	return result
}
