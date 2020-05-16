package persist

import "log"

func ItemSaver() chan interface{} {
	itemChannel := make(chan interface{})
	itemCount := 1
	go func() {
		for {
			item := <- itemChannel
			log.Printf("ItemSaver Get #%d item %v\n", itemCount, item)
			itemCount++
		}
	}()
	return itemChannel
}
