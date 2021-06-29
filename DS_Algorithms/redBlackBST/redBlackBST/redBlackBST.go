package redBlackBST

import (
	"fmt"
	"log"
)

const (
	RED   bool = true
	BLACK bool = false
)

type Node struct {
	Key       string
	Value     int
	LeftNode  *Node
	RightNode *Node
	Color     bool
	Size      int
}

func (node *Node) isRed() bool {
	if node == nil {
		return false
	}
	return node.Color == RED
}

// Size returns the number of leaves for the node in question
func Size(root *Node) int {
	if root == nil {
		return 0
	}
	return root.Size
}

func rotateLeft(h *Node) *Node {
	/*
		Converts a right leaning red black BST (RED link present to the right) to a left leaning red black BST
			   1. Copy the right link to a temp variable
			   2. make temps left link to the right link of root
			   3. Make the temp variable root
			   4. temp.left will be the previous root
	*/
	temp := h.RightNode
	h.RightNode = temp.LeftNode
	temp.LeftNode = h
	temp.Color = h.Color
	h.Color = RED
	temp.Size = h.Size
	h.Size = 1 + Size(h.LeftNode) + Size(h.RightNode)
	return temp
}

func rotateRight(h *Node) *Node {
	temp := h.LeftNode
	h.LeftNode = temp.RightNode
	temp.RightNode = h
	temp.Color = h.Color
	h.Color = RED
	temp.Size = h.Size
	h.Size = 1 + Size(h.LeftNode) + Size(h.RightNode)
	return temp
}

func flipColors(h *Node) {
	h.Color = RED
	h.LeftNode.Color = BLACK
	h.RightNode.Color = BLACK
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

// ======================================================================

// RedBlackBST implements red Black tree
type RedBlackBST struct {
	Root *Node
}

func NewRedBlackBST() *RedBlackBST {
	var newNode *Node
	RedBlackBST := RedBlackBST{newNode}
	return &RedBlackBST
}

// Put places the Key in the right position
func (bst *RedBlackBST) Put(Key string, Value int) *Node {
	println("Put Key", Key, " Value ", Value)
	var newNode *Node
	if bst.Root == nil {
		println("No root present for the BST. Making the Key sent as the Node")
		newNode = &Node{Key, Value, nil, nil, RED, 1}
		bst.Root = newNode
	} else {
		println("Root is present for the BST. Calling func to find the right place to put the Key")
		bst.Root = put(bst.Root, Key, Value)
	}
	bst.Root.Color = BLACK
	return bst.Root
}

func put(root *Node, key string, value int) *Node {

	if root == nil {
		log.Println("Returning the node address")
		return &Node{key, value, nil, nil, RED, 1}
	}

	if key < root.Key {
		log.Println("Root is greater than the Key")
		root.LeftNode = put(root.LeftNode, key, value)
	} else if key > root.Key {
		log.Println("Root is less than the Key")
		root.RightNode = put(root.RightNode, key, value)
	} else {
		log.Println("Root is equal to the Key")
		root.Value = value
		root.Key = key
	}

	// Right Node is red and left node is not red
	if root.RightNode.isRed() && !root.LeftNode.isRed() {
		root = rotateLeft(root)
	}
	if root.LeftNode.isRed() && root.LeftNode.LeftNode.isRed() {
		// On the left side 2 red nodes are contiguos
		root = rotateRight(root)
	}
	if root.LeftNode.isRed() && root.RightNode.isRed() {
		flipColors(root)
	}

	root.Size = 1 + Size(root.LeftNode) + Size(root.RightNode)

	return root
}

// Height returns the height of the binary Tree
func (bst *RedBlackBST) Height(root *Node) int {
	var h int
	h = height(root)
	return h
}

// LevelOrderTraversal or Breadth First Traversal starts at the tree root (or some arbitrary node of a graph, sometimes referred to as a 'search Key'[1]),
// and explores all of the neighbor nodes at the present depth prior to moving on to the nodes at the next depth level
func (bst *RedBlackBST) LevelOrderTraversal(root *Node) {
	levelOrderTraversal(root)
	fmt.Println("")
}

func levelOrderTraversal(root *Node) {
	h := height(root)
	for i := 1; i <= h; i++ {
		// From the root traverse through each level and print the leaves

		printGivenLevel(root, i)
	}
}

//  printGivenLevel prints the nodes at the level
// For example when level = 3 then all the leaves at level 3 are printed. This is acheived by level = 1 condition
// When we call recursively from the root level gets decreased by 1 and we call root.left and root.right
// At the end we reach the right level to print the leaves
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

// InOrderTraversal prints the tree in ascending order
// Left Root Right
func (bst *RedBlackBST) InOrderTraversal() {
	inorderTraversal(bst.Root)
}

func inorderTraversal(node *Node) {
	if node != nil {
		inorderTraversal(node.LeftNode)
		fmt.Print(node.Key, " ")
		inorderTraversal(node.RightNode)
	}
}
