package linkedList

import "log"

//Node is a singly linked List. Value holds the data we pushed and next holds
//the reference
type Node struct {
	value int
	next  *Node
}

//List we maintain what is the head node(First Element) and what is the tail element
type List struct {
	head *Node
	tail *Node
}

// AddMem Adds members to List
func (L *List) AddMem(value int) error {
	node := &Node{
		value: value,
	}
	//log.Println("Address of the node inserting => ", &node, " value of the node => ", node.value, "node.next", &node.next)
	//When we have head is Nil. This is the first element. so we make head as the node
	/*
		In the first run for the first insert head and tail will be same
		From the second run when we get the tail we are essentially getting the Head value and
		assigning the head.next to the node
		Later we are replacing tail with the new value
		In second push we take the tail that is the second element and assign the next to the pushed value
		Then we create a new tail which is the last element

		In every run the tail will have new element address. The tail is to ensure that we need to loop through the elements

	*/
	if L.head == nil {
		L.head = node
	} else {
		// if the head is already present we traverse through the list and come to a
		// point where the node is nil
		currentNode := L.tail
		// for currentNode.next != nil {
		// 	currentNode = currentNode.next
		// }
		currentNode.next = node
		//After reaching the end we add the node
	}
	L.tail = node

	//log.Println("Address of the head node POINTER after assignment", &L.head.next, "Head value ", L.head)
	//log.Println("Address of the tail node POINTER after assignment", &L.tail.next, "Tail value", L.tail)

	return nil
}

//ShowHead returns the Head Node
func (L *List) ShowHead() *Node {
	return L.head
}

//ShowTail returns the pointer to tail
func (L *List) ShowTail() *Node {
	return L.tail
}

//PrintList prints the elements inside the list;
func (L *List) PrintList() {
	//First Check if the List is nil
	if L.head == nil {
		//log.Println("Oops the List is nil")
		return
	}
	currentNode := L.head
	log.Println(currentNode.value, " ", &currentNode.next)
	for currentNode.next != nil {
		currentNode = currentNode.next
		log.Println(currentNode.value, " ", &currentNode.next)
	}
}

//First returns the head Node
func (L *List) First() *Node {
	return L.head
}

//FirstVal returns the head Node
func (L *List) FirstVal() int {
	return L.head.value
}

//Next return next pointer in the Node
func (n *Node) Next() *Node {
	return n.next

}

//NextVal return next pointer in the Node
func (n *Node) NextVal() int {
	return n.next.value

}

//CurrVal return next pointer in the Node
func (n *Node) CurrVal() int {
	return n.value

}
