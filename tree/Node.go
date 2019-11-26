package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

// GO语言中的结构体是没有自身的构造函数的，它不像JAVA。如果我们要控制构造结构体的行为，那我们可以定义工厂方法
// GO语言中的所有变量，都是值传递的
func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}

// 为结构体定义方法
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if(node == nil){
		fmt.Println("Current node is nil, ignored it.")
		return
	}
	node.Value = value
}

