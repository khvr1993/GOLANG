package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/khvr1993/GOLANG/DS_Algorithms/binarySearch/binarySearch"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int

	for arItr := 0; arItr < int(arCount); arItr++ {
		arItemTemp, err := strconv.ParseInt(arTemp[arItr], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, int(arItem))
	}

	targetElement, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	log.Println("targetElement => ", targetElement)

	_, result := binarySearch.BinarySearch(ar, int(targetElement))

	fmt.Fprintf(writer, "%v\n", result)

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
