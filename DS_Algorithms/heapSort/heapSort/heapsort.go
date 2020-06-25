package heapSort

import (
	"github.com/khvr1993/GOLANG/DS_Algorithms/PrintType"
	"github.com/khvr1993/GOLANG/utils"
)

//HeapSort sorts the array using heap sort
func HeapSort(a *[]int) {
	var k int
	arrSize := len(*a)
	for k = arrSize / 2; k >= 1; k-- {
		sink(a, k, arrSize)
	}
	PrintType.PrintArray(a)

	for arrSize > 1 {
		utils.Swap(a, 0, arrSize-1)
		arrSize--
		sink(a, 1, arrSize)
	}
}

func sink(a *[]int, k int, N int) {
	var j, swapElem int
	for 2*k <= N {
		j = 2 * k
		//to maintain 0 based array
		if j < N && (*a)[2*k-1] < (*a)[2*k] {
			j = 2*k + 1
			swapElem = 2 * k
		} else {
			j = 2 * k
			swapElem = 2*k - 1
		}
		if (*a)[swapElem] < (*a)[k-1] {
			break
		}
		// to  maintain 0 based array
		utils.Swap(a, k-1, swapElem)
		k = j
	}
}
