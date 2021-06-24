package graph

import (
	"github.com/khvr1993/GOLANG/DS_Algorithms/bag/Bag"
	"log"
)

/*
	Create a bags for vertices
*/
type Graph struct {
	AdjVertice []Bag.Bag
	V          int
}

/*
	Constructor to create the adj Vertice
*/
func NewGraph(V int) *Graph {
	log.Println("Begin NewGraph")
	s := &Graph{}
	s.AdjVertice = make([]Bag.Bag, 13)
	s.V = V
	s.graph(V)
	return s
}

/*
	for a given number of vertices the func initializes the bags and assigns them
*/
func (s *Graph) graph(V int) {

	//For every vertex we create a bag type data and assign it to it
	for i := 0; i < V; i++ {
		/*
			Adj[i] will store the pointer to the Bag for the i vertex
		*/
		bag := Bag.NewBag()
		s.AdjVertice[i] = *bag
	}
}

/*
	This function will create links for v and w by pushing the data into bag
	Imagine it as a line with 2 end points and from which ever side you take you
	should see the other point
*/
func (s *Graph) AddEdge(v int, w int) {
	s.AdjVertice[v].AddItem(w)
	s.AdjVertice[w].AddItem(v)
}

/*
GetBag Returns the Bag for the requested vertex. This Bag will have all the elements
that are adjacent to the vertex*/
func (s *Graph) GetBag(v int) Bag.Bag {
	return s.AdjVertice[v]
}

//showAdjVertices Displayes the vertices attached to the given edge
func (s *Graph) showAdjVertices(v int) {
	bag := s.AdjVertice[v]
	for k := range bag.Iterable() {
		log.Println(k)
	}
}

//Iterable returns the map of objects that are adjacent to the vertex
func (s *Graph) Iterable(v int) map[interface{}]int {
	return s.AdjVertice[v].Iterable()
}
