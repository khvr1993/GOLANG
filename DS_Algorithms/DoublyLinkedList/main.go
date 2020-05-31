package main

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/DoublyLinkedList/doublyLinkedList"
)

func main() {
	log.Println("Testing Doubly Linked List")
	var l doublyLinkedList.List
	l.AddMem(1)
	l.AddMem(2)
	l.AddMem(3)
	l.AddMem(4)
	l.PrintList()
	log.Println("Printing the List in Reverse")
	l.PrintListReverse()
	n := l.GetHeadNode()
	log.Println("The value of head node ", n.Value())
	n = n.Next()
	log.Println("The value of node ", n.Value())
	n = n.Prev()
	log.Println("The value of node ", n.Value())

	result, ok := l.FindValue(2)
	if ok {
		log.Println("result address ", &result)
	}

}
