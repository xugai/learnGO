package parser

import (
	"io/ioutil"
	"learnGO/crawler/model"
	"testing"
)

func TestParseHouse(t *testing.T) {
	content, err := ioutil.ReadFile("house_test_data.html")
	if err != nil {
		panic(err)
	}

	expectedResult := model.House{
		Name: "整租·北正街 1室1厅 南/北",
		Area: "面积：40㎡",
		Towards: "朝向：南 北",
		Maintain: "维护：5天前",
		CheckIn: "入住：随时入住",
		Floor: "楼层：高楼层/6层",
		Elevator: "电梯：无",
		ParkingSpace: "车位：暂无数据",
		WaterUsed: "用水：暂无数据",
		ElectricityUsed: "用电：暂无数据",
		GasUsed: "燃气：暂无数据",
		Heating: "采暖：暂无数据",
		LeaseTerm: "租期：暂无数据",
		HouseVisit: "看房：需提前预约",
	}

	result := ParseHouse(content, "整租·北正街 1室1厅 南/北")
	if len(result.Items) == 0 {
		t.Errorf("Expected result.Items' length is 1, but get 0")
	}
	if result.Items[0].(model.House) != expectedResult {
		t.Errorf("Expected result is %v, but get %v", expectedResult, result.Items[0])
	}
}
