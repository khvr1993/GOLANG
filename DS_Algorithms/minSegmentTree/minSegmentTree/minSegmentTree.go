package minSegmentTree

import (
	"github.com/khvr1993/GOLANG/utils"
	"math"
)

const MAXINT int = math.MaxInt32

/*
	Initialises the array and populates maximum values in it
*/
func initArrToMax(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = MAXINT
	}
}

/*
	This is a recursive function which calls itself and builds the min segment tree.
	The leaf nodes will have the array values and root nodes maintain the minimum value
*/
func buildTree(inp *[]int, segTree *[]int, lo int, hi int, pos int) {

	//log.Printf("Parameters => lo %d high %d pos %d ", lo, hi, pos)

	if lo == hi {
		(*segTree)[pos] = (*inp)[lo]
		return
	}
	mid := (lo + hi) / 2
	buildTree(inp, segTree, lo, mid, 2*pos+1)
	buildTree(inp, segTree, mid+1, hi, 2*pos+2)
	(*segTree)[pos] = utils.Min((*segTree)[2*pos+1], (*segTree)[2*pos+2])
}

/*
	If the value of n is equal to pow of 2 then return n*2-1
	If the value of n is not equal to the pow of 2 then get the next power of 2 and return
	For eg. if n = 5 then 2*8-1
*/
func getSegmentArrLen(n int) int {
	powCount := 0
	var size int
	for {
		size = int(math.Pow(float64(2), float64(powCount)))
		if size >= n {
			break
		} else {
			powCount++
		}
	}
	return size*2 - 1
}

/*
	Main calling procedure which builds the segment tree
*/
func BuildMinSegmentTree(arr *[]int) *[]int {

	sizeOfTree := getSegmentArrLen(len(*arr))
	segTree := make([]int, sizeOfTree)

	initArrToMax(&segTree)
	/*
		Build Tree starting from Root
	*/
	buildTree(arr, &segTree, 0, len(*arr)-1, 0)
	return &segTree
}
