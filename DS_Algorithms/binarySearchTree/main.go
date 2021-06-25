package main

import (
	bst2 "github.com/khvr1993/GOLANG/DS_Algorithms/binarySearchTree/bst"
	"log"
)

func main() {
	log.Println("Testing Binary Search Tree")
	bst, root := bst2.NewBST(50, "a")
	bst.Put(100, "b")
	bst.Put(40, "c")
	bst.Put(60, "d")
	bst.Put(110, "e")
	bst.Put(40, "f")
	log.Println("Height ", bst.Height(root))
	bst.LevelOrderTraversal(root)
	log.Println(root.LeftNode.Size)
	bst.DeleteMin()
	bst.LevelOrderTraversal(root)
}
