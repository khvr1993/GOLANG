package main

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/linkedQueueOfElements/queue"
)

func main() {
	log.Println("linkedQueueOfElements")
	var q = new(queue.Queue)
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)
	q.PrintQueue()

}
