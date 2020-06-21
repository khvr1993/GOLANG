package main

import (
	"github.com/khvr1993/GOLANG/DS_Algorithms/PrintType"
	"github.com/khvr1993/GOLANG/DS_Algorithms/quickSort/quickSort"
)

func main() {
	a := []int{5, 13, 18, 1, 99, 39, 8, 12, 36, 36, 0, 99}
	quickSort.QuickSort(&a)
	//partition.Partition(&a, 0, len(a)-1)
	PrintType.PrintArray(&a)

}
