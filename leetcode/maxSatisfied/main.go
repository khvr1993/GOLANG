package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/khvr1993/GOLANG/leetcode/maxSatisfied/maxSatisfied"
)

func main() {
	stringPassed := os.Args[1:]
	stringPassed2 := os.Args[2:]
	noOfValToChange, _ := strconv.Atoi(os.Args[3])
	var output int
	var res2 []int
	var res4 []int
	res1 := strings.Split(stringPassed[0], ",")
	res3 := strings.Split(stringPassed2[0], ",")
	for i := 0; i < len(res1); i++ {
		val, _ := strconv.Atoi(res1[i])
		val2, _ := strconv.Atoi(res3[i])
		res2 = append(res2, val)
		res4 = append(res4, val2)

	}
	log.Println("The customers are => ", res2)
	log.Println("The grumpy times are => ", res4)
	fmt.Println(noOfValToChange)
	output = maxSatisfied.MaxSatisfied(res2, res4, noOfValToChange)
	fmt.Println(output)
}
