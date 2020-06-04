package searchBitonicArray

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/binarySearch/binarySearch"

	"github.com/khvr1993/GOLANG/DS_Algorithms/findMaxElemBitonicArray/findMaxElemBitonicArray"
)

//FindTarget searches in bitonic Array and returns the index of
//the target
func FindTarget(a []int, target int) bool {
	maxElementIdx := findMaxElemBitonicArray.FindMax(a)

	log.Println("The maximum element in Array is ", maxElementIdx)

	if maxElementIdx == -1 {
		return false
	}

	ok := binarySearch.BinarySearch(a[0:maxElementIdx+1], target)
	if ok {
		return true
	}

	ok = binarySearch.BinarySearch(a[maxElementIdx+1:], target)
	if ok {
		return true
	}
	return false
}
