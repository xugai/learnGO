package main

import (
	"github.com/olivere/elastic/v7"
	"learnGO/crawler/engine"
	"learnGO/crawler/lianjia/parser"
	"learnGO/crawler/persist"
	"learnGO/crawler/scheduler"
)

// 183.48.244.8

var seed = engine.Request{
	Url: 		"https://cq.lianjia.com/zufang/",
	ParserFunc: parser.ParseCity,
}

func main() {
	// set sniff to false, because we run our elastic search single instance in docker.
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	engine := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 5,
		ItemChannel: persist.ItemSaver(client),
	}
	engine.Run(seed)
}
