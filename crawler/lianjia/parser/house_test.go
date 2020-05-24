package parser

import (
	"io/ioutil"
	"learnGO/crawler/model"
	"testing"
)

func TestParseHouse(t *testing.T) {
	content, err := ioutil.ReadFile("house_test_data.txt")
	if err != nil {
		panic(err)
	}

	expectedResult := model.House{
		Name: "整租·欣葆家园一区 1室1厅 南",
		Rent: "4000元/月",
		Area: "50㎡",
		Towards: "南",
		Maintain: "1天前",
		CheckIn: "随时入住",
		Floor: "高楼层/15层",
		Elevator: "有",
		ParkingSpace: "暂无数据",
		WaterUsed: "暂无数据",
		ElectricityUsed: "暂无数据",
		GasUsed: "暂无数据",
		Heating: "暂无数据",
		LeaseTerm: "暂无数据",
		HouseVisit: "需提前预约",
	}

	result := ParseHouse(content,
					"整租·欣葆家园一区 1室1厅 南",
					"4000元/月",
					"https://bj.lianjia.com//zufang/BJ2518580725339930624.html",
					"BJ2518580725339930624")

	if len(result.Items) != 1 {
		t.Errorf("Expected result.Items' length is 1, but get %v\n", len(result.Items))
	}
	if len(result.Requests) != 5 {
		t.Errorf("Expected result.Requests' length is 5, but get %v\n", len(result.Requests))
	}
	if result.Items[0].Payload.(model.House) != expectedResult {
		t.Errorf("Expected result is %v, but get %v\n", expectedResult, result.Items[0])
	}


}
