package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureWorkerChannel(in chan Request)
}

func (c *ConcurrentEngine) Run(seeds ... Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	c.Scheduler.ConfigureWorkerChannel(in)

	for i := 0; i < c.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, seed := range seeds {
		c.Scheduler.Submit(seed)
	}
	itemCount := 0
	for {
		result := <- out
		for _, item := range result.Items {
			log.Printf("Got #%d item %v\n", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for  {
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
