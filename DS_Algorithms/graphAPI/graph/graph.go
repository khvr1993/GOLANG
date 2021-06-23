package graph

import (
	"github.com/khvr1993/GOLANG/DS_Algorithms/bag/Bag"
	"log"
)

/*
	Create a bags for vertices
*/
type adjStruct struct {
	AdjVertice []Bag.Bag
}

/*
	Constructor to create the adj Vertice
*/
func NewGraph(V int) *adjStruct {
	log.Println("Begin NewGraph")
	s := &adjStruct{}
	s.AdjVertice = make([]Bag.Bag, 13)
	s.Graph(V)
	return s
}

/*
	for a given number of vertices the func initializes the bags and assigns them
*/
func (s *adjStruct) Graph(V int) {

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
func (s *adjStruct) AddEdge(v int, w int) {
	s.AdjVertice[v].AddItem(w)
	s.AdjVertice[w].AddItem(v)
}

/*
	Returns the bag to iterate over
*/
func (s *adjStruct) Iterable(v int) Bag.Bag {
	return s.AdjVertice[v]
}

/*
	Displayes the vertices attached to the given edge
*/
func (s *adjStruct) showAdjVertices(v int) {
	bag := s.AdjVertice[v]
	for k := range bag.Iterable() {
		log.Println(k)
	}
}
