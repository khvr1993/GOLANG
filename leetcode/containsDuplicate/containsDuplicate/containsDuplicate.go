package containsDuplicate

func ContainsDuplicate(nums []int) bool {
	hashmap := make(map[int]int)

	for _, val := range nums {
		_, ok := hashmap[val]
		if ok {
			return true
		} else {
			hashmap[val] = 1
		}
		// fmt.Println(hashmap)
	}
	return false
}
