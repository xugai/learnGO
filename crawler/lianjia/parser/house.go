package parser

import (
	"learnGO/crawler/engine"
	"learnGO/crawler/model"
	"regexp"
)

const basicInfoRege = `<li class="fl oneline">([^<]+)</li>`
var basicInfoR = regexp.MustCompile(basicInfoRege)

const recommendRege = `<div class="recommend-item">
                <a href="([^"]+)" data-event_id="[0-9]+" data-event_action="rank_id=0">
                <img src="[^"]+" alt="">
                <p class="title">([^<]+)</p>
                                <p class="desc">[^<]+</p>
                                                <p class="price">(-?[0-9]+)元/月</p>
                                  <p class="tips">[^<]+</p>
                                </a>
              </div>`
var recommendR = regexp.MustCompile(recommendRege)

func ParseHouse(content [] byte, name string, rent string) engine.ParseResult {

	result := engine.ParseResult{}
	basicInfoMatches := basicInfoR.FindAllSubmatch(content, -1)
	recommendMatches := recommendR.FindAllSubmatch(content, -1)
	prefixUrlMatch := prefixUrlR.FindSubmatch(content)
	house := model.House{}
	if len(basicInfoMatches) > 0 {
		house.Name = name
		house.Rent = rent
		house.Area = string(basicInfoMatches[1][1])
		house.Towards = string(basicInfoMatches[2][1])
		house.Maintain = string(basicInfoMatches[4][1])
		house.CheckIn = string(basicInfoMatches[5][1])
		house.Floor = string(basicInfoMatches[7][1])
		house.Elevator = string(basicInfoMatches[8][1])
		house.ParkingSpace = string(basicInfoMatches[10][1])
		house.WaterUsed = string(basicInfoMatches[11][1])
		house.ElectricityUsed = string(basicInfoMatches[13][1])
		house.GasUsed = string(basicInfoMatches[14][1])
		house.Heating = string(basicInfoMatches[16][1])
		house.LeaseTerm = string(basicInfoMatches[17][1])
		house.HouseVisit = string(basicInfoMatches[18][1])

		result.Items = [] interface{}{house}
	}
	if len(recommendMatches) > 0 {
		for _, recommendMatch := range recommendMatches {
			houseName := string(recommendMatch[2])
			rent := string(recommendMatch[3]) + "元/月"
			result.Requests = append(result.Requests, engine.Request{
				Url: "https://" + string(prefixUrlMatch[1]) + string(recommendMatch[1]),
				ParserFunc: func(bytes []byte) engine.ParseResult {
					return ParseHouse(bytes, houseName, rent)
				},
			})
		}
	}
	return result
}
