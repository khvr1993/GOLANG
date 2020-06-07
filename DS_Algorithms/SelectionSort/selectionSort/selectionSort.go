package selectionSort

import "fmt"

//IntArr is the array that needs to be sorted
type IntArr struct {
	Data []int
}

type stringArr struct {
	Data []string
}

//Sort Implements the selection sort
func (A *IntArr) Sort() {
	var i, j int
	len := len(A.Data)
	for i = 0; i < len; i++ {
		min := i
		for j = i + 1; j < len; j++ {
			if A.compareIt(j, min) == 1 {
				min = j
			}
		}
		A.exchange(i, min)
	}
}

//Sort Implements the selection sort
func (A *stringArr) Sort() {
	var i, j int
	len := len(A.Data)
	for i = 0; i < len; i++ {
		min := i
		for j = i + 1; j < len; j++ {
			if A.compareIt(j, min) == 1 {
				min = j
			}
		}
		A.exchange(i, min)
	}
}

func (A *IntArr) compareIt(i, j int) int {
	if A.Data[i] < A.Data[j] {
		return 1
	} else if A.Data[i] == A.Data[i] {
		return 0
	}
	return -1
}

func (A *stringArr) compareIt(i, j int) int {
	if A.Data[i] < A.Data[j] {
		return 1
	} else if A.Data[i] == A.Data[i] {
		return 0
	}
	return -1
}

//exchange swaps the elements
func (A *IntArr) exchange(i int, min int) {
	tempVal := A.Data[i]
	A.Data[i] = A.Data[min]
	A.Data[min] = tempVal
}

//exchange swaps the elements
func (A *stringArr) exchange(i int, min int) {
	tempVal := A.Data[i]
	A.Data[i] = A.Data[min]
	A.Data[min] = tempVal
}

//PrintArray prints the array
func (A *IntArr) PrintArray() {
	var i int
	for i < len(A.Data) {
		fmt.Printf("%v\t", A.Data[i])
		i++
	}
	fmt.Printf("\n")
}

//PopData populates Data
func (A *IntArr) PopData(a []int) {
	var i int
	lengthA := len(A.Data)
	for i < len(a) {
		A.Data[lengthA] = a[i]
		i++
		lengthA++
	}
}
