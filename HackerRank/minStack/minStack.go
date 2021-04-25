package minStack

/*
https://leetcode.com/problems/min-stack/
*/
type MinStack struct {
	value   []int
	minHeap []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	minStack := new(MinStack)
	return *minStack
}

func (this *MinStack) Push(val int) {
	this.value = append(this.value, val)
	this.minHeap = append(this.minHeap, val)
	this.swim()
}

func (this *MinStack) Pop() {

	if len(this.value) == 0 {
		return
	}

	if this.value[len(this.value)-1] == this.minHeap[0] {
		this.minHeap = nil
		// removing the minimum element. We need to rebuild the minHeap
		for i := 0; i < len(this.value)-1; i++ {
			this.minHeap = append(this.minHeap, this.value[i])
			this.swim()
		}
	}
	this.value = this.value[0 : len(this.value)-1]
}

func (this *MinStack) Top() int {
	var top int
	if len(this.value) == 0 {
		return top
	}

	return this.value[len(this.value)-1]
}

func (this *MinStack) GetMin() int {
	var min int
	if len(this.value) == 0 {
		return min
	}
	return this.minHeap[0]
}

func (this *MinStack) swim() {
	k := len(this.minHeap) - 1
	//For the records where the value at root is not less than
	//the leaf push the value of the child to the root
	for k > 0 && this.minHeap[k] < this.minHeap[k/2] {
		//swap the element to the correct position
		temp := this.minHeap[k]
		this.minHeap[k] = this.minHeap[k/2]
		this.minHeap[k/2] = temp

		k = k / 2
	}
}

/**
* Your MinStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(val);
* obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.GetMin();
 */
