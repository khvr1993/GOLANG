package main

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/countInversions/countInversions"
)

func main() {
	a := []int{4, 2, 1, 5}
	invCount := countInversions.CountInversions(&a)
	log.Println("invCount => ", invCount)
}
