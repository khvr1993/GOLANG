package main

/*
First Line - Size of Array
Second Line the Array

eg :
6
1 2 3 4 5 6

*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/khvr1993/GOLANG/DS_Algorithms/searchBitonicArray/searchBitonicArray"
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
		arItem := int(arItemTemp)
		ar = append(ar, arItem)
	}

	target, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := searchBitonicArray.FindTarget(ar, int(target))

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
