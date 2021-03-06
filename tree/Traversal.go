package tree

import "fmt"

func (node *Node) Traverse() {
	//if node != nil {
	//	node.Left.Traverse()
	//	node.Print()
	//	node.Right.Traverse()
	//}
	node.TraverseFunc(func(node *Node) {
		fmt.Print(node.Value, " ")
	})
}

func (node *Node) TraverseFunc(f func(node *Node)) {
	if node != nil {
		node.Left.TraverseFunc(f)
		f(node)
		node.Right.TraverseFunc(f)
	}
}

func (node *Node) TraverseFuncWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()

	return out
}


