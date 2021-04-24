package missingNumber

func MissingNumber(nums []int) int {
	n := len(nums)
	var arr_sum int
	sumOfintegers := n * (n + 1) / 2

	//Calculate the sum by iterating through the array

	for _, val := range nums {
		arr_sum += val
	}

	if (sumOfintegers - arr_sum) == 0 {
		return 0
	} else {
		return (sumOfintegers - arr_sum)
	}

}
