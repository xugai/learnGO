package persist

import (
	"github.com/olivere/elastic/v7"
	"learnGO/crawler/engine"
	"learnGO/crawler/persist"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	_, err := persist.Save(s.Client, item)
	log.Printf("Save item %v\n", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Fatalf("Save item %v error: %v\n", item, err)
	}
	return nil
}
