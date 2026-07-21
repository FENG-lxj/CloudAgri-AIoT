package service

import (
	"IOTA_ZZ_server/mapper"
	"IOTA_ZZ_server/model"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func ICLogToDB() mqtt.MessageHandler {
	return func(mqttClient mqtt.Client, msg mqtt.Message) {
		datas := model.ICLog{}
		err := json.Unmarshal(msg.Payload(), &datas)
		if err != nil {
			log.Println("ICLogToDB 解组消息时出错:", err)
			return
		}
		log.Println("ICLogToDB 收到数据: \n", datas)
		mapper.SaveICLog(datas)
	}
}
