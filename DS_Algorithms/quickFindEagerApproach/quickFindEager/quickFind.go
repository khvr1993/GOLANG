package quickFindEager

import (
	"fmt"
	"log"
)

//UF Creates the struct of slice wherein we can store the
//points
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

//PrintPathTransform prints the slice modified after getting unions
func (uf *UF) PrintPathTransform() {
	var i int
	log.Println("--------------------------")
	for i < len(uf.slice) {
		fmt.Printf("indx => %v  value => %v\n", i, uf.slice[i])
		i++
	}
	log.Println("--------------------------")
	return
}

//IsConnected will check if there is a path defined between two paths
func (uf *UF) IsConnected(p int, q int) bool {
	return uf.slice[p] == uf.slice[q]
}

//Union Creates a path between 2 points
func (uf *UF) Union(p int, q int) {
	var i int
	pid := uf.slice[p]
	qid := uf.slice[q]

	for i < len(uf.slice) {
		if uf.slice[i] == pid {
			uf.slice[i] = qid
		}
		i++
	}
	return
}

//PathCreated returns the path created
func (uf *UF) PathCreated() []int {
	return uf.slice
}
