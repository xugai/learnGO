package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"learnGO/crawler_distributed/persist"
	"learnGO/crawler_distributed/rpcsupport"
	"log"
)

var itemSaverHost = flag.Int("itemsaver_host", 0, "this is use to connect elasticsearch.")

func main() {
	flag.Parse()
	log.Fatal(serveRPC(fmt.Sprintf(":%d", *itemSaverHost)))
}

func serveRPC(host string) error {
	c, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServeRPC(host, &persist.ItemSaverService{
		Client: c,
	})

}
