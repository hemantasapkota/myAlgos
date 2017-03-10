package excercises

import (
	"../linkedlist"
)

func MakeListWithCycles() *linkedlist.Node {
	list := linkedlist.NewNode("1")
	third := list.Add("2").Add("3")
	fourth := third.Add("4")
	fourth.Next = third
	return list
}

func ContainsCycle(node *linkedlist.Node) bool {
	slowRunner := node
	fastRunner := node
	for fastRunner != nil && fastRunner.Next != nil {
		slowRunner = slowRunner.Next
		fastRunner = fastRunner.Next.Next
		if slowRunner == fastRunner {
			return true
		}
	}
	return false
}

func ReverseLinkedList() {
	reverse := func(node *linkedlist.Node) *linkedlist.Node {
		// Init our three pointers
		var head, next, cursor *linkedlist.Node = node, nil, nil

		for head != nil {
			next = head.Next
			head.Next = cursor
			cursor = head
			head = next
		}
		return cursor
	}

	// 1 2 3 4 5
	// 5 4 3 2 1

	list := linkedlist.NewNode("1")
	list.Add("2").Add("3").Add("4")

	list.Traverse()

	rev := reverse(list)
	rev.Traverse()
}
