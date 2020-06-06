package main

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/Deque/Deque"
)

func main() {
	log.Println("Testing Deque")
	var dq = new(Deque.Deque)
	log.Println("Size of Deque ", dq.Size())
	dq.AddFirst(10)
	dq.AddFirst(5)
	dq.AddFirst(4)
	dq.PrintQueue()
	dq.AddLast(15)
	dq.AddLast(20)
	log.Println("Size of Deque ", dq.Size())
	dq.PrintQueue()
	elem, _ := dq.RemoveFirst()
	log.Println("Removed Element ", elem)
	dq.PrintQueue()
	elem, _ = dq.RemoveLast()
	log.Println("Removed Element ", elem)
	dq.PrintQueue()
}
