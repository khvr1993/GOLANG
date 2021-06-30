package avlTree

import (
	"fmt"
	"github.com/khvr1993/GOLANG/utils"
	"log"
)

type Node struct {
	Key       int
	Value     int
	LeftNode  *Node
	RightNode *Node
	Size      int
	Height    int
}

type AVLTree struct {
	Root *Node
}

func NewAVLTree() *AVLTree {
	// newNode := new(Node)
	var newNode *Node
	avlTree := AVLTree{newNode}
	return &avlTree
}

func NewNode(Key int, Value int) *Node {
	newNode := Node{Key, Value, nil, nil, 1, 1}
	return &newNode
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Size
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// nodeBalance will check and return the difference between the heights of left node and right node
// returns 0 if both left and right are perfectly balanced. +ve if left side is taller and -1 if right side is taller
func nodeBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.LeftNode) - height(node.RightNode)
}

// rotateLeft rotates the tree around the node to the left
func rotateLeft(h *Node) *Node {
	log.Println("rotateLeft ", h.Key)
	temp := h.RightNode
	h.RightNode = temp.LeftNode
	temp.LeftNode = h

	temp.Size = 1 + size(temp.LeftNode) + size(temp.RightNode)
	h.Height = 1 + utils.Max(height(h.LeftNode), height(h.RightNode))
	temp.Height = 1 + utils.Max(height(temp.LeftNode), height(temp.RightNode))

	return temp
}

// rotateRight rotates the tree around the node to the left
func rotateRight(h *Node) *Node {
	log.Println("rotateRight ", h.Key)
	temp := h.LeftNode
	h.LeftNode = temp.RightNode
	temp.RightNode = h

	temp.Size = 1 + size(temp.LeftNode) + size(temp.RightNode)
	h.Height = 1 + utils.Max(height(h.LeftNode), height(h.RightNode))
	temp.Height = 1 + utils.Max(height(temp.LeftNode), height(temp.RightNode))

	return temp
}

func (avlTree *AVLTree) Put(Key int, Value int) {
	avlTree.Root = put(avlTree.Root, Key, Value)
}

func put(root *Node, Key int, Value int) *Node {
	if root == nil {
		newNode := NewNode(Key, Value)
		return newNode
	}
	if Key < root.Key {
		root.LeftNode = put(root.LeftNode, Key, Value)
	} else if Key > root.Key {
		root.RightNode = put(root.RightNode, Key, Value)
	}

	root.Size = 1 + size(root.LeftNode) + size(root.RightNode)
	root.Height = 1 + utils.Max(height(root.LeftNode), height(root.RightNode))
	balance := nodeBalance(root)

	log.Println("balance ", balance, " Current Root = ", root.Key)

	// If node is not balances then perform rotation

	// Left - Left
	if balance > 1 && Key < root.LeftNode.Key {
		log.Println("Performing Left Left Operation -> Rotate Right")
		return rotateRight(root)
	}
	// Right - Right
	if balance < -1 && Key > root.RightNode.Key {
		log.Println("Performing Right Right Operation -> RotateLeft")
		return rotateLeft(root)
	}
	// Left - Right
	if balance > 1 && Key > root.LeftNode.Key {
		log.Println("Performing Left Right Rotation -> RotateRight")
		root.LeftNode = rotateLeft(root.LeftNode)
		return rotateRight(root)
	}
	// Right - Left
	if balance < -1 && Key < root.RightNode.Key {
		log.Println("Performing Right Left Rotation -> RotateLeft")
		root.RightNode = rotateRight(root.RightNode)
		return rotateLeft(root)
	}

	return root

}

// InOrderTraversal prints the tree in ascending order
// Left Root Right
func (avlTree *AVLTree) InOrderTraversal() {
	inorderTraversal(avlTree.Root)
}

func inorderTraversal(node *Node) {
	if node != nil {
		inorderTraversal(node.LeftNode)
		fmt.Print(node.Key, " ")
		inorderTraversal(node.RightNode)
	}
}

// LevelOrderTraversal or Breadth First Traversal starts at the tree root (or some arbitrary node of a graph, sometimes referred to as a 'search Key'[1]),
// and explores all of the neighbor nodes at the present depth prior to moving on to the nodes at the next depth level
func (avlTree *AVLTree) LevelOrderTraversal() {
	levelOrderTraversal(avlTree.Root)
	fmt.Println("")
}

func levelOrderTraversal(root *Node) {
	h := height(root)
	for i := 1; i <= h; i++ {
		printGivenLevel(root, i)
		fmt.Println("")
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
