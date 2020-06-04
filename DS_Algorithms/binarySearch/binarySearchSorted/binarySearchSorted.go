package binarySearchSorted

import (
	"log"
)

//BinarySearch searches the given element and returns the index
func BinarySearch(a []int, target int) int {
	low := 0
	high := len(a) - 1
	for low <= high {
		log.Println("high => ", high, "low => ", low)
		middle := (high + low) / 2
		log.Println("middle Element ", middle)
		if a[middle] > target {
			high = middle - 1
			middle = (high + low) / 2
		} else if a[middle] < target {
			low = middle + 1
			middle = (high + low) / 2
		} else {
			return middle
		}
	}
	return -1
}
