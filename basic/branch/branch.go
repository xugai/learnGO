package main

import (
	"io/ioutil"
	"fmt"
	"strconv"
	"os"
	"bufio"
	"io"
	"strings"
)

// go语言中的switch语句是没有break的，进入某个case执行完后，它会自动跳出switch代码块
// 与此同时，switch语句后面可以不跟表达式
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

// for循环，可省略循环变量的初始化条件，保留结束条件与递增条件
func convertToBin(number int) string {
	result := ""
	for ; number > 0; number /= 2 {
		lsb := number % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
// for循环，省略了初始条件与递增条件，只保留结束条件
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContent(file)
}
func printFileContent(reader io.Reader) {

	scanner := bufio.NewScanner(reader)

	// print content with each line
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
// for循环，省略所有的条件，此时也就是无限循环了，效果等同于while(true)或者while(1)
func forever() {
	for {
		fmt.Println("hello!!!")
	}
}
func main() {
	const filename = "basic/abc.txt"
	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}


	//fmt.Println(eval(3, 6, "*"))
	fmt.Println(grade(59), grade(69), grade(79), grade(89), grade(99))

	fmt.Println(convertToBin(3))
	fmt.Println(convertToBin(13))
	printFile("basic/abc.txt")
	//forever()

	s := `abccc
	1234

	kkk
	bbb`

	printFileContent(strings.NewReader(s))
}
