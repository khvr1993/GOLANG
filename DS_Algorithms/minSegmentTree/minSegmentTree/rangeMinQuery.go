package minSegmentTree

import "github.com/khvr1993/GOLANG/utils"

func RangeMinQuery(segTree *[]int, qLow int, qHigh int, lo int, hi int, pos int) int {
	/*
		We have 3 criteria
			1. Total Overlap -> Return the value at position
			2. Partial Overlap -> Check both left and right [2*pos+1 and 2*pos+2]
			3. No Overlap -> return maxvalue
	*/
	if qLow <= lo && qHigh >= hi {
		return (*segTree)[pos]
	} else if qLow > hi || qHigh < lo {
		return MAXINT
	}
	mid := (lo + hi) / 2

	return utils.Min(RangeMinQuery(segTree, qLow, qHigh, lo, mid, 2*pos+1), RangeMinQuery(segTree, qLow, qHigh, mid+1, hi, 2*pos+2))
}
