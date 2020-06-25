package binaryHeap

import "fmt"

//BinaryHeap implements Heap for Priority Queue
type BinaryHeap struct {
	Maxpq []int
	N     int
}

func (PQ *BinaryHeap) isEmpty() bool {
	len := len(PQ.Maxpq)
	if len > 0 {
		return false
	}
	return true
}

//Size returns the size of the heap;
func (PQ *BinaryHeap) Size() int {
	return len(PQ.Maxpq)
}

//Insert inserts into the HEAP and performs swim
func (PQ *BinaryHeap) Insert(key int) {
	if PQ.isEmpty() {
		PQ.Maxpq = append(PQ.Maxpq, 0)
		PQ.Maxpq = append(PQ.Maxpq, key)
		PQ.N = 1
	} else {
		PQ.N++
		PQ.Maxpq = append(PQ.Maxpq, key)
	}
	PQ.swim(PQ.N)
}

func (PQ *BinaryHeap) swim(k int) {

	for k > 1 && PQ.Maxpq[k] > PQ.Maxpq[k/2] {
		PQ.swap(k, k/2)
		k = k / 2
	}
}

func (PQ *BinaryHeap) sink(k int) {
	for 2*k <= PQ.N {
		j := 2 * k
		if j < PQ.N && PQ.Maxpq[j] < PQ.Maxpq[j+1] {
			j++
		}
		if PQ.Maxpq[j] < PQ.Maxpq[k] {
			break
		}
		PQ.swap(k, j)
		k = j
	}
}

func (PQ *BinaryHeap) swap(i int, j int) {
	temp := PQ.Maxpq[j]
	PQ.Maxpq[j] = PQ.Maxpq[i]
	PQ.Maxpq[i] = temp
}

//GetMax deletes the maximum element
func (PQ *BinaryHeap) GetMax() int {
	maxElem := PQ.Maxpq[1]
	PQ.swap(1, PQ.N)
	PQ.Maxpq = PQ.Maxpq[:PQ.N]
	PQ.N--
	PQ.sink(1)
	return maxElem
}

//ShowHeap prints the heap
func (PQ *BinaryHeap) ShowHeap() {
	i := 1
	k := 1
	for i <= PQ.N {
		fmt.Printf("%d\t\t", PQ.Maxpq[i])
		if i >= k {
			k = 2*k + 1
			fmt.Printf("\n")
		}
		i++
	}
	fmt.Printf("\n")
}
