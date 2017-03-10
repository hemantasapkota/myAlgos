package main

import (
	ex "./excercises"
	"fmt"
)

func main() {

	println(ex.ContainsCycle(ex.MakeListWithCycles()))

	println(ex.FirstK([]int{5, 5, 3, 4, 5, 5, 5}, 5))

	// Stack based question
	println(ex.MatchDelimeters("{ [ ( ] ) }"))

	ex.ReverseLinkedList()

	// TODO: Review this
	s := ex.MakeStack()
	ex.ReverseStackUsingRecursion(s)

	val := s.Pop()
	for val != nil {
		fmt.Println(val)
		val = s.Pop()
	}

}
