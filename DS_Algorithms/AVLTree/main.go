package main

import (
	avlTree2 "github.com/khvr1993/GOLANG/DS_Algorithms/AVLTree/avlTree"
	"log"
)

func main() {
	log.Println("testing AVL Tree")
	avlTree := avlTree2.NewAVLTree()
	log.Println(avlTree.Root)
	avlTree.Put(40, 10)
	avlTree.Put(30, 10)
	avlTree.Put(35, 10)
	// avlTree.Put(40, 10)
	// avlTree.Put(50, 10)
	// avlTree.Put(60, 10)
	// avlTree.Put(70, 10)

	log.Println(avlTree.Root.Height)
	log.Println(avlTree.Root)
	log.Println(avlTree.Root.LeftNode)
	log.Println(avlTree.Root.RightNode)
	avlTree.LevelOrderTraversal()

}
