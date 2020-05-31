package queue

import "fmt"

//Queue create a structure of slice of int
type Queue struct {
	slice []int
}

//Enqueue pushes the value into Queue
func (q *Queue) Enqueue(val int) {
	q.slice = append(q.slice, val)
}

//Peek returns the latest entry to be dequeued
func (q *Queue) Peek() int {
	return q.slice[0]
}

//Dequeue removes the element from the Queue
func (q *Queue) Dequeue() int {
	retVal := q.slice[0]
	q.slice = q.slice[1:]
	return retVal
}

//Size returns the size of the queue
func (q *Queue) Size() int {
	return len(q.slice)
}

//PrintQueue prints the queue
func (q *Queue) PrintQueue() {
	fmt.Printf(" --- ")
	var i int
	for i < len(q.slice) {
		fmt.Printf("%v ", q.slice[i])
		i++
	}
	fmt.Printf(" ---\n")
}
