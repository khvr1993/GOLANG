package main

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/queue/queue"
)

func main() {
	var q *queue.Queue = new(queue.Queue)
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)
	q.Enqueue(400)
	q.Enqueue(500)
	q.PrintQueue()
	log.Println("Size -- ", q.Size())
	q.Size()
	q.Dequeue()
	q.PrintQueue()
}
