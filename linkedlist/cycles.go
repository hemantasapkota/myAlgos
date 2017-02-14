package linkedlist

func (n *Node) ContainsCycle() bool {
	slowRunner := n
	fastRunner := n

	for fastRunner != nil && fastRunner.Next != nil {
		slowRunner = slowRunner.Next
		fastRunner = fastRunner.Next.Next
		if slowRunner == fastRunner {
			return true
		}
	}

	return false

}
