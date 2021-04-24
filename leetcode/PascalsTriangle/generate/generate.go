package generate

import "log"

func Generate(numRows int) [][]int {

	opArray := [][]int{}

	if numRows == 0 {
		return opArray
	}

	//Exit condition
	if numRows == 1 {
		opArray = append(opArray, []int{1})
		return opArray
	}

	log.Println("Creating memory array")

	// Create an array for storing the results of the previous calculations
	memList := make([][]int, numRows)
	log.Println(memList)
	memList[0] = append(memList[0], 1)
	memList[1] = append(memList[1], 1, 1)
	log.Println(memList)
	for i := 2; i < numRows; i++ {
		memList[i] = append(memList[i], 1)
		for j := 1; j <= i-1; j++ {
			memList[i] = append(memList[i], memList[i-1][j-1]+memList[i-1][j])
		}
		memList[i] = append(memList[i], 1)
	}
	//Initialise the first 2 datasets

	log.Println(memList)

	return memList
}
