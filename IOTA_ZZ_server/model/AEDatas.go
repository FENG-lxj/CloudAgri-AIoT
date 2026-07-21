package model

type AEDatas struct {
	AEID      int32   `json:"AEID"`      //设备id
	LI        float32 `json:"LI"`        //光照强度
	PH        float32 `json:"PH"`        //土壤温度
	ST        float32 `json:"ST"`        //土壤温度
	SM        float32 `json:"SM"`        //土壤湿度
	N         float32 `json:"N"`         //土壤氮
	P         float32 `json:"P"`         //土壤磷
	K         float32 `json:"K"`         //土壤钾
	AT        float32 `json:"AT"`        //空气温度
	AH        float32 `json:"AH"`        //空气湿度
	AQ        float32 `json:"AQ"`        //空气质量
	CO2       float32 `json:"CO2"`       //二氧化碳浓度
	CH4O      float32 `json:"CH4O"`      //甲醇浓度
	Longitude float64 `json:"Longitude"` //经度
	Dimension float64 `json:"Dimension"` //纬度
}

//type AEDDateLD struct {
//	AEID      int32   `json:"AEID"`      //设备id
//	Longitude float64 `json:"Longitude"` //经度
//	Dimension float64 `json:"Dimension"` //纬度
//}
