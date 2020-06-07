package engine

import (
	"fmt"
	"learnGO/crawler/fetcher"
	"log"
)

type SingleEngine struct {

}

func (s SingleEngine) Run(seeds ... Request) {
	var requests [] Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		result, err := Worker(request)
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

func Worker(request Request) (ParseResult, error) {
	//log.Printf("Fetching url: %s\n", request.Url)
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetcher error: fetching Url %s:  %v", request.Url, err)
		return ParseResult{}, err
	}
	result := request.Parser.Parse(content)
	return result, nil
}
