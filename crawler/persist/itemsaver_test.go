package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"learnGO/crawler/engine"
	"learnGO/crawler/model"
	"testing"
)

func TestSaver(t *testing.T) {
	// define test data
	expectedPayload := model.House{
		Name: "整租·北正街 1室1厅 南/北",
		Rent: "700元/月",
		Area: "40㎡",
		Towards: "南 北",
		Maintain: "5天前",
		CheckIn: "2020-11-11",
		Floor: "高楼层/6层",
		Elevator: "无",
		ParkingSpace: "暂无数据",
		WaterUsed: "暂无数据",
		ElectricityUsed: "暂无数据",
		GasUsed: "暂无数据",
		Heating: "暂无数据",
		LeaseTerm: "暂无数据",
		HouseVisit: "需提前预约",
	}
	expectedItem := engine.Item{
		Url: "https://aq.lianjia.com/zufang/AQ2464281991093428226.html",
		Id: "AQ2464281991093428226",
		Payload: expectedPayload,
	}

	// save test data to elastic search engine
	client, _ := elastic.NewClient(elastic.SetSniff(false))
	_, err := Save(client, expectedItem)
	if err != nil {
		panic(err)
	}

	// get saved data
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().
		Index("dating_house").
		Id(expectedItem.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s\n", resp.Source)

	// verify actualItem
	var actualResult engine.Item
	json.Unmarshal(resp.Source, &actualResult)
	house, err := model.ConvertToHouseFromjsonObj(actualResult.Payload)
	if err != nil {
		panic(err)
	}
	actualResult.Payload = house

	if actualResult != expectedItem {
		t.Errorf("expected %v, but get %v\n", expectedItem, actualResult)
	}
}
