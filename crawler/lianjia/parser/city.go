package parser

import (
	"fmt"
	"learnGO/crawler/engine"
	"regexp"
	"strconv"
)

const prefixUrlRege = `<a href="//([^"]+)" target="_blank">首页</a>`

const mainHouseRege = `<a
      class="content__list--item--aside" target="_blank"      href="(/zufang/[0-9a-zA-Z]+\.html)"
      title="([^>]+)">`
const streetRege = `<p class="content__list--item--des">
        <a target="_blank" href="/zufang/[0-9a-zA-Z]+/">([^<]+)</a>-<a href="/zufang/[0-9a-zA-Z]+/" target="_blank">([^<]+)</a>-<a title="[^"]+" href="/zufang/[0-9a-zA-Z]+/" target="_blank">([^<]+)</a>
        <i>/</i>
        ([^<]+)
        <i>/</i>([^<]+)<i>/</i>
          ([^<]+)<span class="hide">
          <i>/</i>
          ([^<]+)
                  </span>
      </p>`
const priceRege = `<span class="content__list--item-price"><em>(-?[0-9]+)</em> 元/月</span>`
const pageRege = `<div class="content__pg" data-el="page_navigation" data-url="/zufang/pg{page}/" data-totalPage=-?[0-9]+ data-curPage=([0-9]+)>`

var prefixUrlR = regexp.MustCompile(prefixUrlRege)
var mainHouseR = regexp.MustCompile(mainHouseRege)
var streetR = regexp.MustCompile(streetRege)
var priceR = regexp.MustCompile(priceRege)
var pageR = regexp.MustCompile(pageRege)

func ParseCity(content [] byte) engine.ParseResult {

	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Abnormal Error occured: ", err)
		}
	}()

	prefixUrlMatches := prefixUrlR.FindAllSubmatch(content, -1)
	mainHouseMatches := mainHouseR.FindAllSubmatch(content, -1)
	priceMatches := priceR.FindAllSubmatch(content, -1)
	pageMatch := pageR.FindSubmatch(content)
	result := engine.ParseResult{}


	for i, m := range mainHouseMatches {
		houseName := string(m[2])
		rent := string(priceMatches[i][1]) + "元/月"
		url := "https://" + string(prefixUrlMatches[0][1]) + string(m[1])
		id := StringExtractor(string(m[1]), `/zufang/([^\.]+).html`)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: getParseHouseFunc(houseName, rent, url, id),
		})

		//result.Items = append(result.Items,
		//	string(m[2]) + " " + string(streetmatches[i][1]) + string(streetmatches[i][2]) + string(streetmatches[i][3]) + " " + string(streetmatches[i][4]) + " " +
		//	strings.Trim(string(streetmatches[i][5]), " ") + strings.Trim(string(streetmatches[i][6]), " ") + " " + strings.Join(strings.Fields(string(streetmatches[i][7])), "") +
		//	" " + string(priceMatches[i][1]) + "元/月")
	}

	if len(pageMatch) > 0 {
		result.Requests = append(result.Requests, engine.Request{
			Url: "https://" + string(prefixUrlMatches[0][1]) + "/zufang/" + "pg" + GetNextPage(string(pageMatch[1])),
			ParserFunc: ParseCity,
		})
	}
	return result
}

func GetNextPage(currPage string) string {
	i, err := strconv.Atoi(currPage)
	if err != nil {
		panic(err)
	}
	a := strconv.Itoa(i + 1)
	return a
}

func getParseHouseFunc(houseName string, rent string, url string, id string) func([] byte) engine.ParseResult {
	return func(bytes []byte) engine.ParseResult {
		return ParseHouse(bytes, houseName, rent, url, id)
	}
}