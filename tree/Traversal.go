package tree

func (node *Node) Traverse() {
	if node != nil {
		node.Left.Traverse()
		node.Print()
		node.Right.Traverse()
	}
}
