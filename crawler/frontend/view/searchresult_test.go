package view

import (
	"learnGO/crawler/engine"
	"learnGO/crawler/frontend/model"
	common "learnGO/crawler/model"
	"os"
	"testing"
)

func TestGetSearchResultView(t *testing.T) {
	temp := GetSearchResultView("template.html")
	page := model.SearchResult{}
	item := engine.Item{
		Url: "https://aq.lianjia.com/zufang/AQ2464281991093428224.html",
		Id: "AQ2464281991093428224",
		Payload: common.House{
			Name: "整租·北正街 1室1厅 南/北",
			Rent: "700元/月",
			Area: "40㎡",
			Towards: "南 北",
			Maintain: "5天前",
			CheckIn: "随时入住",
			Floor: "高楼层/6层",
			Elevator: "无",
			ParkingSpace: "暂无数据",
			WaterUsed: "暂无数据",
			ElectricityUsed: "暂无数据",
			GasUsed: "暂无数据",
			Heating: "暂无数据",
			LeaseTerm: "暂无数据",
			HouseVisit: "需提前预约",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	page.Hits = 10
	page.Start = 1
	page.PrevFrom = 1
	page.NextFrom = 2
	out, err := os.Create("template.test.html")

	err = temp.Render(out, page)
	if err != nil {
		panic(err)
	}
}
