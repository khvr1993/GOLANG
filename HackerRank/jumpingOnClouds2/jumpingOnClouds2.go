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

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	var jumps int32
	var start, end, windowLen int32
	lengthOfStr := int32(len(c))
	for end < lengthOfStr {
		if c[end] == 0 {
			log.Println("The end value ", end)
			//increment the zero Window
			end++
		} else {
			log.Println("WindowLen => ", windowLen, "Found 1 at => ", end)
			windowLen = end - start
			jumps += windowLen / 2
			end++
			start = end
			jumps++
			log.Println("jumps => ", jumps)
		}
	}
	if start != end {
		log.Println("Start and end have not reached last pos. There is end window => ", end, start)
		windowLen = end - start
		jumps += windowLen / 2
	}
	log.Println("jumps => ", jumps)
	return jumps
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := jumpingOnClouds(c)

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
