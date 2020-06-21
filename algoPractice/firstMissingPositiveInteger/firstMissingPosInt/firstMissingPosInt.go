package firstMissingPosInt

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/PrintType"
	"github.com/khvr1993/GOLANG/utils"
)

var containsOne bool

/* All the elements which have no impact : negative elements and the elements
which are greater than the val  */
func removenoImpactElem(a *[]int) int {
	var i, j int
	for i < len(*a) {
		if (*a)[i] == 1 {
			containsOne = true
		}

		if (*a)[i] > len(*a) || (*a)[i] <= 0 {
			(*a)[i] = 1
		}
		i++
	}
	PrintType.PrintArray(a)
	return j
}

//FirstMissingPosInt returns the first positive integer missing in an Array
func FirstMissingPosInt(a *[]int) int {
	var i, j int
	removenoImpactElem(a)
	for i < len(*a) {
		k := utils.Abs((*a)[i]) - 1
		log.Println("Value of k", k)
		if (*a)[k] > 0 {
			(*a)[k] = -1 * (*a)[k]
		}
		i++
	}
	PrintType.PrintArray(a)
	if !(containsOne) {
		return 1
	}

	for j < len(*a) {
		if (*a)[j] > 0 {
			return j + 1
		}
		j++
	}
	PrintType.PrintArray(a)
	return j + 1
}
