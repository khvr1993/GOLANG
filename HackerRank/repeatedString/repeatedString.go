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

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	lengthOfaStr := int64(len(s))
	timesToRepeat := n / int64(lengthOfaStr)
	remainder := n % int64(lengthOfaStr)
	var countA, i, j int64

	for i < lengthOfaStr {
		if string(s[i]) == "a" {
			countA++
		}
		i++
	}
	log.Println("Count of the a's => ", countA)

	countA = countA * timesToRepeat

	log.Println("countA after repeat => ", countA)

	for j < remainder {
		if string(s[j]) == "a" {
			countA++
		}
		j++
	}

	return countA
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
