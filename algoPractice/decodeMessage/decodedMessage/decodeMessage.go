package decodedMessage

import (
	"log"
)

//Solution performs the recursive operation to solve the decoding problem
func Solution(a []int) int {
	n := len(a)
	j := decode(a, n)
	log.Println("No of decodings ", j)
	return j
}

func decode(a []int, n int) int {
	var j int
	if n == 1 {
		j++
		return 1
	}
	if n == 2 {
		if a[0] <= 2 {
			j += 2
			return 2
		}
		return 1
	}
	if a[n-2] <= 2 && a[n-1] < 7 {
		return decode(a[0:n-2], n-2) + decode(a[0:n-1], n-1)
	}
	return decode(a[:n-1], n-1)

}
