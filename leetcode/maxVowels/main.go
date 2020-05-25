package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/khvr1993/GOLANG/leetcode/maxVowels/maxVowels"
)

func main() {
	stringPassed := os.Args[1]
	var outputLength int
	fmt.Println(stringPassed)
	k, _ := strconv.Atoi(os.Args[2])
	fmt.Println(k)
	outputLength = maxVowels.MaxVowels(stringPassed, k)
	fmt.Println(outputLength)
}
