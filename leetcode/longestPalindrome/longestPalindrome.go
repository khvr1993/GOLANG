package main

import (
	"log"
	"os"
)

func main() {
	stringPassed := os.Args[1:]
	var opPalindrome string
	log.Println(stringPassed)
	opPalindrome = longestPalindrome(stringPassed[0])
	log.Println("opPalindrome => ", opPalindrome)
}

func longestPalindrome(s string) string {
	var maxPalindrome string
	//set window size to maximumlength -1
	windowSize := len(s)
	runes := []rune(s)
	//log.Println("WindowSize ", windowSize)
	for windowSize > 0 {
		var windowStart int
		var windowEnd int
		windowEnd = windowSize
		for windowEnd <= len(s) {
			//log.Println("The start and end windows are ", windowStart, "", windowEnd)
			if s[windowStart] == s[windowEnd-1] {
				palindromicTestString := string(runes[windowStart:windowEnd])
				//log.Println("palindromicTestString => ", palindromicTestString)
				if checkPalindrome(palindromicTestString, windowSize) {
					log.Println("The passed string is Palindrome => ", palindromicTestString)
					return string(runes[windowStart:windowEnd])
				}
			}
			windowStart++
			windowEnd++
		}
		log.Println("Decrement. the window size is ", windowSize)
		windowSize--
	}

	return maxPalindrome
}

func checkPalindrome(s string, strLen int) bool {
	var i int
	if strLen%2 == 0 {
		n := strLen / 2
		for i < n {
			if s[n-i-1] != s[n+i] {
				return false
			}
			i++
		}
		return true
	}
	n := strLen / 2
	for i < n {
		if s[n-i-1] != s[n+i+1] {
			return false
		}
		i++
	}
	return true

}
