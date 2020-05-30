package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */

func nonDivisibleSubset(k int32, s []int32) int32 {
	// Calculate the remainders and create a map with remainder as Key and count as value
	// increment the count when remainder is encountered.
	var i, j, maxSum int32
	j = 1
	var remainders = make(map[int32]int32)
	//var keys []int
	for i < int32(len(s)) {
		rem := s[i] % k
		_, ok := remainders[rem]
		if ok {
			remainders[rem]++
		} else {
			remainders[rem] = 1
		}
		i++
	}

	log.Println("Remainder List ", remainders, len(remainders))

	for key := range remainders {
		if key == 0 {
			maxSum++
		} else if 2*key%k == 0 {
			maxSum++
		}
	}

	for j <= k/2 {
		log.Println("k/2 => ", k/2)
		if 2*j%k != 0 {
			maxSum += max(remainders[j], remainders[k-j])
		}
		j++
	}

	log.Println("maxSum => ", maxSum)
	return maxSum

}

func max(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	result := nonDivisibleSubset(k, s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
