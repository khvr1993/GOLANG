package mergeSort

import (
	"github.com/khvr1993/GOLANG/DS_Algorithms/merge/merge"
)

//Sort sorts the array using merge Sort
func mergeSort(a *[]int, aux *[]int, lo int, hi int) {
	mid := (hi + lo) / 2
	if lo == hi {
		return
	}

	if lo < hi {
		mergeSort(a, aux, lo, mid)
		mergeSort(a, aux, mid+1, hi)
		merge.Merge(a, aux, lo, hi, mid)
	}
}

//Sort implements mergeSort
func Sort(a []int) {
	aux := make([]int, len(a))
	mergeSort(&a, &aux, 0, len(a)-1)
}
