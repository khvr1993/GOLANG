package twoSum

import (
	"sort"
)

//TwoSum returns true if the target can be acheived by summing 2 nums
func TwoSum(a []int, target int) bool {
	sort.Ints(a)
	var start int
	end := len(a) - 1
	for start != end {
		if a[start]+a[end] == target {
			return true
		} else if a[start]+a[end] > target {
			end--
		} else if a[start]+a[end] < target {
			start++
		}
	}
	return false
}
