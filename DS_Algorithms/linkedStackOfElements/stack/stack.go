package stack

import (
	"fmt"
)

//Node contains value and the pointer to the next Element
type Node struct {
	value string
	next  *Node
}

//Stack is the implementation of stack through linked list
type Stack struct {
	top  *Node
	size int
}

//Push inserts the value at the top
func (S *Stack) Push(val string) {
	node := &Node{
		value: val}

	if S.top == nil {
		S.top = node
	} else {
		node := &Node{
			value: val,
			next:  S.top}
		S.top = node
	}
	S.size++
}

//Pop returns the latest element pushed
func (S *Stack) Pop() string {
	topVal := S.top.value
	S.top = S.top.next
	S.size--
	return topVal
}

//ShowStack Prints the elements present in the stack
func (S *Stack) ShowStack() {
	currentNode := S.top
	for currentNode.next != nil {
		fmt.Printf(currentNode.value + " ")
		currentNode = currentNode.next
	}
	fmt.Printf(currentNode.value + " ")
	fmt.Printf("\n")
}

//Size returns the size of the stack
func (S *Stack) Size() int {
	return S.size
}
