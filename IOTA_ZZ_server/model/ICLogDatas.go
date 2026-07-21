package model

type ICLog struct {
	ICID        int32  `json:"ICID"`        //摄像头设备id
	Datetime    string `json:"Datetime"`    //日期时间
	Level1Class int    `json:"Level1Class"` //一级分类
	Log         string `json:"Log"`         //日志内容
}

type ICHB struct {
	ICID int32 `json:"ICID"` //摄像头设备id
	HB   int   `json:"HB"`   //心跳状态
}
