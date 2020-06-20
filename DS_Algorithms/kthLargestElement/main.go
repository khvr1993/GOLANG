package main

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/kthLargestElement/kthLargestEleemnt"
)

func main() {
	a := []int{5, 13, 18, 1, 4, 39, 8, 12, 25, 36, 0, 99}
	val := kthLargestEleemnt.KthLargestElem(&a, 6)
	log.Println("val => ", val)
}
