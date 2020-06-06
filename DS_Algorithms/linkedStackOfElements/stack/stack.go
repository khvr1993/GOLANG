package stack

import (
	"errors"
	"fmt"
)

//Node contains value and the pointer to the next Element
type Node struct {
	value interface{}
	next  *Node
}

//Stack is the implementation of stack through linked list
type Stack struct {
	top      *Node
	size     int
	iterator *Node
}

//Push inserts the value at the top
func (S *Stack) Push(val interface{}) {
	node := &Node{
		value: val}

	if S.top == nil {
		S.top = node
	} else {
		//Assign the current element to value and the top element to the next
		//Then point the top to the node
		node := &Node{
			value: val,
			next:  S.top}
		S.top = node
	}
	S.size++
}

//Pop returns the latest element pushed
func (S *Stack) Pop() interface{} {
	if S.top == nil {
		return nil
	}
	topVal := S.top.value
	if S.top.next != nil {
		S.top = S.top.next
	} else {
		S.top = nil
	}
	S.size--
	return topVal
}

//ShowStack Prints the elements present in the stack
func (S *Stack) ShowStack() {
	if S.top == nil {
		return
	}
	currentNode := S.top
	for currentNode.next != nil {
		fmt.Printf("%v\t", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Printf("%v\t", currentNode.value)
	fmt.Printf("\n")
	return
}

//Size returns the size of the stack
func (S *Stack) Size() int {
	return S.size
}

//PointToHead will point the iterator to Head
func (S *Stack) PointToHead() {
	if S.top != nil {
		S.iterator = S.top
	}
}

//NextItem returns the item net to the iterator
func (S *Stack) NextItem() (interface{}, error) {
	var returnVal interface{}

	if S.iterator == nil {
		return nil, errors.New("Set the iterator")
	}
	if S.iterator.next == nil {
		returnVal = S.iterator.value
		error := errors.New("Reached end of the List")
		return returnVal, error
	}
	returnVal = S.iterator.value
	S.iterator = S.iterator.next

	return returnVal, nil
}
