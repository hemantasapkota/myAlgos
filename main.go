package main

import (
	"./bst"
	"./linkedlist"
	// "fmt"
)

func main() {

	list := linkedlist.NewNode("1")

	third := list.Add("2").Add("3")
	fourth := third.Add("4")
	fourth.Next = third

	println(list.ContainsCycle())
	println(bst.FirstK([]int{5, 5, 3, 4, 5, 5, 5}, 5))

	tree := bst.NewBSTree(5)

	println("Traversing trees")

	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(8)

	w := []int{}

	tree.InOrder(w)

	// fmt.Println(w)

}
