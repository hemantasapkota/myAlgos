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
