package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	stringPassed := os.Args[1:]
	// noOfValToChange, _ := strconv.Atoi(os.Args[2])
	var output []int
	var res2 []int
	fmt.Println(stringPassed)
	res1 := strings.Split(stringPassed[0], ",")
	fmt.Println(res1)
	// fmt.Println(noOfValToChange)
	for i := 0; i < len(res1); i++ {
		val, _ := strconv.Atoi(res1[i])
		res2 = append(res2, val)
	}
	output = numMovesStonesII(res2)
	fmt.Println(output)
}
func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	fmt.Println(stones)
	var windowStart int
	var windowEnd int
	var windowLength int
	var maxWindowLength int
	var minMax []int
	lengthofArray := len(stones)
	for windowEnd < lengthofArray {
		if stones[windowEnd]-stones[windowStart] == windowEnd-windowStart {
			windowLength = windowEnd - windowStart + 1
			maxWindowLength = max(maxWindowLength, windowLength)
			fmt.Println("The window Length is ", windowLength, " for indx ", windowEnd, " and window Start ", windowStart)
			windowEnd++

		} else {
			windowStart++
		}
	}
	fmt.Println("The Maximum Window sorted is ", maxWindowLength)
	return minMax
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
