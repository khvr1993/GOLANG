package main

import "github.com/khvr1993/GOLANG/DS_Algorithms/shellSort/shellSort"

func main() {
	a := []int{8, 5, 4, 7, 2, 1}
	shellSort.Sort(&a)
	shellSort.PrintArray(&a)
}
