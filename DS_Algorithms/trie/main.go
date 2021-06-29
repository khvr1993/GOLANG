package main

import (
	trie2 "github.com/khvr1993/GOLANG/DS_Algorithms/trie/trie"
	"log"
)

func main() {
	log.Println("Testing trie datastructure")
	trie := trie2.NewTrie()
	trie.Insert("Harsha")
	log.Println(trie.Root)
	trie.Insert("Hars")
	log.Println(trie.Root)

	log.Println(trie.Search("Hars"))
}
