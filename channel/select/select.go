package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generator() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

type worker struct {
	in chan int
	done chan bool
	//wg *sync.WaitGroup
}

func createWorker(i int) worker {
	w := worker{
		in: make(chan int),
		done: make(chan bool),
	}
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Printf("Received %d from worker %d\n", <- w.in, i)
			//w.done <- true
		}
	}()
	return w
}

// select channel 的特性
/*
	只要是nil channel，在select作用域里面它就是阻塞的，不会被select到
*/
func main() {

	var c1, c2 = generator(), generator()
	//c1, c2 = generator(), generator()
	w := createWorker(0)
	var values []int
	var activeNum = 0
	//hasValue := false
	n := 0
	tm := time.After(10 * time.Second)	// 规定10s后结束任务
	// 每隔1s定时统计暂存起来的number个数(定时任务)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan int
		if len(values) > 0 {
			activeWorker = w.in
			activeNum = values[0]
		}
		select {
		case n = <- c1:
			values = append(values, n)
		case n = <- c2:
			values = append(values, n)
		case activeWorker <- activeNum:
			values = values[1:]
		case <- time.After(800 * time.Millisecond):
			fmt.Println("Time out!")	// 说明生产者(generator)的生产时间慢于0.8s
		case <- tick:
			fmt.Println("length of values:", len(values))
		case <- tm:
			fmt.Println("Good bye!")
			return
		}
	}

}
