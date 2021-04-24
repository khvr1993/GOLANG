package main

import (
	"fmt"

	singleNumber "github.com/khvr1993/GOLANG/leetcode/SingleNumber/SingleNumber"
)

/*

https://leetcode.com/problems/single-number/

Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

Follow up: Could you implement a solution with a linear runtime complexity and without using extra memory?
This can be done using XOR Operation
A XOR 0 == A
A XOR A == 0


*/

func main() {
	numRows := []int{1, 2, 2, 1, 4}
	Op := singleNumber.SingleNumber(numRows)
	fmt.Println(Op)
}
