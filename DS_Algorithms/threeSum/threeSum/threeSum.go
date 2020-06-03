package threeSum

import (
	"sort"

	"github.com/khvr1993/GOLANG/DS_Algorithms/twoSumSetImpl/twoSumSet"
)

//ThreeSum checks if the target can be acheived by adding any three elements
func ThreeSum(a []int, target int) bool {
	var i int
	sort.Ints(a)
	for i < len(a)-2 {
		newTarget := target - a[i]
		ok := twoSumSet.TwoSum(a[i+1:], newTarget)
		if ok {
			return true
		}
		i++
	}
	return false
}
