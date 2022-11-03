package path

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/tools/go/callgraph"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"strings"
)

func NewEdgeChain(startNode *callgraph.Node) *EdgeChain {
	return &EdgeChain{
		Chain: nil,
		Start: startNode,
	}
}

func IsCallGraphAccurateOnNode(node *callgraph.Node) bool {
	return len(node.In) > 1
}

func BacktraceCallChain(source *callgraph.Node, visited map[int]*callgraph.Edge) (result *EdgeChain) {
	result = NewEdgeChain(source)
	for visited[source.ID] != nil {
		edge := visited[source.ID]
		result.Chain = append(result.Chain, edge)
		source = edge.Callee
	}
	return
}

func ComputeScope(funcs []*ssa.Function, lcaConfig *LcaConfig) (map[*ssa.Function][]*EdgeChain, error) {
	ret := make(map[*ssa.Function][]*EdgeChain)
	var err error
	visitedFuncs := make(map[*ssa.Function]struct{})
	for _, targetFunc := range funcs {
		if _, ok := visitedFuncs[targetFunc]; !ok {
			var callchain *EdgeChain
			callchain, err = ComputeCallChain(config.CallGraph.Nodes[targetFunc], lcaConfig)
			if err != nil {
				return nil, err
			}
			if _, ok := ret[callchain.Start.Func]; !ok {
				ret[callchain.Start.Func] = []*EdgeChain{}
			}
			ret[callchain.Start.Func] = append(ret[callchain.Start.Func], callchain)
			visitedFuncs[targetFunc] = struct{}{}
		}
	}
	return ret, err
}

// ComputeCallChain computes the shortest call chain starting from an entry function to the sink. Once it finds an entry
// function, it stops search and returns.
func ComputeCallChain(sink *callgraph.Node, config *LcaConfig) (result *EdgeChain, err error) {
	// key is the node id in source, value is the predecessor in the call chain.
	err = fmt.Errorf("Call chain not found for " + sink.Func.Name())
	visited := make(map[int]*callgraph.Edge)
	queue := make([]*callgraph.Node, 1)
	queue[0] = sink
	visited[sink.ID] = nil
	head := -1
	for head < len(queue)-1 {
		head += 1
		headNode := queue[head]
		if IsCallGraphAccurateOnNode(headNode) {
			if config.GiveUpWhenCallGraphIsInaccurate {
				return nil, ErrInaccurateCallgraph
			}
			//fmt.Println(ErrInaccurateCallgraph)
		}
		if headNode.Func.Name() == "main" ||
			(strings.HasPrefix(headNode.Func.Name(), "Test") &&
				!strings.Contains(headNode.Func.Name(), "$")) {
			return BacktraceCallChain(headNode, visited), nil
		}
		for _, calleeEdge := range headNode.In {
			caller := calleeEdge.Caller
			if !IsFunctionIncludedInAnalysis(caller) {
				continue
			}
			if _, ok := visited[caller.ID]; !ok {
				visited[caller.ID] = calleeEdge
				queue = append(queue, caller)
			}
		}
	}
	return
}