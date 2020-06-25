package binaryHeap

//BinaryHeap implements Heap for Priority Queue
type BinaryHeap struct {
	Maxpq []int
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
	var lastElem int
	if PQ.isEmpty() {
		PQ.Maxpq[1] = key
	} else {
		lastElem = len(PQ.Maxpq)
		PQ.Maxpq[lastElem] = key
	}
	PQ.swim(lastElem)
}

func (PQ *BinaryHeap) swim(k int) {

	for k > 1 && PQ.Maxpq[k] > PQ.Maxpq[k/2] {
		PQ.swap(k, k/2)
		k = k / 2
	}
}

func (PQ *BinaryHeap) swap(i int, j int) {
	temp := PQ.Maxpq[j]
	PQ.Maxpq[j] = PQ.Maxpq[i]
	PQ.Maxpq[i] = temp
}
