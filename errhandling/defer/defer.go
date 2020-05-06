package main

import (
	"bufio"
	"errors"
	"fmt"
	"learnGO/functional/fio"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	panic("error occured!")
	fmt.Println(4)
}

type myError func() string

func NewError() myError {
	 return func() string {
		 return "This is myError by func"
	 }
}

func (me myError) Error() string {
	return "This is myError!!!"
}

func writeFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE|os.O_TRUNC, 0666)
	err = errors.New("This is a Customer Error!")
	if err != nil {
		// 以此方式来判断当前error是哪种类型的error
		if pathError, ok := err.(*os.PathError); !ok {
			myerr := NewError()
			panic(myerr)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	// 等当前方法执行退出的时候，调用下面的方法，将缓存中的内容写入到文件里
	defer writer.Flush()


	f := fio.Fibonacca()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}



func main() {
	//tryDefer()
	writeFile("fib.txt")
}
