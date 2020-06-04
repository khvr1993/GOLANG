package findMaxElemBitonicArray

import "log"

//FindMax finds the maximum element in a bitonic Array
func FindMax(a []int) int {

	length := len(a)
	low, high := 0, length-1
	//handling cases where the array length is just 2
	if length < 3 {
		return -1
	}

	for low <= high {
		middle := (low + high) / 2
		log.Println("low => ", low, " high => ", high, " middle => ", middle)

		if a[middle] > a[middle+1] && a[middle] > a[middle-1] {
			//bitonic Array Max
			return middle
		} else if a[middle] < a[middle+1] {
			//Undershot the middle
			low = middle
		} else if a[middle] < a[middle-1] {
			//Overshot the middle
			high = middle
		}

	}

	return -1
}
