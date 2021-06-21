package stack

import "log"

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	ShowStack()
	Size() int
	CheckTop(interface{}) bool
}

//Stack creates a slice with 0 len and 0 capacity
type interfaceImplStack struct {
	slice []interface{}
}

func NewStack() *interfaceImplStack {
	return &interfaceImplStack{make([]interface{}, 0)}
}

//Push pushes the value on the stack
func (s *interfaceImplStack) Push(val interface{}) {

	s.slice = append(s.slice, val)
}

//Peek returns the latest element on the Stack
func (s *interfaceImplStack) Peek() interface{} {
	return s.slice[len(s.slice)-1]
}

//Pop pops out the value onto the stack
func (s *interfaceImplStack) Pop() interface{} {
	retVal := s.slice[len(s.slice)-1]
	s.slice = s.slice[0 : len(s.slice)-1]
	return retVal
}

//ShowStack prints the elements in the stack
func (s *interfaceImplStack) ShowStack() {
	log.Println("----")
	i := len(s.slice) - 1
	for i >= 0 {
		log.Println(s.slice[i])
		i--
	}

	log.Println("----")
}

//Size returns the current size of the stack
func (s *interfaceImplStack) Size() int {
	return len(s.slice)
}

//CheckTop matches val to top of interfaceImplStack
func (s *interfaceImplStack) CheckTop(val interface{}) bool {
	topVal := s.Peek()
	if topVal == val {
		return true
	}
	return false
}
