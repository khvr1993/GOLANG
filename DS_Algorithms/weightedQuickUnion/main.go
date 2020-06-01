package main

/*
1. In this we are taking input from the commad Line.
2. The first line of the argument will be the number of pairs we are inputting.
3. The subsequent lines will have arrays
eg : The arguments can be entered as Follows
						10  -- No of points
						6    --No of Unions
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

	"github.com/khvr1993/GOLANG/DS_Algorithms/weightedQuickUnion/quickUnion"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	//Path Size
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	UFSize := int(nTemp)
	var uf *quickUnion.UF = new(quickUnion.UF)
	//Path Size
	uf.CreatePoints(UFSize)

	//uf.PrintPathTransform()

	//No of Union Paths
	nTemp, err = strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)
	//No of Union Paths

	//Union Points
	for i := 0; i < int(n); i++ {
		arrTemp := strings.Split(readLine(reader), " ")

		aTemp, err := strconv.ParseInt(arrTemp[0], 10, 64)
		checkError(err)

		p := int(aTemp)

		bTemp, err := strconv.ParseInt(arrTemp[1], 10, 64)
		checkError(err)
		q := int(bTemp)
		uf.Union(p, q)
	}
	//Union Points

	/*
		//for connection exists check
		arrTemp := strings.Split(readLine(reader), " ")
		aTemp, err := strconv.ParseInt(arrTemp[0], 10, 64)
		checkError(err)

		p := int(aTemp)

		bTemp, err := strconv.ParseInt(arrTemp[1], 10, 64)
		checkError(err)
		q := int(bTemp)

		//for connection exists Check
		ok := uf.IsConnected(p, q)

		log.Println("IsConnected ", ok)
	*/

	treeDepth := uf.Size(3)
	log.Println("TreeDepth ", treeDepth)

	result := uf.PathCreated()

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
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
