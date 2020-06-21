package main

import (
	"log"

	"github.com/khvr1993/GOLANG/algoPractice/firstMissingPositiveInteger/firstMissingPosInt"
)

func main() {
	a := []int{5, 7, 1, -8, 6}
	missingPosInt := firstMissingPosInt.FirstMissingPosInt(&a)
	log.Println("missingPosInt ", missingPosInt)
}
