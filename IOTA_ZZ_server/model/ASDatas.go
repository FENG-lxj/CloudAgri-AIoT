package model

type ASDatas struct {
	ASID int32   `json:"ASID"` //设备id
	T    float32 `json:"T"`    //温度
	H    float32 `json:"H"`    //湿度
	CO2  float32 `json:"CO2"`  //二氧化碳浓度
	O2   float32 `json:"O2"`   //氧气浓度
}
