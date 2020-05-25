package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/khvr1993/GOLANG/leetcode/maxVowels/maxVowels"
	"github.com/khvr1993/leetcode/greet"
)

func main() {
	stringPassed := os.Args[1:]
	var outputLength int
	fmt.Println(stringPassed)
	k, _ := strconv.Atoi(os.Args[2])
	fmt.Println(k)
	greet.Hello()
	outputLength = maxVowels.MaxVowels(stringPassed[0], k)
	fmt.Println(outputLength)
}
