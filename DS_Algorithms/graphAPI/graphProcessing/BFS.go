package graphProcessing

import (
	"github.com/khvr1993/GOLANG/DS_Algorithms/graphAPI/graph"
	queue2 "github.com/khvr1993/GOLANG/DS_Algorithms/queue/queue"
	"log"
)

type BreadthFirstProcessor struct {
	marked []bool
	edgeTo []int
	s      int
}

// NewBreadthFirstPaths acts as a constructor and returns the pointer to a struct which stores
// all the variables required to do DFS
func NewBreadthFirstPaths(graph graph.Graph, s int) *DepthFirstProcessor {
	breadthFirstPaths := new(DepthFirstProcessor)
	breadthFirstPaths.marked = make([]bool, graph.V)
	for i := range breadthFirstPaths.marked {
		breadthFirstPaths.marked[i] = false
	}
	breadthFirstPaths.edgeTo = make([]int, graph.V)
	for i := range breadthFirstPaths.edgeTo {
		breadthFirstPaths.edgeTo[i] = -1
	}
	breadthFirstPaths.s = s
	return breadthFirstPaths
}

// bfs performs breadth first search
func (bfs *BreadthFirstProcessor) BreadthFirstPaths(graph graph.Graph, sourceVertex int) {
	bfs.bfs(graph, sourceVertex)
}

/*
	bfs performs the following
	1. Step 1 is to insert the source vertex into a queue
	2. For that vertex dequeue it and mark all the vertices which are adjacent to it as visited if they
		are not visited earlier and insert them into the queue
	3. repeat this process till the queue becomes empty
*/
func (bfs *BreadthFirstProcessor) bfs(graph graph.Graph, s int) {
	queue := queue2.Queue{}
	queue.Enqueue(s)
	bfs.marked[s] = true
	for queue.Size() > 0 {
		vertex := queue.Dequeue()
		for i := range graph.Iterable(vertex) {
			log.Println("Checking for vertex ", i)
			if !bfs.marked[i.(int)] {
				queue.Enqueue(i.(int))
				bfs.marked[i.(int)] = true
				bfs.edgeTo[i.(int)] = vertex
			}
		}
	}
}
