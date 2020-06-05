package queue

import (
	"fmt"
)

//Node creates node which holds value and pointer to next node
type Node struct {
	val  interface{}
	next *Node
}

//Queue creates the datatype queue
type Queue struct {
	front *Node
	tail  *Node
	size  int
}

//Enqueue add the value to the queue
func (Q *Queue) Enqueue(val interface{}) {
	node := &Node{
		val: val,
	}
	if Q.front == nil {
		Q.front = node
	} else {
		Q.tail.next = node
	}
	Q.tail = node
	Q.size++
}

//Dequeue pop out the first element from the queue
func (Q *Queue) Dequeue() interface{} {
	returnVal := Q.front.val
	Q.front = Q.front.next
	Q.size--
	return returnVal
}

//Size returns the size of the queue
func (Q *Queue) Size() int {
	return Q.size
}

//PrintQueue prints the queue
func (Q *Queue) PrintQueue() {
	currentNode := Q.front
	for currentNode.next != nil {
		fmt.Printf("%v\t", currentNode.val)
		currentNode = currentNode.next
	}
	fmt.Printf("%v\t\n", currentNode.val)
}
