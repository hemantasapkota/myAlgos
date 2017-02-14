package linkedlist

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{Data: data, Next: nil}
}

func (n *Node) Add(data interface{}) *Node {
	newNode := &Node{Data: data, Next: nil}
	n.Next = newNode
	return newNode
}

func (n *Node) Traverse() {
	node := n
	for node != nil {
		fmt.Printf("%v", node.Data)
		node = node.Next
	}
}
