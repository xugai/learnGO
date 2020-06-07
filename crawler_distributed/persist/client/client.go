package client

import (
	"learnGO/crawler/engine"
	"learnGO/crawler_distributed/config"
	"learnGO/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) chan engine.Item {
	itemChannel := make(chan engine.Item)
	c, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	itemCount := 1
	go func() {
		for {
			item := <- itemChannel
			log.Printf("ItemSaver save #%d item %v\n", itemCount, item)
			itemCount++

			// save item into elastic search engine
			result := ""
			err = c.Call(config.ITEMSAVER_SERVICE, item, &result)
			if err != nil {
				log.Printf("Save item %v failed. error: %v\n", item, err)
			}
		}
	}()
	return itemChannel
}
