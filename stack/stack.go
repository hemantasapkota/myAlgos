package stack

import (
	"log"
)

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	Value interface{}
	Next  *Element
}

func New() *Stack {
	log.Printf("Creating an empty stack")
	return &Stack{}
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{Value: value, Next: s.top}
	s.size++
	// log.Printf("Pushing %v, Updated Size: %d", value, s.size)
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		value := s.top.Value
		s.top = s.top.Next
		s.size--
		// log.Printf("Popping %v, Updated Size: %d", value, s.size)
		return value
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if s.size > 0 {
		value := s.top.Value
		// log.Printf("Peeking: %v", value)
		return value
	}
	return nil
}
