package main

import (
	"learnGO/crawler/engine"
	"learnGO/crawler/model"
	"learnGO/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaverService(t *testing.T) {

	// start itemsaverservice server end
	go serveRPC()
	// start client end
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(":1234")
	if err != nil {
		panic(err)
	}
	// client end call itemsaverservice'service with jsonrpc
	house := model.House{
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
	}
	item := engine.Item{
		Url: "https://aq.lianjia.com/zufang/AQ2464281991093428226.html",
		Id: "AQ2464281991093428226",
		Payload: house,
	}
	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)

	// verify
	if err != nil || result != "ok" {
		t.Errorf("Get err: %s, result is: %s\n", err, result)
	}
}
