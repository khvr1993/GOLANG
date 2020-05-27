package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/khvr1993/GOLANG/leetcode/maxTurbulenceSize/maxTurbulenceSize"
)

func main() {
	stringPassed := os.Args[1:]
	var output int
	var res2 []int
	res1 := strings.Split(stringPassed[0], ",")
	for i := 0; i < len(res1); i++ {
		val, _ := strconv.Atoi(res1[i])
		res2 = append(res2, val)
	}
	log.Println("res2 => ", res2)
	output = maxTurbulenceSize.MaxTurbulenceSize(res2)
	fmt.Println(output)
}
