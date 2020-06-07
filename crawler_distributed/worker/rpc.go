package worker

import (
	"learnGO/crawler/engine"
	"log"
)

type CrawlService struct {

}

func (CrawlService) Process(request Request, result *ParseResult) error {
	req, err := DeserializeRequest(request)
	if err != nil {
		return err
	}
	log.Printf("Get request from client: %s\n", request.Url)
	parseResult, err := engine.Worker(req)
	if err != nil {
		return err
	}
	*result = SerializeParseResult(parseResult)
	return nil
}

