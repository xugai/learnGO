package controller

import (
	"context"
	elastic "github.com/olivere/elastic/v7"
	"learnGO/crawler/engine"
	"learnGO/crawler/frontend/model"
	"learnGO/crawler/frontend/view"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultViewHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

func InitialSearchResultViewHandler(fileName string) SearchResultViewHandler {

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultViewHandler{
		view: view.GetSearchResultView(fileName),
		client: client,
	}
}

// localhost:8888/search?q=朝南 集中供暖&from=10
func (s SearchResultViewHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	q := strings.TrimSpace(request.FormValue("q"))
	from, err := strconv.Atoi(request.FormValue("from"))
	if err != nil {
		from = 0
	}
	//fmt.Fprintf(writer, "q = %s, from = %v", q, from)
	var page model.SearchResult
	page, err = s.getSearchResult(q, from)

	s.view.Render(writer, page)
}

func (s SearchResultViewHandler) getSearchResult(q string, from int) (model.SearchResult, error) {

	var result model.SearchResult
	resp, err := s.client.Search("dating_house").
					Query(elastic.NewQueryStringQuery(s.rewriteQueryString(q))).From(from).Do(context.Background())
	if err != nil {
		return model.SearchResult{}, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Query = q
	result.PrevFrom = from - 10
	result.NextFrom = from + 10
	for _, v := range resp.Each(reflect.TypeOf(engine.Item{})) {
		result.Items = append(result.Items, v.(engine.Item))
	}
	return result, nil
}

func (s SearchResultViewHandler) rewriteQueryString(q string) string {
	rege := regexp.MustCompile(`([A-Za-z]+):`)
	return rege.ReplaceAllString(q, "Payload.$1:")
}

