package decodeString

import (
	"log"
	"strconv"
)

/*
	Takes the n and returns the repeated string
*/
func repeatString(s string, n string) string {
	itr, err := strconv.Atoi(n)
	var op string
	if err == nil {
		log.Println("Begin repeat String ", n)
		for itr > 0 {
			op = op + s
			itr--
		}
	}
	return op
}
