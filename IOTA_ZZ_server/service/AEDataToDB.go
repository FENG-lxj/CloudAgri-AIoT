package service

import (
	"IOTA_ZZ_server/mapper"       // 数据库操作层（增删改查SQL）
	"IOTA_ZZ_server/model"        // 数据模型，AE设备结构体在这里定义
	"encoding/json"               // JSON序列化/反序列化
	mqtt "github.com/eclipse/paho.mqtt.golang" // MQTT客户端依赖
	"log"                         // 日志打印
)

// AEDataToDB 返回 mqtt.MessageHandler 类型（MQTT标准消息回调函数）
func AEDataToDB() mqtt.MessageHandler {
	// 返回匿名回调函数，MQTT收到对应主题消息时自动执行这个函数
	return func(mqttClient mqtt.Client, msg mqtt.Message) {
		// 1. 定义空结构体，用来存放解析后的AE设备数据
		datas := model.AEDatas{}

		// 2. msg.Payload() 获取MQTT消息二进制字节流，json.Unmarshal JSON转结构体
		err := json.Unmarshal(msg.Payload(), &datas)
		if err != nil {
			// JSON格式错误、字段不匹配时打印日志，直接return终止本次处理
			log.Println("AEDataToDB 解组消息时出错:", err)
			return
		}

		// 打印解析成功后的设备数据，方便调试查看
		log.Println("AEDataToDB 收到数据: \n", datas)

		// 3. 查询数据库：该AE设备是否已入库
		// 参数1：设备类型AE；参数2：设备唯一ID
		if !mapper.SelectDevExist("AE", datas.AEID) {
			// 数据库查不到该设备 → 新增设备基础信息
			mapper.InsertDevInfo("AE", datas.AEID)
		}

		// 4. 把本次上报的传感器时序数据存入数据库（MySQL/InfluxDB由mapper实现）
		mapper.SaveAEDatas(datas)
	}
}
//// 位置
//func AEDLDataToDB() mqtt.MessageHandler {
//	return func(mqttClient mqtt.Client, msg mqtt.Message) {
//
//		data := model.AEDDateLD{}
//		err := json.Unmarshal(msg.Payload(), &data)
//		if err != nil {
//			log.Println("AEDLDataToDB 解组消息时出错:", err)
//			return
//		}
//		log.Println("AEDLDataToDB 收到数据: \n", data)
//		if !mapper.SelectDevExist("AE", data.AEID) {
//			mapper.InsertDevInfo("AE", data.AEID)
//		}
//		mapper.SaveAEDevGPS(data.AEID, data.Longitude, data.Dimension)
//	}
//}
