package main

import (
	"fmt"
	"learnGO/crawler/engine"
	"learnGO/crawler/lianjia/parser"
	"learnGO/crawler_distributed/config"
	"learnGO/crawler_distributed/rpcsupport"
	"learnGO/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawleService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRPC(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		t.Error(err)
	}

	req := engine.Request{
		Url: "https://bj.lianjia.com/zufang/BJ2510698080752369664.html",
		Parser: parser.NewHouseParser("整租·时代国际嘉园 1室1厅 南",
										"9000",
											"https://bj.lianjia.com/zufang/BJ2510698080752369664.html",
											"BJ2510698080752369664"),
	}

	var result worker.ParseResult
	err = client.Call(config.CRAWLER_SERVICE, worker.SerializeRequest(req), &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%v\n", result)
	}
}
