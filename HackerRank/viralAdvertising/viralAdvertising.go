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

// Complete the viralAdvertising function below.
func viralAdvertising(n int32) int32 {
	firstDay := int32(5)
	var i int32 = 1
	var noOfLiked int32
	var noOfShared int32
	var cumulativeLiked int32
	for i <= n {
		if i == 1 {
			noOfLiked += firstDay / 2
			noOfShared := firstDay
			cumulativeLiked += noOfLiked
			log.Println("Shared after first DayEnd ", noOfShared)
		} else {

			noOfShared = noOfLiked * 3
			noOfLiked = noOfShared / 2
			cumulativeLiked += noOfLiked
			log.Println("Cumulatively Liked ", noOfLiked, " at the end of ", i)
			log.Println("Shared to ", noOfShared, " at the end of ", i)
		}
		i++
	}
	log.Println("Cumulatively liked ", cumulativeLiked)
	return cumulativeLiked

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

	result := viralAdvertising(n)

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
