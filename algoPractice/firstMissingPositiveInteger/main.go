package main

import (
	"log"

	"github.com/khvr1993/GOLANG/algoPractice/firstMissingPositiveInteger/firstMissingPosIntHash"
)

func main() {
	a := []int{5, 7, 0, 1, 2, 3, 4, 5, 6, -8, 6}
	missingPosInt := firstMissingPosIntHash.FirstMissingPosIntHash(&a)
	log.Println("missingPosInt ", missingPosInt)
}
