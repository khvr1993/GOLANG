package Bag

import "log"

type Bag struct {
	data map[interface{}]int
}

func NewBag() *Bag {
	b := Bag{data: nil}
	b.data = make(map[interface{}]int) //Initialize the map
	return &b
}

func (b *Bag) AddItem(val interface{}) {
	_, ok := b.data[val]
	if ok {
		b.data[val] += 1
	} else {
		b.data[val] = 1
	}
}

func (b *Bag) Size() int {
	return len(b.data)
}

func (b *Bag) IsEmpty() bool {
	if len(b.data) == 0 {
		return true
	}
	return false
}

func (b *Bag) DisplayContents() {
	for k, v := range b.data {
		log.Println("key ", k, " Value ", v)
	}
}
