package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	stringPassed := os.Args[1:]
	var outputLength int
	log.Println(stringPassed)
	res1 := strings.Split(stringPassed[0], "")
	log.Println(res1)

	outputLength = lengthOfLongestSubstring(stringPassed[0])
	log.Println(outputLength)
}
func lengthOfLongestSubstring(s string) int {
	var output int

	// var windowStart int
	//var iterator int
	//Convert string to array
	//var windowSize int
	var i, j int32
	length := int32(len(s))
	var mapArray = make(map[byte]int32)
	for i < length {
		maxLength := 0
		log.Println("Current value of i => ", i)
		for j < length {
			log.Println("Current value of j => ", j, string(s[j]))
			_, exist := mapArray[s[j]]
			log.Println("The value ", string(s[j]), " in map ", mapArray, " is ", exist)
			if !exist {
				mapArray[s[j]] = j
				maxLength = len(mapArray)
				output = max(output, maxLength)
				j++
			} else {
				delete(mapArray, s[i])
				log.Println("Deleting value of i => and output is ", string(s[i]), output)
				break
			}
		}
		/* j = i
			maxLength := 0
			for _, exist := mapArray[string(s[j])]; !exist && j < len(s); {
				fmt.Println("Adding to ARRAY for j ", j, string(s[j]))
				mapArray[string(s[j])] = j
				maxLength++
				output = max(output, maxLength)
				j++
			}
			fmt.Println(mapArray, output)
			i++
		} */
		i++
	}
	return output

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
