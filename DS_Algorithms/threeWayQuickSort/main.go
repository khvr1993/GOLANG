package main

import "github.com/khvr1993/GOLANG/DS_Algorithms/threeWayQuickSort/threeWayQuickSort"

func main() {
	a := []int{5, 3, 2, 1, 4, 1, 3, 1, 4, 1, 2, 1, 9, 8, 8, 7, 6}
	threeWayQuickSort.QuickSort(&a, 0, len(a)-1)
}
