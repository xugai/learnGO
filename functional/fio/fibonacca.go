package fio

import (
	"io"
	"fmt"
	"strings"
)

// 如何通过结合reader与scanner的方式，遍历输出一定数量的斐波那契数?

func Fibonacca() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}



type intGen func() int

// golang中很特别的一点是，函数也可以当做对象来看待，实现接口的方法
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	// 将整型转换成字符串
	s := fmt.Sprintf("%d\n", next)
	// 委托其他类型的Reader，进行Read操作
	return strings.NewReader(s).Read(p)
}
