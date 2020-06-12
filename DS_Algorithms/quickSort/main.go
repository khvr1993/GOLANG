package main

import "github.com/khvr1993/GOLANG/DS_Algorithms/quickSort/partition"

func main() {
	a := []int{5, 11, 12, 1, 2, 6, 8, 10, 3, 4}
	partition.Partition(&a, 0, len(a)-1)
}
