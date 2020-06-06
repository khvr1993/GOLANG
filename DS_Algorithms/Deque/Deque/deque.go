package Deque

import (
	"errors"
	"fmt"
)

//Node creates a struct with val and a next of pointer Node
type Node struct {
	val  interface{}
	next *Node
}

//Deque creates a struct with head and tail
type Deque struct {
	front *Node
	back  *Node
	size  int
}

//Size returns the size of the queue
func (Q *Deque) Size() int {
	return Q.size
}

//AddFirst adds it to the end of the list
func (Q *Deque) AddFirst(val interface{}) {
	node := &Node{
		val: val,
	}
	if Q.front == nil {
		Q.front = node
		Q.back = node
	} else {
		currentNode := Q.front
		Q.front = node
		Q.front.next = currentNode
	}
	Q.size++
}

//AddLast adds the value to the end of the list
func (Q *Deque) AddLast(val interface{}) {
	node := &Node{
		val: val,
	}
	if Q.back == nil {
		Q.front = node
		Q.back = node
	} else {
		//First point the nex of the last element to the node
		// The new last element will be pointing to the back
		Q.back.next = node
		Q.back = Q.back.next
	}
	Q.size++
}

//RemoveFirst removes the item from the head
func (Q *Deque) RemoveFirst() (interface{}, error) {
	if Q.size == 0 {
		return nil, errors.New("Queue is empty")
	}
	currentVal := Q.front.val
	nextNode := Q.front.next
	Q.front = nextNode
	return currentVal, nil
}

//RemoveLast remove the last element from the queue
func (Q *Deque) RemoveLast() (interface{}, error) {
	if Q.size == 0 {
		return nil, errors.New("Queue is empty")
	}
	returnVal := Q.back.val
	currentNode := Q.front
	for currentNode.next.next != nil {
		currentNode = currentNode.next
	}
	Q.back = currentNode
	Q.back.next = nil
	return returnVal, nil
}

//PrintQueue prints the queue
func (Q *Deque) PrintQueue() {
	//	log.Println("Q.back", Q.back)
	currentNode := Q.front
	for currentNode.next != nil {
		fmt.Printf("%v\t", currentNode.val)
		currentNode = currentNode.next
	}
	fmt.Printf("%v\t\n", currentNode.val)
}
