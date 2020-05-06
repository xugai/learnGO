package engine

import (
	"fmt"
	"learnGO/crawler/fetcher"
	"log"
)

func Run(seeds ... Request) {
	var requests [] Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		fmt.Printf("Fetching url: %s\n", request.Url)
		content, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("Fetcher error: fetching Url %s:  %v", request.Url, err)
			continue
		}
		result := request.ParserFunc(content)
		requests = append(requests, result.Requests ...)
		for _, item := range result.Items {
			fmt.Printf("Get item: %v\n", item)
		}
		fmt.Printf("Fetch %v items\n", len(result.Items))
	}
}
