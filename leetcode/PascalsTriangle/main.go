package main

import (
	"fmt"

	"github.com/khvr1993/GOLANG/leetcode/PascalsTriangle/generate"
)

/*
https://leetcode.com/problems/pascals-triangle/

Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

*/

func main() {
	numRows := 6
	outputLength := generate.Generate(numRows)
	fmt.Println(outputLength)
}
