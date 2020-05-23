package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stringPassed := os.Args[1:]
	noOfValToChange, _ := strconv.Atoi(os.Args[2])
	var output []int
	var res2 []int
	fmt.Println(stringPassed)
	res1 := strings.Split(stringPassed[0], ",")
	fmt.Println(res1)
	fmt.Println(noOfValToChange)
	for i := 0; i < len(res1); i++ {
		val, _ := strconv.Atoi(res1[i])
		res2 = append(res2, val)
	}
	output = maximumSubArray(res2, noOfValToChange)
	fmt.Println(output)
}

func maximumSubArray(stdIn []int, K int) []int {
	var windowEnd int
	var windowStart int
	var maxSum int
	var sum int
	var maxArray []int

	for windowEnd < len(stdIn) {
		if K > 0 {
			K--
			sum += stdIn[windowEnd]
			windowEnd++
			fmt.Println("Current value of K ", K, "and windowEnd is  ", windowEnd, " and sum calc is ", sum)
			if sum > maxSum {
				maxArray = nil
				for i := windowStart; i < windowEnd; i++ {
					maxArray = append(maxArray, stdIn[i])
				}
				maxSum = max(sum, maxSum)
				//push the Array

			}
		} else {
			K++
			sum -= stdIn[windowStart]
			windowStart++

		}
	}
	fmt.Println("Max Sum upto this Point ", maxSum, " and max Array is", maxArray)

	return maxArray

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
