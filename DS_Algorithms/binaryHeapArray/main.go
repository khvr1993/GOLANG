package main

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/binaryHeapArray/binaryHeap"
)

func main() {
	maxPQ := new(binaryHeap.BinaryHeap)
	maxPQ.Insert(6)
	maxPQ.Insert(7)
	maxPQ.Insert(10)
	maxPQ.Insert(12)
	maxPQ.Insert(17)
	maxPQ.Insert(5)
	maxPQ.ShowHeap()
	maxelem := maxPQ.GetMax()
	log.Println("maxelem => ", maxelem)
	maxPQ.ShowHeap()
	maxelem = maxPQ.GetMax()
	log.Println("maxelem => ", maxelem)
	maxPQ.ShowHeap()
	maxelem = maxPQ.GetMax()
	log.Println("maxelem => ", maxelem)
	maxPQ.ShowHeap()
}
