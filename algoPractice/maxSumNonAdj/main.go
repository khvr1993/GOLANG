package main

import (
	"log"

	"github.com/khvr1993/GOLANG/algoPractice/maxSumNonAdj/maxSumNonAdj"
)

func main() {
	a := []int{5, 1, 1, 0, 5}
	val := maxSumNonAdj.MaxSumNonAdj(a)
	log.Println("val => ", val)
}
