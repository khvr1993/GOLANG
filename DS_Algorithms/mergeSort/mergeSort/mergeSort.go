package mergeSort

import (
	"github.com/khvr1993/GOLANG/DS_Algorithms/merge/merge"
)

//Sort sorts the array using merge Sort
func mergeSort(a []int, lo int, hi int) {
	mid := (hi + lo) / 2
	if lo == hi {
		return
	}

	if lo < hi {
		mergeSort(a, lo, mid)
		mergeSort(a, mid+1, hi)
		merge.Merge(&a, lo, hi, mid)
	}
}

//Sort implements mergeSort
func Sort(a []int) {
	mergeSort(a, 0, len(a)-1)
}
