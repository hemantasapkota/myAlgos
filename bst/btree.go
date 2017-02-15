// Know how to implment a binary tree and basic operations on it

//        A
//    B       C
//          D    E

// Review:

// Height of node: The number of edges from the node to the leaf
// Depth of a tree: Height of the root
// Traversal: BFS ( InOrder, PreOrder, PostOrder ), DFS ( Level Order )
// InOrder: Left, Parent, Right
// PreOrder: Parent, Left, Right
// PostOrder: Left, Right, Parent

package main

import (
	"fmt"
)

type BTree struct {
	Data  interface{}
	left  *BTree
	right *BTree
}

func (t *BTree) Height() int {
	// Base Condition
	// Height of the left tree
	// Height of the right tree
	// Compare height and return the greatest
	if t == nil {
		return -1
	}
	lh := 1 + t.left.Height()
	rh := 1 + t.right.Height()
	if lh > rh {
		return lh
	}
	return rh
}

func (t *BTree) Preorder(w []interface{}) []interface{} {
	if t != nil {
		w = append(w, t.Data)
		w = t.left.Preorder(w)
		w = t.right.Preorder(w)
	}
	return w
}

func (t *BTree) Inorder(w []interface{}) []interface{} {
	if t != nil {
		w = t.left.Inorder(w)
		w = append(w, t.Data)
		w = t.right.Inorder(w)
	}
	return w
}

func (t *BTree) Postorder(w []interface{}) []interface{} {
	if t != nil {
		w = t.left.Postorder(w)
		w = t.right.Postorder(w)
		w = append(w, t.Data)
	}
	return w
}

func (t *BTree) IsBalanced() bool {
	return false
}

func main_() {
	node := &BTree{"A", &BTree{"B", nil, nil}, &BTree{"C", &BTree{"D", nil, nil}, &BTree{"E", nil, nil}}}

	height := node.Height()
	fmt.Printf("Height of the tree: %d", height)
	println()

	list := make([]interface{}, 0)
	list = node.Preorder(list)
	fmt.Printf("%v", list)
	println()

	list = make([]interface{}, 0)
	list = node.Inorder(list)
	fmt.Printf("%v", list)

	println()
	list = make([]interface{}, 0)
	list = node.Postorder(list)
	fmt.Printf("%v", list)
}
