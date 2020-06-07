package main

import (
	"flag"
	"fmt"
	"learnGO/crawler_distributed/rpcsupport"
	"learnGO/crawler_distributed/worker"
	"log"
)

var crawlerHost = flag.Int("crawler_host", 0, "this is use to connect crawler service.")

func main() {
	flag.Parse()
	log.Fatal(rpcsupport.ServeRPC(fmt.Sprintf(":%d", *crawlerHost), worker.CrawlService{}))

}
