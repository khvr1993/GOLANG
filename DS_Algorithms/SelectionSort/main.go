package main

import "github.com/khvr1993/GOLANG/DS_Algorithms/SelectionSort/selectionSort"

func main() {
	A := selectionSort.IntArr{Data: []int{6, 8, 4, 5}}
	A.Sort()
	A.PrintArray()
}
