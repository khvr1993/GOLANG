package PrintType

import "fmt"

//PrintArray prints the array
func PrintArray(a *[]int) {
	var i int
	for i < len((*a)) {
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
	for i < len((*a)) {
		fmt.Printf("%v\t", (*a)[i])
		i++
	}
	fmt.Printf("\n")
}
