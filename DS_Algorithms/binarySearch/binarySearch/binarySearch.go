package binarySearch

import (
	"log"
	"sort"
)

//BinarySearch searches the given element
func BinarySearch(a []int, target int) (int, bool) {
	low := 0
	high := len(a) - 1
	sort.Ints(a)
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
			return middle, true
		}
	}
	return 0, false
}
