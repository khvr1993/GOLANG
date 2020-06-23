package maxSumNonAdj

import (
	"log"

	"github.com/khvr1993/GOLANG/utils"
)

//MaxSumNonAdj returns the maximum sum of non adjacent numbers
func MaxSumNonAdj(a []int) int {
	var maxValue int
	incl := a[0]
	excl := 0
	i := 1
	/* When you include the first val in array then the exclusions is 0.
	include the maximum value of incl and excl.
	When we move to the second val to include this we need to add it to the exclusion
	of the previous result. again return the maximum
	*/

	for i < len(a) {
		log.Println("incl => ", incl, "excl => ", excl, "a[i] => ", a[i])
		maxValue = utils.Max(incl, excl)
		incl = excl + a[i]
		excl = maxValue

		i++
	}

	return utils.Max(incl, excl)
}
