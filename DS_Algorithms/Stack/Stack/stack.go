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

type IntStack struct {
	slice []int
}

type StrStack struct {
	slice []string
}

func NewIntStack() *IntStack {
	return &IntStack{make([]int, 0)}
}

func NewStrStack() *StrStack {
	return &StrStack{make([]string, 0)}
}

//Push pushes the value on the stack
func (s *IntStack) Push(val interface{}) {
	s.slice = append(s.slice, val.(int))
}

//Peek returns the latest element on the Stack
func (s *IntStack) Peek() interface{} {
	return s.slice[len(s.slice)-1]
}

//Pop pops out the value onto the stack
func (s *IntStack) Pop() interface{} {
	retVal := s.slice[len(s.slice)-1]
	s.slice = s.slice[0 : len(s.slice)-1]
	return retVal
}

//ShowStack prints the elements in the stack
func (s *IntStack) ShowStack() {
	log.Println("----")
	i := len(s.slice) - 1
	for i >= 0 {
		log.Println(s.slice[i])
		i--
	}

	log.Println("----")
}

//Size returns the current size of the stack
func (s *IntStack) Size() int {
	return len(s.slice)
}

//CheckTop matches val to top of IntStack
func (s *IntStack) CheckTop(val interface{}) bool {
	topVal := s.Peek()
	if topVal == val {
		return true
	}
	return false
}

//Push pushes the value on the stack
func (s *StrStack) Push(val interface{}) {
	s.slice = append(s.slice, val.(string))
}

//Peek returns the latest element on the Stack
func (s *StrStack) Peek() interface{} {
	return s.slice[len(s.slice)-1]
}

//Pop pops out the value onto the stack
func (s *StrStack) Pop() interface{} {
	retVal := s.slice[len(s.slice)-1]
	s.slice = s.slice[0 : len(s.slice)-1]
	return retVal
}

//ShowStack prints the elements in the stack
func (s *StrStack) ShowStack() {
	log.Println("----")
	i := len(s.slice) - 1
	for i >= 0 {
		log.Println(s.slice[i])
		i--
	}

	log.Println("----")
}

//Size returns the current size of the stack
func (s *StrStack) Size() int {
	return len(s.slice)
}

//CheckTop matches val to top of Stack
func (s *StrStack) CheckTop(val interface{}) bool {
	topVal := s.Peek()
	if topVal == val {
		return true
	}
	return false
}
