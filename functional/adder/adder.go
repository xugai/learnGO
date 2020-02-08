package main

import "fmt"

/*
	正统的函数式编程有以下两点要求：
	1、函数体内不能有状态
	2、函数只能携带一个参数
	但幸运的是，golang中的函数式编程可以不用以此做得这么严格
*/

func adder() func(value int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}

type iAdder func(value int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return v + base, adder2(v + base)
	}
}

func main(){
	a := adder2(0)
	s := 0
	for i := 0; i < 10; i++ {
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}