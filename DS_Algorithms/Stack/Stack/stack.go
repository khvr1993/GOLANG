package stack

import "log"

//Stack creates a slice with 0 len and 0 capacity
type Stack struct {
	slice []int
}

//Push pushes the value on the stack
func (s *Stack) Push(val int) {
	s.slice = append(s.slice, val)
}

//Peek returns the latest element on the Stack
func (s *Stack) Peek() int {
	return s.slice[len(s.slice)-1]
}

//Pop pops out the value onto the stack
func (s *Stack) Pop() int {
	retVal := s.slice[len(s.slice)-1]
	s.slice = s.slice[0 : len(s.slice)-1]
	return retVal
}

//ShowStack prints the elements in the stack
func (s *Stack) ShowStack() {
	log.Println("----")
	i := len(s.slice) - 1
	for i >= 0 {
		log.Println(s.slice[i])
		i--
	}

	log.Println("----")
}

//Size returns the current size of the stack
func (s *Stack) Size() int {
	return len(s.slice)
}

//CheckTop matches val to top of Stack
func (s *Stack) CheckTop(val int) bool {
	topVal := s.Peek()
	if topVal == val {
		return true
	}
	return false
}
