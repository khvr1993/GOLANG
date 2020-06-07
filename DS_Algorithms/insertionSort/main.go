package main

import "github.com/khvr1993/GOLANG/DS_Algorithms/insertionSort/insertionSort"

func main() {
	a := []int{5, 3, 6, 8, 1}

	insertionSort.Sort(&a)
	insertionSort.PrintArray(&a)
}
