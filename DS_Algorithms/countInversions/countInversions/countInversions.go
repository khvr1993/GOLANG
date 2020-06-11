package countInversions

import "github.com/khvr1993/GOLANG/DS_Algorithms/countInversions/countInvMerge"

//CountInversions counts the number of inversions in an Array
func CountInversions(a *[]int) int {
	var length, arrayStart, invCount int
	aux := make([]int, len(*a))
	for length = 1; length < len(*a); length = 2 * length {
		//Length increments in intervals 1,2,4,8,16...
		for arrayStart = 0; arrayStart < len(*a)-length; arrayStart += 2 * length {
			//When length is 1 arrays are (0),(1),(2)
			//For merge we pass (0,1,0) middle element is 0
			low := arrayStart
			middle := low + length - 1
			high := low + 2*length - 1

			if high > len(*a)-1 {
				high = len(*a) - 1
			}

			invCount += countInvMerge.CountInvMerge(a, &aux, low, high, middle)
		}
	}
	return invCount
}
