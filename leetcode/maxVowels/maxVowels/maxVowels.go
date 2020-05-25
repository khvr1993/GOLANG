package maxVowels

import "log"

//MaxVowels For a given window returns the maximum count of vowels
func MaxVowels(s string, k int) int {
	var maxCount int
	var windowStart, windowEnd int
	var vowelMap = make(map[string]int)
	vowelMap["a"] = 0
	vowelMap["e"] = 0
	vowelMap["i"] = 0
	vowelMap["o"] = 0
	vowelMap["u"] = 0
	log.Println("len(s)", len(s))
	for windowEnd < len(s) {
		log.Println("windowEnd ", windowEnd)
		if k > 0 {
			log.Println("The window is at positions ", windowStart, ",", windowEnd, " string(s[windowEnd]) is ", string(s[windowEnd]))
			k--
			_, ok := vowelMap[string(s[windowEnd])]
			if ok {
				vowelMap[string(s[windowEnd])]++
			}
			maxCount = max(maxCount, returnMapSum(vowelMap))
			windowEnd++
		} else {
			log.Println("The Current mapValues are ", vowelMap)
			k++
			_, ok := vowelMap[string(s[windowStart])]
			if ok {
				vowelMap[string(s[windowStart])]--
			}
			windowStart++
			log.Println("The window is at positions ", windowStart, ",", windowEnd)
		}
	}
	log.Println("maxCount => ", maxCount)
	return maxCount
}

func returnMapSum(a map[string]int) int {
	var sum int
	for _, val := range a {
		sum += val
	}
	return sum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
