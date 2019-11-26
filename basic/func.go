package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

// 单个返回值的函数
func eval(a, b int, ops string) (int, error) {
	switch ops {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("Unsupported operation symbol: %s", ops)
	}
}

// 具有多个返回值的函数
func div(a, b int) (int ,int) {
	return a / b, a % b
}

// GO语言特点之一函数式编程，可将函数作为一个参数使用
func assign(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Call function %s with args: (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func magic(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	//if result, err := eval(12, 3, "0"); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(result)
	//}
	//fmt.Println(div(12, 5))
	//fmt.Println(assign(func (a, b int) int {
	//	return int(math.Pow(float64(a), float64(b)))
	//}, 3, 3))

	//fmt.Println(magic(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
