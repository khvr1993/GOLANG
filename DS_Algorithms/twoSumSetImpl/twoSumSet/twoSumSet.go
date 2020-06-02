package twoSumSet

import "log"

//TwoSum returns whehter the target is possible with 2 elements Sum
func TwoSum(a []int, target int) bool {
	var i int
	var twoSumMap = make(map[int]int)
	for i < len(a) {
		_, ok := twoSumMap[a[i]]
		if ok {
			twoSumMap[a[i]]++
		} else {
			twoSumMap[a[i]] = 1
		}
		i++
	}
	log.Println("The generated map is ", twoSumMap)

	for keys := range twoSumMap {
		supportingVal := target - keys
		log.Println("supporting Val ", supportingVal)

		_, ok := twoSumMap[supportingVal]
		if ok {
			if supportingVal == keys && twoSumMap[supportingVal] > 1 {
				return true
			} else if supportingVal == keys && twoSumMap[supportingVal] == 1 {
				return false
			}
			return true
		}
	}
	return false
}
