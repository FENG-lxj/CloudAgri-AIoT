package model

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"goskeleton/app/global/variable"
	"log"
	"time"
)

func CreateInfluxDBAEFactory() *AEModel {
	return &AEModel{}
}

func CreateInfluxDBAEDatasFactory() *AEModels {
	return &AEModels{}
}

func CreateInfluxDBASFactory() *ASModel {
	return &ASModel{}
}

func CreateInfluxDBASDatasFactory() *ASModels {
	return &ASModels{}
}

type AEModel struct {
	DateTime string
	AEID     int32   `json:"AEID"` //设备id
	LI       float32 `json:"LI"`   //光照强度
	PH       float32 `json:"PH"`   //土壤温度
	ST       float32 `json:"ST"`   //土壤温度
	SM       float32 `json:"SM"`   //土壤湿度
	N        float32 `json:"N"`    //土壤氮
	P        float32 `json:"P"`    //土壤磷
	K        float32 `json:"K"`    //土壤钾
	AT       float32 `json:"AT"`   //空气温度
	AH       float32 `json:"AH"`   //空气湿度
	AQ       float32 `json:"AQ"`   //空气质量
	CO2      float32 `json:"CO2"`  //二氧化碳浓度
	CH4O     float32 `json:"CH4O"` //甲醇浓度
}

type AEModels struct {
	AEDatas []AEModel `json:"AEDatas"`
}

type ASModel struct {
	DateTime string
	ASID     int32   `json:"ASID"` //设备id
	T        float32 `json:"T"`    //环境温度
	H        float32 `json:"H"`    //环境湿度
	CO2      float32 `json:"CO2"`  //二氧化碳浓度
	O2       float32 `json:"O2"`   //氧气浓度
}

type ASModels struct {
	ASDatas []ASModel `json:"ASDatas"`
}

func (ae *AEModel) GetAEdata(AEID int32) *AEModel {
	ByzyToken := variable.ConfigYml.GetString("InfluxDB.BYZY_TOKEN")
	url := variable.ConfigYml.GetString("InfluxDB.Url")
	client := influxdb2.NewClient(url, ByzyToken)

	org := variable.ConfigYml.GetString("InfluxDB.Org")

	queryAPI := client.QueryAPI(org)
	bucket := variable.ConfigYml.GetString("InfluxDB.Bucket")
	MeasureAE := variable.ConfigYml.GetString("InfluxDB.MeasureAE")
	queryFlux := fmt.Sprintf(`from(bucket: "%s")
				  |> range(start: -24h)
				  |> filter(fn: (r) => r["_measurement"] == "%s")
				  |> filter(fn: (r) => r["Device"] == "%d")
				  |> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
				  |> sort(columns: ["_time"], desc: true)
			      |> limit(n: 1)
				  |> yield(name: "AEdata")`, bucket, MeasureAE, AEID)
	results, err := queryAPI.Query(context.Background(), queryFlux)
	if err != nil {
		log.Println("GetAEdata influxdb 查询出错:", err)
		return nil
	}
	// 使用一个for循环，遍历结果集中的每个记录（Record）对象，每个记录代表一个时间序列数据的单元，包含以下四个部分：
	// 时间戳（timestamp）：表示数据的采集时间，通常是一个精确到纳秒的时间值。
	// 测量值（measurement）：表示数据的类别或者来源，通常是一个字符串，例如"cpu"，"temperature"，"event"等。
	// 标签（tag）：表示数据的元数据，通常是一组键值对，用于区分和过滤数据，例如"host"，"region"，"status"等。
	// 字段（field）：表示数据的实际值，通常是一组键值对，用于存储和计算数据，例如"value"，"count"，"speed"等。

	// 检查结果集中是否有错误，如果有，打印错误并返回
	if err := results.Err(); err != nil {
		log.Println("GetAEdata influxdb 查询结果集出错:", err)
		return nil
	}

	// 循环遍历查询结果
	for results.Next() {
		// 获取当前记录
		record := results.Record()

		ae.DateTime = record.Time().In(time.Local).Format(variable.DateFormat)
		ae.LI = float32(record.ValueByKey("LI").(float64))
		ae.PH = float32(record.ValueByKey("PH").(float64))
		ae.ST = float32(record.ValueByKey("ST").(float64))
		ae.SM = float32(record.ValueByKey("SM").(float64))
		ae.N = float32(record.ValueByKey("N").(float64))
		ae.P = float32(record.ValueByKey("P").(float64))
		ae.K = float32(record.ValueByKey("K").(float64))
		ae.AT = float32(record.ValueByKey("AT").(float64))
		ae.AH = float32(record.ValueByKey("AH").(float64))
		ae.AQ = float32(record.ValueByKey("AQ").(float64))
		ae.CO2 = float32(record.ValueByKey("CO2").(float64))
		ae.CH4O = float32(record.ValueByKey("CH4O").(float64))
		//fmt.Printf("%+v\n", ae)
	}

	return ae
}

