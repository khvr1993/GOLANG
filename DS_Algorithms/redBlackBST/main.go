package main

import (
	redBlackBST2 "github.com/khvr1993/GOLANG/DS_Algorithms/redBlackBST/redBlackBST"
	"log"
)

func main() {
	println("Testing red black BST")
	redBlackBST := redBlackBST2.NewRedBlackBST()
	redBlackBST.Put("S", 1)
	redBlackBST.Put("E", 1)
	redBlackBST.Put("A", 1)
	redBlackBST.Put("R", 1)
	redBlackBST.Put("C", 1)
	redBlackBST.Put("H", 1)
	redBlackBST.Put("E", 1)
	redBlackBST.Put("X", 1)
	redBlackBST.Put("A", 1)
	redBlackBST.Put("M", 1)
	redBlackBST.Put("P", 1)
	redBlackBST.Put("L", 1)
	redBlackBST.Put("E", 1)
	log.Println("++==================++")
	redBlackBST.LevelOrderTraversal(redBlackBST.Root)
}
