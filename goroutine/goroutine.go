package main

import (
	//"fmt"
	"time"
	"fmt"
	"runtime"
)

// Coroutine ----> 协程
/*
	协程具有以下特点：
	1、是一种轻量级“线程”
	2、非抢占式多任务处理，由协程主动交出控制权
	3、编译器/解释器/虚拟机层面的多任务
	4、多个协程可能在一个或多个线程上运行
	5、操作系统有自己的调度器，golang中也有自己的调度器。协程由调度器控制
*/

func main() {

	var a [10]int

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				// 手动交出控制权
				runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
