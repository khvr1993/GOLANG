package main

import (
	"log"
	"time"

	"github.com/khvr1993/GOLANG/DS_Algorithms/mergeSort/mergeSort"
)

func main() {
	start := time.Now()
	a := []int{183,
		876}
	mergeSort.Sort(a)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	//PrintType.PrintArray(&a)

}
