package parser

import (
	"learnGO/crawler/engine"
	"regexp"
)

const citylistRege = `<a href="(https://[a-z]+\.lianjia\.com/)">([^<]+)</a>`
var r = regexp.MustCompile(citylistRege)

func ParseCityList(content [] byte) engine.ParseResult {

	matches := r.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]) + "zufang/",
			ParserFunc: ParseCity,
		})
		//result.Items = append(result.Items, string(m[2]))
	}
	return result
}