func (aes *AEModels) GetAEdatas(AEID int32, StartTime string) *AEModels {
	ByzyToken := variable.ConfigYml.GetString("InfluxDB.BYZY_TOKEN")
	url := variable.ConfigYml.GetString("InfluxDB.Url")
	client := influxdb2.NewClient(url, ByzyToken)

	org := variable.ConfigYml.GetString("InfluxDB.Org")

	queryAPI := client.QueryAPI(org)
	bucket := variable.ConfigYml.GetString("InfluxDB.Bucket")
	MeasureAE := variable.ConfigYml.GetString("InfluxDB.MeasureAE")
	queryFlux := fmt.Sprintf(`from(bucket: "%s")
				  |> range(start: %s)
				  |> filter(fn: (r) => r["_measurement"] == "%s")
				  |> filter(fn: (r) => r["Device"] == "%d")
				  |> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
				  |> yield(name: "AEdatas")`, bucket, StartTime, MeasureAE, AEID)
	results, err := queryAPI.Query(context.Background(), queryFlux)
	if err != nil {
		log.Println("GetAEdatas influxdb 查询出错:", err)
		return nil
	}
	if err := results.Err(); err != nil {
		log.Println("GetAEdatas influxdb 查询结果集出错:", err)
		return nil
	}
	// 循环遍历查询结果
	for results.Next() {
		AEdata := AEModel{}
		// 获取当前记录
		record := results.Record()
		AEdata.DateTime = record.Time().In(time.Local).Format(variable.DateFormat)
		AEdata.AEID = AEID
		AEdata.LI = float32(record.ValueByKey("LI").(float64))
		AEdata.PH = float32(record.ValueByKey("PH").(float64))
		AEdata.ST = float32(record.ValueByKey("ST").(float64))
		AEdata.SM = float32(record.ValueByKey("SM").(float64))
		AEdata.N = float32(record.ValueByKey("N").(float64))
		AEdata.P = float32(record.ValueByKey("P").(float64))
		AEdata.K = float32(record.ValueByKey("K").(float64))
		AEdata.AT = float32(record.ValueByKey("AT").(float64))
		AEdata.AH = float32(record.ValueByKey("AH").(float64))
		AEdata.AQ = float32(record.ValueByKey("AQ").(float64))
		AEdata.CO2 = float32(record.ValueByKey("CO2").(float64))
		AEdata.CH4O = float32(record.ValueByKey("CH4O").(float64))
		aes.AEDatas = append(aes.AEDatas, AEdata)
	}
	return aes
}

func (as *ASModel) GetASdata(ASID int32) *ASModel {
	ByzyToken := variable.ConfigYml.GetString("InfluxDB.BYZY_TOKEN")
	url := variable.ConfigYml.GetString("InfluxDB.Url")
	client := influxdb2.NewClient(url, ByzyToken)

	org := variable.ConfigYml.GetString("InfluxDB.Org")

	queryAPI := client.QueryAPI(org)
	bucket := variable.ConfigYml.GetString("InfluxDB.Bucket")
	MeasureAS := variable.ConfigYml.GetString("InfluxDB.MeasureAS")
	queryFlux := fmt.Sprintf(`from(bucket: "%s")
				  |> range(start: -24h)
				  |> filter(fn: (r) => r["_measurement"] == "%s")
				  |> filter(fn: (r) => r["Device"] == "%d")
				  |> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
				  |> sort(columns: ["_time"], desc: true)
			      |> limit(n: 1)
				  |> yield(name: "ASdata")`, bucket, MeasureAS, ASID)
	results, err := queryAPI.Query(context.Background(), queryFlux)
	if err != nil {
		log.Println("GetASdata influxdb 查询出错:", err)
		return nil
	}
	// 检查结果集中是否有错误，如果有，打印错误并返回
	if err := results.Err(); err != nil {
		log.Println("GetASdata influxdb 查询结果集出错:", err)
		return nil
	}

	// 循环遍历查询结果
	for results.Next() {
		// 获取当前记录
		record := results.Record()

		as.DateTime = record.Time().In(time.Local).Format(variable.DateFormat)
		as.T = float32(record.ValueByKey("T").(float64))
		as.H = float32(record.ValueByKey("H").(float64))
		as.CO2 = float32(record.ValueByKey("CO2").(float64))
		as.O2 = float32(record.ValueByKey("O2").(float64))
		//fmt.Printf("%+v\n", as)
	}

	return as
}

func (ass *ASModels) GetASdatas(ASID int32, StartTime string) *ASModels {
	ByzyToken := variable.ConfigYml.GetString("InfluxDB.BYZY_TOKEN")
	url := variable.ConfigYml.GetString("InfluxDB.Url")
	client := influxdb2.NewClient(url, ByzyToken)

	org := variable.ConfigYml.GetString("InfluxDB.Org")

	queryAPI := client.QueryAPI(org)
	bucket := variable.ConfigYml.GetString("InfluxDB.Bucket")
	MeasureAS := variable.ConfigYml.GetString("InfluxDB.MeasureAS")
	queryFlux := fmt.Sprintf(`from(bucket: "%s")
				  |> range(start: %s)
				  |> filter(fn: (r) => r["_measurement"] == "%s")
				  |> filter(fn: (r) => r["Device"] == "%d")
				  |> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
				  |> yield(name: "ASdatas")`, bucket, StartTime, MeasureAS, ASID)
	results, err := queryAPI.Query(context.Background(), queryFlux)
	if err != nil {
		log.Println("GetASdatas influxdb 查询出错:", err)
		return nil
	}
	if err := results.Err(); err != nil {
		log.Println("GetASdatas influxdb 查询结果集出错:", err)
		return nil
	}
	// 循环遍历查询结果
	for results.Next() {
		ASdata := ASModel{}
		// 获取当前记录
		record := results.Record()
		ASdata.DateTime = record.Time().In(time.Local).Format(variable.DateFormat)
		ASdata.ASID = ASID
		ASdata.T = float32(record.ValueByKey("T").(float64))
		ASdata.H = float32(record.ValueByKey("H").(float64))
		ASdata.CO2 = float32(record.ValueByKey("CO2").(float64))
		ASdata.O2 = float32(record.ValueByKey("O2").(float64))
		ass.ASDatas = append(ass.ASDatas, ASdata)
	}
	return ass
}
