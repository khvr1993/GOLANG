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
	var output int
	var res2 []int
	fmt.Println(stringPassed)
	res1 := strings.Split(stringPassed[0], ",")
	fmt.Println(res1)
	fmt.Println(noOfValToChange)
	for i := 0; i < len(res1); i++ {
		val, _ := strconv.Atoi(res1[i])
		res2 = append(res2, val)
	}
	output = longestOnes(res2, noOfValToChange)
	fmt.Println(output)
}

func longestOnes(A []int, K int) int {
	var windowStart, windowEnd int
	for windowEnd < len(A) {
		if A[windowEnd] == 0 {
			K--
		}
		if K < 0 {
			if A[windowStart] == 0 {
				K++
			}
			windowStart++
		}
		windowEnd++
	}
	return windowEnd - windowStart
}
