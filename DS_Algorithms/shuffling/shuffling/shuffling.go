package shuffling

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//Shuffle shuffles the array randomly
func Shuffle(a *[]int) {
	var i int
	length := len(*a)
	//Top-level functions, such as Float64 and Int, use a default shared Source that
	//produces a deterministic sequence of values each time a program is run.
	//Use the Seed function to initialize the default Source
	//if different behavior is required for each run.
	rand.Seed(time.Now().UnixNano())
	for i < length {
		//generate a ranom number
		randNum := rand.Intn(i + 1)
		log.Println("randNum => ", randNum)
		tempVal := (*a)[randNum]
		(*a)[randNum] = (*a)[i]
		(*a)[i] = tempVal
		i++
	}
}

//PrintArray Prints the values
func PrintArray(a *[]int) {
	var i int
	for i < len(*a) {
		fmt.Printf("%v\t", (*a)[i])
		i++
	}
	fmt.Printf("\n")
}
