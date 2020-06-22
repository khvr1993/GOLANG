package threeWayQuickSort

import "github.com/khvr1993/GOLANG/utils"

//QuickSort does sorting using 3 way quick sort
func QuickSort(a *[]int, lo int, hi int) {
	//Implementing the 3 way partition
	//we need to have 3 pointers one gt, one lt,one i pointer and one pivot
	if hi <= lo {
		return
	}

	lt, gt, i := lo, hi, lo
	pivot := (*a)[lo]
	for i <= gt {
		if (*a)[i] < pivot {
			//eg: 2 1 3 5 pivot =2;lt:= 0 ; gt := 3;i=1
			//after swap 1 2 3 5 pivot =2;lt:= 1 ; gt := 3;i=2
			utils.Swap(a, lt, i)
			lt++
			i++
		} else if (*a)[i] > pivot {
			utils.Swap(a, i, gt)
			gt--
		} else {
			i++
		}
	}
	utils.PrintArray(a)
	QuickSort(a, lo, lt-1)
	QuickSort(a, gt+1, hi)
}
