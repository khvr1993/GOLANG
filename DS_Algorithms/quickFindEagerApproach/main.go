package main

/*
1. In this we are taking input from the commad Line.
2. The first line of the argument will be the number of pairs we are inputting.
3. The subsequent lines will have arrays
eg : The arguments can be entered as Follows
						10
						6
						1 2
						3 4
						4 5
						1 3
						1 4
						2 6
*/
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	UFSize := int32(nTemp)
	log.Println("UF Size", UFSize)

	nTemp, err = strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	for i := 0; i < int(n); i++ {
		arrTemp := strings.Split(readLine(reader), " ")

		aTemp, err := strconv.ParseInt(arrTemp[0], 10, 64)
		checkError(err)

		p := int32(aTemp)

		bTemp, err := strconv.ParseInt(arrTemp[1], 10, 64)
		checkError(err)
		q := int32(bTemp)
		log.Println("p => ", p)
		log.Println("q => ", q)

	}

	result := []int{1, 2, 3}

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
