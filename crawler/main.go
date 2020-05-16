package main

import (
	"learnGO/crawler/engine"
	"learnGO/crawler/lianjia/parser"
	"learnGO/crawler/persist"
	"learnGO/crawler/scheduler"
)

var seed = engine.Request{
	Url: 		"https://www.lianjia.com/city/",
	ParserFunc: parser.ParseCityList,
}

func main() {
	engine := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChannel: persist.ItemSaver(),
	}
	engine.Run(seed)
}
