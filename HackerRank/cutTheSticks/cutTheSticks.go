package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the cutTheSticks function below.
func cutTheSticks(arr []int32) []int32 {
	var sticksPresent []int32
	var newArr []int
	var i, indx, slicer int32
	for i < int32(len(arr)) {
		newArr = append(newArr, int(arr[i]))
		i++
	}
	sort.Ints(newArr)
	log.Println("Sorted Array ", newArr)
	stickCount := int32(len(newArr))
	for stickCount > 0 {
		log.Println("Stick Count ", stickCount)
		sticksPresent = append(sticksPresent, int32(stickCount))
		indx = 0
		shortestStick := newArr[0]
		for indx < int32(len(newArr)) {
			log.Println("Shortest Stick ", newArr[0])
			newArr[indx] -= shortestStick
			if newArr[indx] == 0 {
				log.Println("Slice the value at ", indx)
				slicer = indx + 1
				log.Println("slicer ", slicer)
			}
			indx++
		}
		newArr = newArr[slicer:]
		log.Println("New array after slice ", newArr)
		stickCount = int32(len(newArr))
	}
	return sticksPresent
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

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := cutTheSticks(arr)

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
