package main

import (
	".."
	"fmt"
)

/*
	记住，一个目录只能有一个package，main package中的main function是GO语言程序的执行入口
	使用其他package的东西，就要显示import进来
	结构体中，变量名首字母大写表示public，首字母小写表示private
	文件名跟package的名字可以不一样
*/

// 为已定义好的结构体扩展方法
// 后序遍历
type myTreeNode struct {
	*tree.Node  // Embeding 内嵌的方式
}

// 重载父类函数

func (myNode *myTreeNode) Traverse() {
	fmt.Println("This is a shadowed method.")
}

func (myNode *myTreeNode) PosterOrder() {
	if myNode == nil || myNode.Node == nil{
		return
	}

	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}

	left.PosterOrder()
	right.PosterOrder()
	myNode.Print()
}

func main() {
	//fmt.Println(root)
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)
	root.Right.Left.SetValue(4)
	root.SetValue(100)

	fmt.Print("root.Node.Traverse(): ")
	root.Node.Traverse()
	fmt.Println()
	fmt.Print("root.Traverse(): ")
	root.Traverse()
	fmt.Println()
	root.PosterOrder()
}
