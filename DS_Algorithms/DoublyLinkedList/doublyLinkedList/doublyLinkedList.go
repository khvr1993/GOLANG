package doublyLinkedList

import "log"

//Node will create a Node for Doubly Linked List
type Node struct {
	value int
	prev  *Node
	next  *Node
}

//List creates a struct with a head Node and a tail Node
type List struct {
	head *Node
	tail *Node
}

//AddMem add the data passed to linked List
func (L *List) AddMem(value int) {
	node := &Node{
		value: value,
	}

	//Check whether this is the first element => Head will be Null
	if L.head == nil {
		L.head = node
	} else {
		currentNode := L.tail
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = node
		node.prev = L.tail
	}
	L.tail = node
	return
}

//PrintList Prints all the elements in the List
func (L *List) PrintList() {
	if L.head == nil {
		log.Println("Ooops the List is empty")
		return
	}
	currentNode := L.head
	for currentNode.next != nil {
		log.Println("Node", currentNode.value)
		currentNode = currentNode.next
	}
	//Print the Last Element
	currentNode = L.tail
	log.Println("The Last Node => ", currentNode.value)
	return
}

//PrintListReverse prints the list in the reverse Order
func (L *List) PrintListReverse() {
	if L.head == nil {
		log.Println("Ooops the List is empty")
		return
	}
	currentNode := L.tail
	for currentNode.prev != nil {
		log.Println("Node", currentNode.value)
		currentNode = currentNode.prev
	}
	//Print the Last Element
	currentNode = L.head
	log.Println("The First Node => ", currentNode.value)
	return
}

//Prev returns the next Node
func (n *Node) Prev() *Node {
	return n.prev
}

//GetHeadNode Get the Head Node
func (L *List) GetHeadNode() *Node {
	if L.head == nil {
		log.Println("The node requested cannot be returned as the list is empty")
		return nil
	}
	return L.head
}

//Value returns the value of the node
func (n *Node) Value() int {
	return n.value
}

//Next returns the pointer to the next node
func (n *Node) Next() *Node {
	return n.next
}

//FindValue value exists
func (L *List) FindValue(a int) (*Node, bool) {
	//get the head
	currentNode := L.head
	for currentNode.next != nil {
		if currentNode.value == a {
			return currentNode, true
		}
		currentNode = currentNode.next
	}
	currentNode = L.tail
	if currentNode.value == a {
		return currentNode, true
	}
	return nil, false
}
