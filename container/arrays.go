package main

import "fmt"

// Go语言中，数组作为参数传递的话，默认情况它也是属于值拷贝传递，而不是引用传递
func printArray(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {

	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	maxIndex := -1
	maxValue := -1
	for i, v := range arr3 {
		if maxValue < v {
			maxValue = v
			maxIndex = i
		}
	}
	fmt.Println(maxIndex, maxValue)

	fmt.Println("Call func printArray: ")
	printArray(arr1[:])
	printArray(arr3[:])
	fmt.Println(arr1, arr3)
}
