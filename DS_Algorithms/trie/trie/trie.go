package trie

import "log"

type trieNode struct {
	EndOfString bool
	Children    map[string]*trieNode
}

type Trie struct {
	Root *trieNode
}

func newTrieNode() *trieNode {
	emptyMap := make(map[string]*trieNode)
	return &trieNode{false, emptyMap}
}

// NewTrie initialises Trie DS
func NewTrie() *Trie {
	trie := new(Trie)
	trie.Root = newTrieNode()

	return trie
}

// Insert will create a Trie datastructure for all the strings given
func (t *Trie) Insert(key string) {
	log.Println("Insert")
	/*
		Get the root of the trie
		For the characters in the string check if the character is present in the trie node. If it is present then move to next node
		else create a new trie node and for the current node assign the current character
	*/
	node := t.Root
	log.Println("Root Node", node)

	for i := 0; i < len(key); i++ {
		index := string(key[i])
		_, ok := node.Children[index]
		if !ok {
			node.Children[index] = newTrieNode()
			log.Println("node ", node, " index ", index)
		}
		node = node.Children[index]

	}
	node.EndOfString = true
}

// Search will return if the string is present in the list of strings
func (t *Trie) Search(key string) bool {
	node := t.Root

	for i := 0; i < len(key); i++ {
		index := string(key[i])
		if _, ok := node.Children[index]; !ok {
			return false
		}
		node = node.Children[index]
	}
	return node.EndOfString
}
