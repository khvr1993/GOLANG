package main

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/linkedStackOfElements/stack"
)

func main() {
	log.Println("testing linkedStackOfElements")
	var s = new(stack.Stack)
	s.Push("A")
	s.Push("B")
	s.Push("C")
	s.Push("D")
	s.Push("E")
	s.Push("F")
	log.Println("Size of the Stack ", s.Size())
	val := s.Pop()
	log.Println("Popped Element ", val)
	val = s.Pop()
	log.Println("Popped Element ", val)
	log.Println("Size of the Stack after Pop ", s.Size())
	s.ShowStack()
}
