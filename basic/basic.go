package main

import (
	"fmt"
	"math/cmplx"
	"math"
)
/*
	go 语言中，没有所谓的全局变量，其不存在全局变量这一说，只有包内变量
	因此定义在函数外部的变量，就属于包内变量。这样就避免了全局变量中可能
	出现的重名变量而相互覆盖的情况
*/

var	(
	aa = 33
	bb = "bb"
 	cc = false
)


// 定义变量，编译器自动赋初值
func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a, s)
	fmt.Printf("%d, %q\n", a, s)
}

// 定义变量，我们为其赋初值
func variableInitialValue() {
	var a, b int = 1, 2
	a = 3
	var s string = "hello, berio"
	fmt.Println(a, b, s)
}

// 定义变量，但不声明类型，编译器会自动识别
func variableTypeDeduction() {
	var a, b, c, d = 1, 2.2, true, "abc"
	fmt.Println(a, b, c, d)
}

// 定义变量，但不使用 var 关键字，使用 :=，但这种写法只能定义在函数内，不能定义在函数外
func variableShorter() {
	a, b, c, d := 1, 2, true, "abc"
	fmt.Println(a, b, c, d)
}

// 欧拉公式，复数类型的验证
// cmplx.Pow()，求指数的结果，第一个入参是底数，第二个入参是指数
// cmplx.Exp()，求底数为e的指数的值
func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
}

// 不同于其他语言，在go语言中，只有强制类型转换，没有隐式类型转换
func triangle() {
	a, b := 3, 4
	c := 0
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func calcTriangle(a int, b int) int {

	c := 0
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}

func consts() {
	const (
		cpp = iota
		_
		nodejs
		python
		golang
		javascript
	)
	//const (
	//	a = 3
	//	b = 4
	//)
	const (
		b = 1 << (10 * iota)  	// 1 * 2的0次方
		kb			// 1 * 2的10次方
		mb			// 1 * 2的20次方
		gb			// 1 * 2的30次方
		tb 			// 1 * 2的40次方
		pb			// 1 * 2的50次方
	)
	fmt.Println(cpp, javascript, nodejs, python, golang)
	//c := 1
	//c = int(math.Sqrt(a * a + b * b))
	//fmt.Println(c)
	fmt.Println(b, kb, mb, gb, int64(tb), int64(pb))
}

func main() {
	//fmt.Println("Hello World!")
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShorter()
	//fmt.Println(aa, bb, cc)
	//euler()
	//triangle()
	consts()
}
