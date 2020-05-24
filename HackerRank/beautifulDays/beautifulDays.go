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

// Complete the beautifulDays function below.
func beautifulDays(i int32, j int32, k int32) int32 {
	var countOfbeautifulDays int32
	for i <= j {
		reversalOfI := reverseNumber(i)
		log.Println("reverseNumber ", reversalOfI)
		difference := i - reversalOfI
		if difference%k == 0 {
			log.Println("This day is beautiful ", i, reversalOfI)
			countOfbeautifulDays++
		}
		i++
	}
	log.Println("no Of beautiful Days ", countOfbeautifulDays)
	return countOfbeautifulDays
}

func reverseNumber(i int32) int32 {
	newint := 0
	for i > 0 {
		remainder := int(i) % 10
		newint *= 10
		newint += remainder
		i /= 10
	}
	return int32(newint)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	ijk := strings.Split(readLine(reader), " ")

	iTemp, err := strconv.ParseInt(ijk[0], 10, 64)
	checkError(err)
	i := int32(iTemp)

	jTemp, err := strconv.ParseInt(ijk[1], 10, 64)
	checkError(err)
	j := int32(jTemp)

	kTemp, err := strconv.ParseInt(ijk[2], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := beautifulDays(i, j, k)

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
