package main

import (
	"github.com/khvr1993/GOLANG/leetcode/decodeString/decodeString"
	"log"
)

/*
An encoded string is given, the task is to decode it. The pattern in which the strings are encoded is as follows.
<count>[sub_str] ==> The substring 'sub_str' appears <count> times.
Input : str = "1[b]"
Output : b
Input : str = "2[ab]"
Output : abab
Input : str = "2[a2[b]]"
Output : abbabb
Input : str = "3[b2[ca]]"
Output : bcacabcacabcaca
Input : str = "2[a2[b23]]"
Output : ab23b23ab23b23
Input : str = "2[a2[b]1[z]]"
Output : abbzabbz

*/
func main() {
	log.Println("Decode String ")
	str := "2[a2[b]1[z]2[h]]"
	op := decodeString.DecodeMessage(str)
	log.Println("Op => ", op)
}
