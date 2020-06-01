package quickUnion

import (
	"fmt"
	"log"
)

//UF will create a struct of slice
type UF struct {
	slice      []int
	size       []int
	actualList []int
}

//CreatePoints creates the Number of points passed
func (uf *UF) CreatePoints(n int) *UF {
	var i int
	for i < n {
		uf.slice = append(uf.slice, i)
		uf.size = append(uf.size, 1)
		uf.actualList = append(uf.actualList, i)
		i++
	}
	return uf
}

//Root will return the root value of a given point
func (uf *UF) Root(point int) int {
	//First check whether the element is root of itself.
	//if not then set the point to the root and check if that point is the root(Moving up the tree)
	for point != uf.slice[point] {
		uf.slice[point] = uf.slice[uf.slice[point]] //path Compression
		//The above step does this . it checked that the root node and the index are different.
		//Then we assign the value where the root node points to
		// eg indx -> 3 val -> 4 (A[3] == 4) =>  4 is the root node of 3
		// we replace A[3] to A[A[4]]. there by getting the root node of 4
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
	if pRoot == qRoot {
		return
	}
	if uf.size[pRoot] <= uf.size[qRoot] {
		uf.slice[pRoot] = qRoot //qRoot will become root node and pRoot will be attached as Child
		uf.size[qRoot] += uf.size[pRoot]
	} else {
		uf.slice[qRoot] = pRoot //pRoot will become root node and qRoot will become child
		uf.size[pRoot] += uf.size[qRoot]
		uf.actualList[pRoot] = qRoot
	}

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

//Size returns the size of the tree
func (uf *UF) Size(point int) int {
	return uf.size[point]
}

//ActualList returns the maximum value
func (uf *UF) ActualList(x int) int {
	return uf.actualList[uf.Root(x)]
}
