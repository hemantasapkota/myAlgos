package main

import (
	ex "./excercises"
)

func main() {

	println(ex.ContainsCycle(ex.MakeListWithCycles()))

	println(ex.FirstK([]int{5, 5, 3, 4, 5, 5, 5}, 5))

	// Stack based question
	println(ex.MatchDelimeters("{ [ ( ] ) }"))

	ex.ReverseLinkedList()

	ex.ReverseStackUsingRecursion()

	ex.ReverseCharsInPlace()
}
