package bst

type BSTree struct {
	Key   int
	Left  *BSTree
	Right *BSTree
}

func NewBSTree(key int) *BSTree {
	return &BSTree{Key: key}
}

func (t *BSTree) Insert(key int) {
	if t == nil {
		return
	}

	switch {
	case t.Key == key:
		return

	case t.Key < key:
		if t.Left == nil {
			t.Left = &BSTree{Key: key}
			return
		}
		t.Left.Insert(key)

	case t.Key > key:
		if t.Right == nil {
			t.Right = &BSTree{Key: key}
		}
		t.Right.Insert(key)

	default:
		return
	}

}

func (t *BSTree) InOrder(w []int) {
	if t != nil {
		t.Left.InOrder(w)
		println(t.Key)
		t.Right.InOrder(w)
	}
}
