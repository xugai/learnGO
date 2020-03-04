package main

import (
	"fmt"
	"time"
	)

func createWorker(i int) chan <- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Received %c from worker %d\n", <- c, i)
		}
	}()
	return c
}

func worker(i int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %d\n", i, n)
	}
}

func chanDemo() {
	var channels [10]chan <- int

	// 开启10个协程
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	for i := 0; i < 10; i++ {
		close(channels[i])
	}

	time.Sleep(time.Minute)
}

func bufferedChannel(i int) {
	c := make(chan int, 3)
	go worker(i, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Minute)
}


func main() {
	fmt.Println("Channel as the first-class citizen")
	chanDemo()
	fmt.Println("BufferedChannel and close channel with range")
	bufferedChannel(0)
}
