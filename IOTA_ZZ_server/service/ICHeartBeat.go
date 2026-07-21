package service

import (
	"IOTA_ZZ_server/mapper"
	"IOTA_ZZ_server/model"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func ICHeartBeat() mqtt.MessageHandler {
	return func(mqttClient mqtt.Client, msg mqtt.Message) {
		datas := model.ICHB{}
		err := json.Unmarshal(msg.Payload(), &datas)
		if err != nil {
			log.Println("ICHeartBeat 解组消息时出错:", err)
			return
		}
		log.Println("ICHeartBeat 收到数据: \n", datas)
		if !mapper.SelectDevExist("IC", datas.ICID) {
			mapper.InsertDevInfo("IC", datas.ICID)
		}
		mapper.UPICHeartBeat(datas.ICID)
	}
}
