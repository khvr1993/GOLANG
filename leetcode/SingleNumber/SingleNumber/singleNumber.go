package singleNumber

func SingleNumber(nums []int) int {
	hashtable := make(map[int]int)

	// Populate map with the count
	for _, val := range nums {
		_, ok := hashtable[val]
		if ok {
			hashtable[val] += 1
		} else {
			hashtable[val] = 1
		}
	}

	//Iterate through the array and return where the count = 1

	for key, value := range hashtable {
		if value == 1 {
			return key
		}
	}

	return 0
}
