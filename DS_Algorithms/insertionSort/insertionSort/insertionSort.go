package insertionSort

import (
	"fmt"
)

//Sort sorts the elements using insertion sort
func Sort(a *[]int) {
	var i, j int
	i = 1
	len := len(*a)
	for i < len {
		j = i - 1
		key := (*a)[i]

		for j >= 0 && (*a)[j] > key {
			swap(a, j)
			j--
		}
		(*a)[j+1] = key //replace the value in the array where the iteration stopped
		i++
	}
}

func swap(a *[]int, j int) {
	(*a)[j+1] = (*a)[j]
}

//PrintArray Prints the values
func PrintArray(a *[]int) {
	var i int
	for i < len(*a) {
		fmt.Printf("%v\t", (*a)[i])
		i++
	}
	fmt.Printf("\n")
}
