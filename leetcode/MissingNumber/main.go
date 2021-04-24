package main

import (
	"fmt"

	"github.com/khvr1993/GOLANG/leetcode/MissingNumber/missingNumber"
)

/*
https://leetcode.com/problems/missing-number/

Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

Follow up: Could you implement a solution using only O(1) extra space complexity and O(n) runtime complexity?

Sum of first n integers is n*(n+1)/2
*/

func main() {
	inp := []int{0}
	Op := missingNumber.MissingNumber(inp)
	fmt.Println(Op)
}
