package main

import (
	".."
	"fmt"
)
/*
	有时候我们想对一个类型新增方法或者属性，也就是扩展现有类型，这有点像是JAVA里面的继承。
	一般来说，扩展类型的思路有三种：通过组合、采用别名方式、内嵌方式
*/

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	//q.Push("string")
	//q.Push(true)
	//q.Push(7.4)
	//fmt.Println(q.Pop())
	//fmt.Println(q.Pop())
	//fmt.Println(q.Pop())
}
