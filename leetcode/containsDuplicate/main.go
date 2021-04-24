package main

import (
	"fmt"

	"github.com/khvr1993/GOLANG/leetcode/containsDuplicate/containsDuplicate"
)

/*
https://leetcode.com/problems/contains-duplicate/

Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.
*/

func main() {
	inp := []int{1, 2, 3, 4, 1}
	Op := containsDuplicate.ContainsDuplicate(inp)
	fmt.Println(Op)
}
