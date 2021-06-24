package graphProcessing

import (
	stack2 "github.com/khvr1993/GOLANG/DS_Algorithms/Stack/Stack"
	"github.com/khvr1993/GOLANG/DS_Algorithms/graphAPI/graph"
	"log"
)

/*
	Marked => If there is a path from source to vertex then marked will store true
	edgeTo => From which vertex did we reach the querying vertex
	s => Source vertex
*/
type DepthFirstProcessor struct {
	marked []bool
	edgeTo []int
	s      int
}

// NewDepthFirstPaths acts as a constructor and returns the pointer to a struct which stores
// all the variables required to do DFS
func NewDepthFirstPaths(graph graph.Graph, s int) *DepthFirstProcessor {
	depthFirstPaths := new(DepthFirstProcessor)
	depthFirstPaths.marked = make([]bool, graph.V)
	for i := range depthFirstPaths.marked {
		depthFirstPaths.marked[i] = false
	}
	depthFirstPaths.edgeTo = make([]int, graph.V)
	for i := range depthFirstPaths.edgeTo {
		depthFirstPaths.edgeTo[i] = -1
	}
	depthFirstPaths.s = s
	return depthFirstPaths
}

func (dfs *DepthFirstProcessor) DepthFirstPaths(graph graph.Graph, s int) {
	dfs.dfs(graph, s)
}

/* dfs performs the depth first search using recursion. For a given vertex mark visited as true
and visit the adjacent vertices. If the adjacent vertex is already visited then come back and go to next vertex
*/
func (dfs *DepthFirstProcessor) dfs(graph graph.Graph, vertex int) {
	log.Println("Working on vertex ", vertex)
	///Marking this as true since the vertex is visited
	dfs.marked[vertex] = true

	//for the adjacent elements of this vertex we need to visit them and mark them as true
	for w := range graph.Iterable(vertex) {
		log.Println("Checking the vertex ", w)
		if !dfs.marked[w.(int)] {
			dfs.edgeTo[w.(int)] = vertex
			log.Println("Call Recusion for ", w)
			dfs.dfs(graph, w.(int))
		} else {
			log.Println("Already visited vertex ", w)
		}
	}
}

// HasPathTo returns true if a vertex can be reached from source
func (dfs *DepthFirstProcessor) HasPathTo(v int) bool {
	return dfs.marked[v]
}

//Pathto performs the follwing operations
//1. First checks if there is a path from source to the given vertex. If No then return Nil
// If there is a path then keep on adding the edge elements to a stack
// for vertex x the edge which connects to it is stored in edgeTo[x]
// first we insert x then check from where did we reach x and from where did we reach to that vertex and so on

func (dfs *DepthFirstProcessor) PathTo(v int) *stack2.InterfaceImplStack {
	if !dfs.HasPathTo(v) {
		emptyStack := stack2.NewStack()
		return emptyStack
	}
	path := stack2.NewStack()
	for x := v; x != dfs.s; x = dfs.edgeTo[x] {
		path.Push(dfs.edgeTo[v])
	}
	path.Push(dfs.s)
	return path

}

func (dfs *DepthFirstProcessor) ReturnEdgeConnections() *[]int {
	return &dfs.edgeTo
}
