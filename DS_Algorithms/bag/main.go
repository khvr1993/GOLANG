package main

import (
	"github.com/khvr1993/GOLANG/DS_Algorithms/bag/Bag"
	"log"
)

func main() {
	log.Println("Testing Bag ")
	bag := Bag.NewBag()
	bag.AddItem("s")
	bag.AddItem("a")
	bag.AddItem("a")
	log.Println("Bag Size ", bag.Size())
	bag.DisplayContents()
}
