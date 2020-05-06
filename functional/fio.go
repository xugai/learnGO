package main

import (
	"bufio"
	"fmt"
	"io"
	"learnGO/functional/fio"
)

func printIntWithReader(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println("loop times.")
		fmt.Println(scanner.Text())
	}

}

func main() {
	// 斐波那契数列
	f := fio.Fibonacca()
	printIntWithReader(f)

}