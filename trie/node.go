package trie

type Node struct {
	Children map[rune]*Node
}

func (n *Node) MakeChildNode(character rune) {
	n.Children[character] = &Node{}
}

func (n *Node) GetChildNode(character rune) *Node {
	return n.Children[character]
}

func (n *Node) HasChildNode(character rune) bool {
	_, ok := n.Children[character]
	return ok
}

type Trie struct {
	Root *Node
}
