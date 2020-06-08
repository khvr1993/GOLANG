package shellSort

import (
	"fmt"
)

//Sort sorts the array using shell Sort
func Sort(a *[]int) {
	var h, i, j int
	length := len(*a)
	h = 1
	// get the maximum value of h
	for length > h {
		h = h*3 + 1
	} //it gives values 1,4,13,40...

	for h >= 1 {
		//Perform Insertion Sort

		for i = h; i < length; i = i + h {
			keyVal := (*a)[i]
			j = i - h
			for j >= 0 && (*a)[j] > keyVal {
				(*a)[j+h] = (*a)[j]
				j = j - h
			}
			(*a)[j+h] = keyVal
		}
		h = h / 3
	}
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
