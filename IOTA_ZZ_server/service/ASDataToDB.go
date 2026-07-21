package service

import (
	"IOTA_ZZ_server/mapper"
	"IOTA_ZZ_server/model"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func ASDataToDB() mqtt.MessageHandler {
	return func(mqttClient mqtt.Client, msg mqtt.Message) {
		datas := model.ASDatas{}

		err := json.Unmarshal(msg.Payload(), &datas)

		if err != nil {
			log.Println("ASDataToDB 解组消息时出错:", err)
			return
		}
		log.Println("ASDataToDB 收到数据: \n", datas)
		if !mapper.SelectDevExist("AS", datas.ASID) {
			mapper.InsertDevInfo("AS", datas.ASID)
		}
		mapper.SaveASDatas(datas)
		mapper.UpASDevLatestNewsTime(datas.ASID)
	}
}
