package main

import (
	"learnGO/crawler/engine"
	"learnGO/crawler/lianjia/parser"
	"learnGO/crawler/scheduler"
)

var seed = engine.Request{
	Url: 		"https://www.lianjia.com/city/",
	ParserFunc: parser.ParseCityList,
}

func main() {
	engine := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	engine.Run(seed)
}
// 6525   6535    6536    6469