package stack

import "log"

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	ShowStack()
	Size() int
	CheckTop(interface{}) bool
	IsEmpty() bool
}

//InterfaceImplStack creates a slice with 0 len and 0 capacity
type InterfaceImplStack struct {
	slice []interface{}
}

func NewStack() *InterfaceImplStack {
	return &InterfaceImplStack{make([]interface{}, 0)}
}

//Push pushes the value on the stack
func (s *InterfaceImplStack) Push(val interface{}) {

	s.slice = append(s.slice, val)
}

//Peek returns the latest element on the Stack
func (s *InterfaceImplStack) Peek() interface{} {
	return s.slice[len(s.slice)-1]
}

//Pop pops out the value onto the stack
func (s *InterfaceImplStack) Pop() interface{} {
	if len(s.slice) == 0 {
		return nil
	}
	retVal := s.slice[len(s.slice)-1]
	s.slice = s.slice[0 : len(s.slice)-1]
	return retVal
}

//ShowStack prints the elements in the stack
func (s *InterfaceImplStack) ShowStack() {
	log.Println("----")
	i := len(s.slice) - 1
	for i >= 0 {
		log.Println(s.slice[i])
		i--
	}

	log.Println("----")
}

//Size returns the current size of the stack
func (s *InterfaceImplStack) Size() int {
	return len(s.slice)
}

//CheckTop matches val to top of InterfaceImplStack
func (s *InterfaceImplStack) CheckTop(val interface{}) bool {
	topVal := s.Peek()
	if topVal == val {
		return true
	}
	return false
}

//checks if the stack is empty
func (s *InterfaceImplStack) IsEmpty() bool {
	if len(s.slice) > 0 {
		return false
	} else {
		return true
	}
}
