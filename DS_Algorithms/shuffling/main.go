package main

import "github.com/khvr1993/GOLANG/DS_Algorithms/shuffling/shuffling"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}

	shuffling.Shuffle(&a)
	shuffling.PrintArray(&a)
}
