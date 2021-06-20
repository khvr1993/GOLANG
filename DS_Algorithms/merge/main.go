package main

import "github.com/khvr1993/GOLANG/DS_Algorithms/merge/merge"

func main() {
	a := []int{7, 11, 10, 1, 2, 3}
	var aux []int
	merge.Merge(&a, &aux, 0, 5, 2)
}
