package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the squares function below.
func squares(a int32, b int32) int32 {
	var count, val int32
	// First find the Sqrt of a
	// floor the value and this becomes the starting Point
	// increment the starting point and get the sqrts
	//Once the new sqrt value is greater exit the loop
	startingPoint := math.Floor(math.Sqrt(float64(a - 1)))
	startingPoint++
	val = int32(math.Pow(startingPoint, 2))
	for val <= b && val >= a {
		log.Println("Computing square for ", startingPoint)
		val = int32(math.Pow(startingPoint, 2))

		if val <= b {
			count++
		} else {
			break
		}
		startingPoint++
	}
	log.Println("count ", count)
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		ab := strings.Split(readLine(reader), " ")

		aTemp, err := strconv.ParseInt(ab[0], 10, 64)
		checkError(err)
		a := int32(aTemp)

		bTemp, err := strconv.ParseInt(ab[1], 10, 64)
		checkError(err)
		b := int32(bTemp)

		result := squares(a, b)

		fmt.Fprintf(writer, "%d\n", result)
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
