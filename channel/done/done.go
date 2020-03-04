package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	//done chan bool
	wg *sync.WaitGroup
}

func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		//done: make(chan bool, 2),
		wg: wg,
	}
	go func() {
		for {
			fmt.Printf("Received %c from worker %d\n", <- w.in, i)
			w.wg.Done()
		}
	}()
	return w
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	// 开启10个协程
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}

	// wait all of them done

	//for _, worker := range workers {
	//	<- worker.done
	//	<- worker.done
	//}

	wg.Wait()
}



func main() {
	fmt.Println("Channel as the first-class citizen")
	chanDemo()
}
