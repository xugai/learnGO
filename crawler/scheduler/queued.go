package scheduler

import (
	"learnGO/crawler/engine"
)

type QueuedScheduler struct {
	RequestChannel chan engine.Request
	WorkerChannel chan chan engine.Request
}

func (q *QueuedScheduler) Submit(request engine.Request) {
	q.RequestChannel <- request
}

func (q *QueuedScheduler) ConfigureWorkerChannel() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) WorkerReady(readyChannel chan engine.Request) {
	q.WorkerChannel <- readyChannel
}

func (q *QueuedScheduler) Run() {
	q.RequestChannel = make(chan engine.Request)
	q.WorkerChannel = make(chan chan engine.Request)
	var requestQueue []  engine.Request
	var workerChannelQueue [] chan engine.Request
	go func() {
		for {
			var currentRequest engine.Request
			var currentWorkerChannel chan engine.Request
			if len(requestQueue) > 0 && len(workerChannelQueue) > 0 {
				currentRequest = requestQueue[0]
				currentWorkerChannel = workerChannelQueue[0]
			}
			select {
			case request := <- q.RequestChannel:
				requestQueue = append(requestQueue, request)
			case workerChannel := <- q.WorkerChannel:
				workerChannelQueue = append(workerChannelQueue, workerChannel)
			case currentWorkerChannel <- currentRequest:
				workerChannelQueue = workerChannelQueue[1:]
				requestQueue = requestQueue[1:]
			}
		}
	}()
}

