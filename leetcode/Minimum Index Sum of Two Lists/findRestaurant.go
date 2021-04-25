package findRestaurant

/*
https://leetcode.com/problems/minimum-index-sum-of-two-lists/
*/
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func findRestaurant(list1 []string, list2 []string) []string {

	//list1hash to store index in the value and key as the restaurant

	list1_hashMap := make(map[string]int)
	minimum_index_val := 9999999
	common_interests := make(map[string]int)
	var op []string

	for index, vals := range list1 {
		list1_hashMap[vals] = index
	}

	for index, vals := range list2 {
		list1_index, ok := list1_hashMap[vals]
		if ok {
			minimum_index_val = min(minimum_index_val, (index + list1_index))
			common_interests[vals] = index + list1_index
		}
	}

	for key, value := range common_interests {
		if value == minimum_index_val {
			op = append(op, key)
		}
	}

	return op

}
