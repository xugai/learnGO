package parser

import (
	"learnGO/crawler/engine"
	"learnGO/crawler/model"
	"regexp"
)

const basicInfoRege = `<li class="fl oneline">[^：]+：([^<]+)</li>`
var basicInfoR = regexp.MustCompile(basicInfoRege)

const recommendRege = `<div class="recommend-item">
                <a href="([^"]+)" data-event_id="[0-9]+" data-event_action="rank_id=[0-9]{1}">
                <img src="[^"]+" alt="">
                <p class="title">([^<]+)</p>
                                <p class="desc">[^<]+</p>
                                                <p class="price">(-?[0-9]+)元/月</p>
                                  <p class="tips">[^<]+</p>
                                </a>
              </div>`
var recommendR = regexp.MustCompile(recommendRege)
const recommendRege2 = `<div class="recommend-item">
                <a href="([^"]+)" data-event_id="[0-9]+" data-event_action="rank_id=[0-9]{1}">
                <img src="[^"]+" alt="">
                <p class="title">([^<]+)</p>
                                <p class="desc">[^<]+</p>
                                                <p class="desc">[^<]+</p>
                                <p class="price">(-?[0-9]+)元/月</p>
                                  <p class="tips">[^<]+</p>
                                </a>
              </div>`
var recommendR2 = regexp.MustCompile(recommendRege2)
func ParseHouse(content [] byte, name string, rent string, url string, id string) engine.ParseResult {

	result := engine.ParseResult{}
	basicInfoMatches := basicInfoR.FindAllSubmatch(content, -1)
	recommendMatches := recommendR.FindAllSubmatch(content, -1)
	recommendMatches2 := recommendR2.FindAllSubmatch(content, -1)
	prefixUrlMatch := prefixUrlR.FindSubmatch(content)
	house := model.House{}
	if len(basicInfoMatches) > 0 {
		house.Name = name
		house.Rent = rent
		house.Area = string(basicInfoMatches[0][1])
		house.Towards = string(basicInfoMatches[1][1])
		house.Maintain = string(basicInfoMatches[2][1])
		house.CheckIn = string(basicInfoMatches[3][1])
		house.Floor = string(basicInfoMatches[4][1])
		house.Elevator = string(basicInfoMatches[5][1])
		house.ParkingSpace = string(basicInfoMatches[6][1])
		house.WaterUsed = string(basicInfoMatches[7][1])
		house.ElectricityUsed = string(basicInfoMatches[8][1])
		house.GasUsed = string(basicInfoMatches[9][1])
		house.Heating = string(basicInfoMatches[10][1])
		house.LeaseTerm = string(basicInfoMatches[11][1])
		house.HouseVisit = string(basicInfoMatches[12][1])

		result.Items = [] engine.Item{
			{
				Url: url,
				Id: id,
				Payload: house,
			},
		}
	}
	if len(recommendMatches) > 0 {
		for _, recommendMatch := range recommendMatches {
			houseName := string(recommendMatch[2])
			rent := string(recommendMatch[3]) + "元/月"
			url := "https://" + string(prefixUrlMatch[1]) + string(recommendMatch[1])
			id := StringExtractor(string(recommendMatch[1]), `/zufang/([^\.]+).html`)
			result.Requests = append(result.Requests, engine.Request{
				Url: url,
				ParserFunc: getParseHouseFunc(houseName, rent, url, id),
			})
		}
	}
	if len(recommendMatches2) > 0 {
		for _, recommendMatch2 := range recommendMatches2 {
			houseName := string(recommendMatch2[2])
			rent := string(recommendMatch2[3]) + "元/月"
			url := "https://" + string(prefixUrlMatch[1]) + string(recommendMatch2[1])
			id := StringExtractor(string(recommendMatch2[1]), `/zufang/([^\.]+).html`)
			result.Requests = append(result.Requests, engine.Request{
				Url: url,
				ParserFunc: getParseHouseFunc(houseName, rent, url, id),
			})
		}
	}
	return result
}

func StringExtractor(target string, rule string) string {
	r := regexp.MustCompile(rule)
	matches := r.FindSubmatch([] byte(target))
	return string(matches[1])
}
