package parser

import (
	"learnGO/crawler/engine"
	"regexp"
)

const citylistRege = `<a href="(https://[a-z]+\.lianjia\.com/)">([^<]+)</a>`

func ParseCityList(content [] byte) engine.ParseResult {
	r := regexp.MustCompile(citylistRege)
	matches := r.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: engine.NilParserFunc,
		})
		result.Items = append(result.Items, string(m[2]))
	}
	return result
}
