package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("s = %v, len(s) = %v, cap(s) = %v\n", s, len(s), cap(s))
}

func main() {

	s := []int{}

	for i := 0; i < 100; i++ {
		//printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}		//通过定义数组的方式，创建slice
	printSlice(s1)

	s2 := make([]int, 16)		//通过调用make函数的方式创建slice，同时指定其长度大小
	printSlice(s2)

	s3 := make([]int, 10, 32)	//通过调用make函数的方式创建slice，同时指定其长度与容量大小
	printSlice(s3)

	fmt.Println("Copying from slice")

	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleteing slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("Popping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]

	fmt.Printf("front = %v, tail = %v\n", front, tail)
	printSlice(s2)
}
