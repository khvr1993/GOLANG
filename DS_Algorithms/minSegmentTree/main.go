package main

import (
	"github.com/khvr1993/GOLANG/DS_Algorithms/minSegmentTree/minSegmentTree"
	"github.com/khvr1993/GOLANG/utils"
	"log"
)

func main() {
	log.Println("Testing Min Segment Tree ")
	arr := []int{-1, 2, 4, 0}
	segTree := minSegmentTree.BuildMinSegmentTree(&arr)

	utils.PrintArray(segTree)

	op := minSegmentTree.RangeMinQuery(segTree, 0, 2, 0, len(arr)-1, 0)
	log.Printf("op is %d", op)
}
