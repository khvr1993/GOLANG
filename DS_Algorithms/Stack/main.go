package main

import (
	"log"

	stack "github.com/khvr1993/GOLANG/DS_Algorithms/Stack/Stack"
)

func main() {
	log.Println("Testing the Stack")
	//var s = stack.NewStrStack()
	test := stack.Stack(stack.NewIntStack())
	test.Push(100)
	test.Push(200)
	test.Push(300)
	test.Push(400)
	test.Push(500)
	test.Push(600)
	test.Size()
	test.ShowStack()
	test.Pop()
	test.ShowStack()
	log.Println("Peek ", test.Peek())
	//test.Push("a")
	//test.Push("b")
	//log.Println("Peek ", test.Peek())
	//log.Println("Pop ", test.Pop())

}
