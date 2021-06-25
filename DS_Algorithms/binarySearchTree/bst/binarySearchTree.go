package bst

import (
	"fmt"
	"log"
)

type Node struct {
	Key       int
	Value     string
	LeftNode  *Node `default:"nil"`
	RightNode *Node `default:"nil"`
	Size      int   `default:"0"`
}

type BinarySearchTree struct {
	root *Node
}

// NewBST creates a BinaryTree structure and returns the pointer to it
func NewBST(Key int, Value string) (*BinarySearchTree, *Node) {
	var bst BinarySearchTree
	root := new(Node)
	root.Key = Key
	root.Value = Value
	root.LeftNode = nil
	root.RightNode = nil
	root.Size = 1
	bst.root = root
	return &bst, root
}

// Put places the Key in the right position
func (bst *BinarySearchTree) Put(Key int, Value string) *Node {
	var newNode *Node
	if bst.root == nil {
		log.Println("No root present for the BST. Making the Key sent as the Node")
		newNode = &Node{Key, Value, nil, nil, 1}
		bst.root = newNode
		return newNode
	} else {
		log.Println("Root is present for the BST. Calling func to find the right place to put the Key")
		newNode = bst.put(bst.root, Key, Value)
	}
	return newNode
}

// put is a private function which will insert the node into the correct place
func (bst *BinarySearchTree) put(root *Node, Key int, Value string) *Node {

	log.Println("Key ", Key, " Value ", Value)

	if bst.root == nil {
		log.Println("The Root of the BST is empty. Assingning the root Node")
		newNode := Node{Key, Value, nil, nil, 1}
		bst.root = &newNode
		return bst.root
	}
	if root == nil {
		log.Println("Returning the node address")
		return &Node{Key, Value, nil, nil, 1}
	}

	// Inserting into the left Node for all the values of root which are < than the Key to the left
	if root.Key > Key {
		log.Println("Root is greater than the Key")
		root.LeftNode = bst.put(root.LeftNode, Key, Value)
	} else if root.Key < Key {
		log.Println("Root is less than the Key")
		root.RightNode = bst.put(root.RightNode, Key, Value)
	} else {
		log.Println("Root is equal to the Key")
		root.Value = Value
		root.Key = Key
	}
	root.Size = 1 + size(root.LeftNode) + size(root.RightNode)
	return root
}

// returns the size of the node (How many leaves are there for a root Node)
func size(root *Node) int {
	if root == nil {
		return 0
	} else {
		return root.Size
	}
}

// Height returns the height of the binary Tree
func (bst *BinarySearchTree) Height(root *Node) int {
	var h int
	h = height(root)
	return h
}

// height is a private function which calculates the height of the binary Tree
func height(root *Node) int {
	var leftHeight, rightHeight int

	if root == nil {
		return 0
	} else {
		// left elements height and right elements height
		leftHeight = height(root.LeftNode)
		rightHeight = height(root.RightNode)
	}
	if leftHeight > rightHeight {
		return 1 + leftHeight
	} else {
		return 1 + rightHeight
	}
}

// LevelOrderTraversal or Breadth First Traversal starts at the tree root (or some arbitrary node of a graph, sometimes referred to as a 'search Key'[1]),
// and explores all of the neighbor nodes at the present depth prior to moving on to the nodes at the next depth level
func (bst *BinarySearchTree) LevelOrderTraversal(root *Node) {
	levelOrderTraversal(root)
	fmt.Println("")
}

func levelOrderTraversal(root *Node) {
	h := height(root)
	for i := 1; i <= h; i++ {
		printGivenLevel(root, i)
	}
}

func printGivenLevel(root *Node, level int) {
	if root == nil {
		return
	}
	if level == 1 {
		fmt.Print(root.Key, " ")
	} else {
		printGivenLevel(root.LeftNode, level-1)
		printGivenLevel(root.RightNode, level-1)
	}
}

func (bst *BinarySearchTree) DeleteMin() {
	deleteMin(bst.root)
}

func deleteMin(root *Node) *Node {
	// exit
	if root.LeftNode == nil {
		return root.RightNode
	} else {
		/*For the left node after reaching the end it will either assign nil to the
		 last node when both left and right nodes are null or it will replace the last node
		with its right leaf
		*/
		root.LeftNode = deleteMin(root.LeftNode)
	}
	root.Size = 1 + size(root.LeftNode) + size(root.RightNode)
	return root
}
