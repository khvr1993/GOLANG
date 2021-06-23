package Bag

import "log"

type Bag struct {
	data map[interface{}]int
}

/*
	Creates a new Bag structure and returns the pointer
*/
func NewBag() *Bag {
	b := Bag{data: nil}
	b.data = make(map[interface{}]int) //Initialize the map
	return &b
}

// AddItem adds the item to the Bag
func (b *Bag) AddItem(val interface{}) {
	_, ok := b.data[val]
	if ok {
		b.data[val] += 1
	} else {
		b.data[val] = 1
	}
}

// Size returns the number of contents in the Bag
func (b *Bag) Size() int {
	return len(b.data)
}

// IsEmpty Checks if the Bag is Empty and returns true
func (b *Bag) IsEmpty() bool {
	if len(b.data) == 0 {
		return true
	}
	return false
}

// DisplayContents Displays count of the unique items and the items
func (b *Bag) DisplayContents() {
	for k, v := range b.data {
		log.Println("key ", k, " Value ", v)
	}
}

// ShowBag Displays the unique items of the bag
func (b *Bag) ShowBag() {
	for k := range b.data {
		log.Println(k)
	}
}

// Iterable returns the map of the Bag so that the calling function can iterate over
func (b *Bag) Iterable() map[interface{}]int {
	return b.data
}
