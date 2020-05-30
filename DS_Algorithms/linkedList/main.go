package main

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/linkedList/linkedList"
)

func main() {
	log.Println("Testing LinkedList ")
	var L linkedList.List
	//var nodeElem linkedList.Node
	L.AddMem(1)
	L.AddMem(2)
	L.AddMem(3)
	L.AddMem(4)
	L.PrintList()

	node := L.First()
	log.Println("value  ", node.CurrVal())
	log.Println("value ", node.NextVal())
	node = node.Next()
	log.Println("value ", node.NextVal())
	node = node.Next()
	log.Println("value ", node.NextVal())
	node = node.Next()
	if node.Next() != nil {
		log.Println("value ", node.NextVal())
	}
}
