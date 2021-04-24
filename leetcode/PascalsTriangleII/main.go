package main

import (
	"fmt"

	"github.com/khvr1993/GOLANG/leetcode/PascalsTriangleII/getRow"
)

/*
https://leetcode.com/problems/pascals-triangle-ii/

Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

*/

func main() {
	numRows := 3
	outputLength := getRow.GetRow(numRows)
	fmt.Println(outputLength)
}
