package QuickUnion

import (
	"fmt"
	"log"
)

//UF will create a struct of slice
type UF struct {
	slice []int
}

//CreatePoints creates the Number of points passed
func (uf *UF) CreatePoints(n int) *UF {
	var i int
	for i < n {
		uf.slice = append(uf.slice, i)
		i++
	}
	return uf
}

//Root will return the root value of a given point
func (uf *UF) Root(point int) int {
	//First check whether the element is root of itself.
	//if not then set the point to the root and check if that point is the root(Moving up the tree)
	for point != uf.slice[point] {
		point = uf.slice[point]
	}
	return point
}

//Union will join 2 points in space
func (uf *UF) Union(p int, q int) {
	//First get the roots of both the points
	// Then assign A[p] = A[q]
	pRoot := uf.Root(p)
	qRoot := uf.Root(q)
	uf.slice[pRoot] = uf.slice[qRoot]
	return
}

//IsConnected will check if the given points are having same root
func (uf *UF) IsConnected(p int, q int) bool {
	return uf.Root(p) == uf.Root(q)
}

//PrintUF will print the array transformed
func (uf *UF) PrintUF() {
	var i int
	log.Println("--------------------------")
	for i < len(uf.slice) {
		fmt.Printf("%v ", uf.slice[i])
		i++
	}
	fmt.Printf("\n")
	log.Println("--------------------------")

}

//PathCreated returns the slice
func (uf *UF) PathCreated() []int {
	return uf.slice
}
