package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"learnGO/crawler/engine"
	"log"
)

/**
	docker run elastic command: docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.7.0
 */

func ItemSaver(client *elastic.Client) chan engine.Item {
	itemChannel := make(chan engine.Item)
	itemCount := 1
	go func() {
		for {
			item := <- itemChannel
			log.Printf("ItemSaver save #%d item %v\n", itemCount, item)
			itemCount++

			// save item into elastic search engine
			_, err := Save(client, item)
			if err != nil {
				log.Printf("Save item %v failed. error: %v\n", item, err)
			}
		}
	}()
	return itemChannel
}

func Save(client *elastic.Client, item engine.Item) (id string, err error){
	// set sniff to false, because we run our elastic search single instance in docker.
	indexService := client.Index().
		Index("dating_house").
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	resp, err := indexService.Do(context.Background())
	if err != nil {
		return "", err
	}

	return resp.Id, nil
}

