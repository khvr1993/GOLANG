package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arrayString := os.Args[1:]
	target, _ := strconv.Atoi(os.Args[2])

	convArray := strings.Split(arrayString[0], ",")

	//convert to int array
	var nums []int
	for _, vals := range convArray {
		i1, _ := strconv.Atoi(vals)
		nums = append(nums, i1)
	}

	fmt.Println(target)
	fmt.Println(nums)

	var twoSumcomb []int
	//fmt.Println(convArray)
	twoSumcomb = twoSum(nums, target)
	fmt.Println(twoSumcomb)
	/*
		for _, values := range convArray {
			fmt.Println(values)
		}
	*/
}

/*
func twoSum(nums []int, target int) []int {
	var returnArray []int
	combFound := false
	for i := 0; i < len(nums)-1; i++ {
		//fmt.Printf("Working on indx i %d value%d \n", i, nums[i])
		for j := i + 1; j < len(nums); j++ {
			//fmt.Printf("Working on indx j %d value%d \n", j, nums[j])
			//fmt.Printf("Sum %d\n", nums[i]+nums[j])
			if nums[i]+nums[j] == target {
				returnArray = append(returnArray, i)
				returnArray = append(returnArray, j)
				combFound = true
				break
			}
		}
		if combFound {
			break
		}
	}
	return returnArray
} */

func twoSum(nums []int, target int) []int {
	var seen = make(map[int]int)
	for i := range nums {
		fmt.Printf("The index is %d and value %d\n", i, nums[i])
		fmt.Println(seen)
		item, exist := seen[target-nums[i]]
		fmt.Println(item)
		fmt.Println(exist)
		if exist {
			fmt.Println(seen)
			return []int{i, item}
		}
		seen[nums[i]] = i
		fmt.Println(seen)
	}
	return []int{0, 0}
}
