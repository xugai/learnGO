package parser

import (
	"learnGO/crawler/engine"
	"learnGO/crawler/model"
	"regexp"
)

const basicInfoRege = `<li class="fl oneline">([^<]+)</li>`

func ParseHouse(content [] byte, name string) engine.ParseResult {
	r := regexp.MustCompile(basicInfoRege)
	result := engine.ParseResult{}
	basicInfoMatches := r.FindAllSubmatch(content, -1)
	house := model.House{}
	if len(basicInfoMatches) > 0 {
		house.Name = name
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
	return result
}
