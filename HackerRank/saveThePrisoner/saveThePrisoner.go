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

// Complete the saveThePrisoner function below.
func saveThePrisoner(n int, m int, s int) int {
	//n  number of prisoners
	//m number of sweets
	//s chair number
	//log.Println("saveThePrisoner ", n, m, s)
	var savePrisoner int
	noOftimesDist := int(m / n)
	remainingSweets := m % n
	log.Println("noOftimesDist => ", noOftimesDist, " remainingSweets => ", remainingSweets)
	if remainingSweets == 0 && noOftimesDist > 0 {
		savePrisoner = s - 1
		if savePrisoner == 0 {
			savePrisoner = n
		} //implies that the sweet ended at last prisoner and size is 1
	} else if noOftimesDist > 0 {
		savePrisoner = s + remainingSweets - 1
	} else {
		savePrisoner = s + m - 1
	}
	if savePrisoner > n { // there are sweets left to distribute and the prisoner position is in between at this point the summ will become greater than the prisoner count
		return savePrisoner - n
	}
	return savePrisoner

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int(tTemp)
	for tItr := 0; tItr < int(t); tItr++ {
		nms := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nms[0], 10, 64)
		checkError(err)
		n := int(nTemp)

		mTemp, err := strconv.ParseInt(nms[1], 10, 64)
		checkError(err)
		m := int(mTemp)

		sTemp, err := strconv.ParseInt(nms[2], 10, 64)
		checkError(err)
		s := int(sTemp)

		result := saveThePrisoner(n, m, s)

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
