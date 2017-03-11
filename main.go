package main

import (
	ex "./excercises"
)

func main() {

	//Problem: Detect cycles in a linked list
	println(ex.ContainsCycle(ex.MakeListWithCycles()))

	//Problem: Find the index of the first leftmost occurrence of k in XS.
	println(ex.FirstK([]int{5, 5, 3, 4, 5, 5, 5}, 5))

	//Problem: Check if delimeters are matched
	println(ex.MatchDelimeters("{ [ ( ] ) }"))

	//Problem: Reverse a linked list
	ex.ReverseLinkedList()

	//Problem: Reverse a stack using recursion
	ex.ReverseStackUsingRecursion()

	//Problem: Reverse characters in place
	ex.ReverseCharsInPlace()
}
