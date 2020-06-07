package main

import (
	"flag"
	"learnGO/crawler/engine"
	"learnGO/crawler/lianjia/parser"
	"learnGO/crawler/scheduler"
	"learnGO/crawler_distributed/config"
	"learnGO/crawler_distributed/persist/client"
	"learnGO/crawler_distributed/rpcsupport"
	client2 "learnGO/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

// 183.48.244.8

var seed = engine.Request{
	Url: 		"https://cd.lianjia.com/zufang/",
	Parser: 	engine.NewFuncParser(parser.ParseCity, "ParseCity"),
}

var (
	itemSaverHost = flag.String("itemsaver_host", "", "this port use to connect to elasticsearch.")
	workerHosts = flag.String("worker_host", "", "this port use to connect crawler server, comma to seprate.")
	)

func main() {
	flag.Parse()
	// 实现多节点分布式抓取信息
	clientChannel, _ := createClientPool(strings.Split(*workerHosts, ","))
	processor, err := client2.CreateProcessor(clientChannel)
	if err != nil {
		panic(err)
	}
	engine := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: config.WORKER_COUNT,
		ItemChannel: client.ItemSaver(*itemSaverHost),
		RequestProcessor: processor,
	}
	engine.Run(seed)
}

func createClientPool(hosts [] string) (chan *rpc.Client, error) {
	var clients [] *rpc.Client
	clientChannel := make(chan *rpc.Client)
	for _, host := range hosts {
		newClient, err := rpcsupport.NewClient(host)
		if err != nil {
			log.Printf("Create new client error: %v\n", err)
			return nil, nil
		}
		clients = append(clients, newClient)
	}
	go func() {
		for {
			for _, c := range clients {
				clientChannel <- c
			}
		}
	}()
	return clientChannel, nil
}
