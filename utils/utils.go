package utils

import (
	"fmt"
)

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//PrintArray prints the array
func PrintArray(a *[]int) {
	var i int
	for i < len(*a) {
		fmt.Printf("%v\t", (*a)[i])
		i++
	}
	fmt.Printf("\n")
}

//PrintArr prints the array
func PrintArr(a []interface{}) {
	var i int
	for i < len(a) {
		fmt.Printf("%v\t", a[i])
		i++
	}
	fmt.Printf("\n")
}

//PrintString prints the array
func PrintString(a *[]string) {
	var i int
	for i < len(*a) {
		fmt.Printf("%v\t", (*a)[i])
		i++
	}
	fmt.Printf("\n")
}

//Max returns the max of the 2 values
func Max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

//Min returns the min of 2 values
func Min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

//Swap the ith element to the jth element position
func Swap(a *[]int, i int, j int) {
	temp := (*a)[j]
	(*a)[j] = (*a)[i]
	(*a)[i] = temp
}
