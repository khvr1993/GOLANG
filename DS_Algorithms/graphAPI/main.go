package main

import (
	"fmt"
	"github.com/khvr1993/GOLANG/DS_Algorithms/graphAPI/graph"
	"github.com/khvr1993/GOLANG/DS_Algorithms/graphAPI/graphProcessing"
	"log"
)

func main() {
	log.Println("Testing Graph API")
	newGraph := graph.NewGraph(13)
	newGraph.AddEdge(0, 5)
	newGraph.AddEdge(4, 3)
	newGraph.AddEdge(0, 1)
	newGraph.AddEdge(9, 12)
	newGraph.AddEdge(6, 4)
	newGraph.AddEdge(5, 4)
	newGraph.AddEdge(0, 2)
	newGraph.AddEdge(11, 12)
	newGraph.AddEdge(9, 10)
	newGraph.AddEdge(0, 6)
	newGraph.AddEdge(7, 8)
	newGraph.AddEdge(9, 11)
	newGraph.AddEdge(5, 3)
	for v := 0; v < newGraph.V; v++ {
		fmt.Print(v, " => ")

		for k := range newGraph.Iterable(v) {
			fmt.Print(" ", k, " ")
		}
		fmt.Println()
	}

	depthFirstPath := graphProcessing.NewDepthFirstPaths(*newGraph, 0)
	depthFirstPath.DepthFirstPaths(*newGraph, 0)
	//log.Println(depthFirstPath.HasPathTo(2))
	edGeconns := depthFirstPath.ReturnEdgeConnections()
	for k, v := range *(edGeconns) {
		fmt.Print("k => ", k, " v => ", v)
		fmt.Println()
	}
}
