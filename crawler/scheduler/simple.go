package scheduler

import "learnGO/crawler/engine"

type SimpleScheduler struct {
	workerChannel chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	//pass request to worker to parse
	//async to process request, which will not block scheduler submit.
	go func() {
		s.workerChannel <- request
	}()
	//s.workerChannel <- request
}

func (s *SimpleScheduler) ConfigureWorkerChannel() chan engine.Request{
	return s.workerChannel
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChannel = make(chan engine.Request)
}


