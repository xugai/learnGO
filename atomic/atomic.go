package main

import (
	"fmt"
	"time"
	"sync"
)

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (ato *atomicInt) increment() {
	fmt.Println("increment atomicInt.")
	ato.lock.Lock()
	defer ato.lock.Unlock()
	ato.value++
}

func (ato *atomicInt) get() int{
	ato.lock.Lock()
	defer ato.lock.Unlock()
	return ato.value
}


func main() {
	ato := atomicInt{
		value: 0,
		lock: sync.Mutex{},
	}
	ato.increment()
	go func() {
		ato.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(ato.get())
}
