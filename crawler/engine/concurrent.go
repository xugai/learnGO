package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChannel chan Item
}

type Scheduler interface {
	WorkerReadyNotifier
	ConfigureWorkerChannel() chan Request
	Submit(Request)
	Run()
}

type WorkerReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ... Request) {
	//in := make(chan Request)
	out := make(chan ParseResult)
	c.Scheduler.Run()
	//c.Scheduler.ConfigureWorkerChannel(in)

	for i := 0; i < c.WorkerCount; i++ {
		createWorker(c.Scheduler.ConfigureWorkerChannel(), out, c.Scheduler)
	}

	for _, seed := range seeds {
		c.Scheduler.Submit(seed)
	}

	for {
		result := <- out
		for _, item := range result.Items {
			go doPersist(item, c.ItemChannel)
		}
		for _, request := range result.Requests {
			if ifDumplicate(request.Url) {
				continue
			}
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, n WorkerReadyNotifier) {
	go func() {
		for  {
			n.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrl = map[string]bool{}
func ifDumplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}

func doPersist(item Item, itemChannel chan Item) {
	itemChannel <- item
}
