package main

import (
	"learnGO/crawler/engine"
	"learnGO/crawler/lianjia/parser"
)

var seed = engine.Request{
	Url: 		"https://www.lianjia.com/city/",
	ParserFunc: parser.ParseCityList,
}

func main() {
	engine.Run(seed)
}
