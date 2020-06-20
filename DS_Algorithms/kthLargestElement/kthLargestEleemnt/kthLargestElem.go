package kthLargestEleemnt

import (
	"github.com/khvr1993/GOLANG/DS_Algorithms/quickSort/partition"
)

//KthLargestElem returns the Kth largest element in Array
func KthLargestElem(a *[]int, k int) int {
	var lo int
	hi := len(*a) - 1
	for hi > lo {
		j := partition.Partition(a, lo, hi)
		if j > k {
			lo = 0
			hi = j - 1
		} else if j < k {
			lo = j + 1
		} else {
			return (*a)[k]
		}
	}
	return (*a)[k]
}
