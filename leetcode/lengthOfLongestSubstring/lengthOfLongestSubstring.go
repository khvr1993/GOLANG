package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	stringPassed := os.Args[1:]
	var outputLength int
	fmt.Println(stringPassed)
	res1 := strings.Split(stringPassed[0], "")
	fmt.Println(res1)

	outputLength = lengthOfLongestSubstring(stringPassed[0])
	fmt.Println(outputLength)
}
func lengthOfLongestSubstring(s string) int {
	var output int
	//var iterator int
	//Convert string to array
	res1 := strings.Split(s, "")
	arrayLen := len(res1)

	for indx := 0; indx < arrayLen; indx++ {
		fmt.Printf("Working on value %d\n", indx)
		var elementsMap = map[string]int{}
		var maxLength int
		for j := indx; j < arrayLen; j++ {
			_, exist := elementsMap[res1[j]]
			fmt.Println(elementsMap)
			fmt.Println(exist)
			if !exist {
				maxLength = maxLength + 1
				if output < maxLength {
					output = maxLength
				}
			} else {
				break
			}
			elementsMap[res1[j]] = j
		}
		if maxLength == arrayLen {
			return maxLength
		}
		if maxLength > arrayLen-indx {
			fmt.Printf("Current Output is %d at Indx %d \n", output, indx)
			return output

		}
		fmt.Printf("max Length after indx %d is %d and output is %d\n", indx, maxLength, output)
	}
	return output
}
