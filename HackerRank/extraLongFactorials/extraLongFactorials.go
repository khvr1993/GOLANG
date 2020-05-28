package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// Complete the extraLongFactorials function below.
func extraLongFactorials(n int32) {
	var newBigInt big.Int
	fact := n
	newBigInt = *calcFact(fact)
	fmt.Println(&newBigInt)
}

func calcFact(n int32) *big.Int {
	//log.Println("Calculating for ", n)
	if n == 0 {
		return big.NewInt(int64(1))
	}
	val := Mul(big.NewInt(int64(n)), calcFact(n-1))
	//log.Println("val", val)
	return val
}

//Mul returns the multiplication
func Mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	extraLongFactorials(n)
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
