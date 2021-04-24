package quickSort

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/quickSort/partition"
)

//QuickSort implements quickSort
func QuickSort(a *[]int) {
	log.Println("QuickSort")
	// shuffle.Shuffle(a)
	sort(a, 0, len(*a)-1)
}

func sort(a *[]int, lo int, high int) {
	log.Println("sort")
	if high <= lo {
		return
	}
	j := partition.Partition(a, lo, high)
	log.Println("the in position element is ", j)
	sort(a, lo, j-1)
	sort(a, j+1, high)
}
