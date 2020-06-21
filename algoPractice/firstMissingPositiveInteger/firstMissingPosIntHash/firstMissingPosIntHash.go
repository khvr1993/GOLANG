package firstMissingPosIntHash

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/PrintType"
	"github.com/khvr1993/GOLANG/utils"
)

var containsOne bool

func removenoImpactElem(a *[]int) int {
	var i, j int
	for i < len(*a) {
		if (*a)[i] == 1 {
			containsOne = true
		}
		if (*a)[i] > len(*a) || (*a)[i] <= 0 {
			utils.Swap(a, i, j)
			j++
		}
		i++
	}
	PrintType.PrintArray(a)
	return j
}

//FirstMissingPosIntHash uses hashmap to find the first missing number
func FirstMissingPosIntHash(a *[]int) int {
	i := 1
	j := removenoImpactElem(a)
	if !containsOne {
		return 1
	}
	//convert slice to map
	elementMap := utils.ConvertsliceToMap((*a)[j:])
	log.Println("elementMap ", elementMap)
	for i < len(*a) {
		_, ok := elementMap[i]
		if !ok {
			return i
		}
		i++
	}
	return 0
}
