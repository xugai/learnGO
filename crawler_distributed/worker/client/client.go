package client

import (
	"learnGO/crawler/engine"
	"learnGO/crawler_distributed/config"
	"learnGO/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChannel chan *rpc.Client) (engine.Processor, error) {
	return func(engineRequest engine.Request) (engine.ParseResult, error) {
		sRequest := worker.SerializeRequest(engineRequest)
		sResult := worker.ParseResult{}
		// 多线程并发场景中，通过go原生的chan，即可实现并发安全，而不用我们显式地通过锁来达到同步控制。单单从语法层面上来说，这样更加简洁了
		c := <- clientChannel
		err := c.Call(config.CRAWLER_SERVICE, sRequest, &sResult)
		if err != nil {
			return engine.ParseResult{}, nil
		}
		return worker.DeserializeParseResult(sResult), nil
	}, nil
}
