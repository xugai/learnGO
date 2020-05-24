package model

import "encoding/json"

type House struct {
	Name string
	Rent string
	Area string
	Towards string
	Maintain string
	CheckIn string
	Floor string
	Elevator string
	ParkingSpace string
	WaterUsed string
	ElectricityUsed string
	GasUsed string
	Heating string
	LeaseTerm string
	HouseVisit string
}

func ConvertToHouseFromjsonObj(obj interface{}) (House, error) {
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	var house House
	err = json.Unmarshal(jsonStr, &house)
	return house, err
}
