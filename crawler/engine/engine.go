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
		result, err := worker(request)
		if err != nil {
			panic(err)
		}
		requests = append(requests, result.Requests ...)
		for _, item := range result.Items {
			fmt.Printf("Get item: %v\n", item)
		}
		fmt.Printf( "Fetch total %v items\n", len(result.Items))
	}
}

func worker(request Request) (ParseResult, error) {
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetcher error: fetching Url %s:  %v", request.Url, err)
		return ParseResult{}, err
	}
	result := request.ParserFunc(content)
	return result, nil
}
