package main

import (
	"log"

	stack "github.com/khvr1993/GOLANG/DS_Algorithms/Stack/Stack"
)

func main() {
	log.Println("Teesting the Stack")
	var s *stack.Stack = new(stack.Stack)
	s.Push(100)
	s.Push(200)
	s.Push(300)
	s.Push(400)
	s.Push(500)
	s.Push(600)
	s.Size()
	s.ShowStack()
	s.Pop()
	s.ShowStack()
	log.Println("Peek ", s.Peek())
}
