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

// Complete the angryProfessor function below.
func angryProfessor(k int32, a []int32) string {
	var noOflatestudents int
	var i int
	var lenOfArray = len(a)
	for i < lenOfArray {
		if a[i] > 0 {
			noOflatestudents++
		}
		i++
	}
	log.Println("No of Late Students ", noOflatestudents)
	if int32(lenOfArray-noOflatestudents) >= k {
		return "NO"
	}
	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	log.Println(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer stdout.Close()
	writer := bufio.NewWriterSize(stdout, 1024*1024)
	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)
	for tItr := 0; tItr < int(t); tItr++ {
		nk := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nk[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		temp, err := strconv.ParseInt(nk[1], 10, 64)
		checkError(err)
		k := int32(temp)

		aTemp := strings.Split(readLine(reader), " ")

		var a []int32

		for i := 0; i < int(n); i++ {
			aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
			checkError(err)
			aItem := int32(aItemTemp)
			a = append(a, aItem)
		}

		result := angryProfessor(k, a)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
