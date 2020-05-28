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

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) string {
	var i, j int32
	initStr := make(map[int32]string)
	finalStr := make(map[int32]string)
	initStrlen := int32(len(s))
	finalStrlen := int32(len(t))
	for i < initStrlen {
		initStr[i] = string(s[i])
		i++
	}
	log.Println("Map converted s => ", initStr)

	for j < finalStrlen {
		finalStr[j] = string(t[j])
		j++
	}
	log.Println("Map Converted t => ", finalStr)

	i, j = 0, 0
	//find the common strings
	for i < initStrlen {
		if initStr[i] == finalStr[i] {
			i++
		} else {
			break
		}
	}
	leftOverStrLenInit := initStrlen - (i)
	leftOverStrLenFinal := finalStrlen - (i)
	log.Println("leftOverStrLenInit => ", leftOverStrLenInit, "leftOverStrLenFinal => ", leftOverStrLenFinal)
	if leftOverStrLenInit+leftOverStrLenFinal > k {
		log.Println("returning No")
		return "No"
	}

	if (leftOverStrLenInit+leftOverStrLenFinal) == 0 && (k%2 == 0 || k > (initStrlen+finalStrlen)) {
		return "Yes"
	} else if (leftOverStrLenInit + leftOverStrLenFinal) == k {
		return "Yes"
	} else if (k-(leftOverStrLenInit+leftOverStrLenFinal))%2 == 0 || (k > (initStrlen + finalStrlen)) {
		return "Yes"
	}
	return "No"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	t := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := appendAndDelete(s, t, k)

	fmt.Fprintf(writer, "%s\n", result)

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
