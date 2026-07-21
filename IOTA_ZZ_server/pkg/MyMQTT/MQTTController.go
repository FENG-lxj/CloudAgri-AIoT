package MyMQTT

import (
	"IOTA_ZZ_server/service"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"log"
)

func Controller() {
	opts := Init()                     // 初始化MQTT配置
	mqttClient := mqtt.NewClient(opts) // 创建MQTT客户端
	// 连接MQTT服务器
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("连接MQTT服务器失败: ", token.Error())
	}
	// 订阅农业生态环境监测数据
	if token := mqttClient.Subscribe(viper.GetString("MQTT.AETopic"), 0, service.AEDataToDB()); token.Wait() && token.Error() != nil {
		log.Fatal(viper.GetString("MQTT.AETopic"), " 主题订阅失败: ", token.Error())
	}
	//if token := mqttClient.Subscribe(viper.GetString("MQTT.AETopic"), 0, service.AEDLDataToDB()); token.Wait() && token.Error() != nil {
	//	log.Fatal(viper.GetString("MQTT.AETopic"), " 主题订阅失败: ", token.Error())
	//}
	// 订阅农业作物存储数据
	if token := mqttClient.Subscribe(viper.GetString("MQTT.ASTopic"), 0, service.ASDataToDB()); token.Wait() && token.Error() != nil {
		log.Fatal(viper.GetString("MQTT.ASTopic"), " 主题订阅失败: ", token.Error())
	}
	// 订阅监控日志
	if token := mqttClient.Subscribe(viper.GetString("MQTT.ICLogTopic"), 2, service.ICLogToDB()); token.Wait() && token.Error() != nil {
		log.Fatal(viper.GetString("MQTT.ICTopic"), " 主题订阅失败: ", token.Error())
	}
	// 订阅监控设备心跳检测
	if token := mqttClient.Subscribe(viper.GetString("MQTT.ICHBTopic"), 2, service.ICHeartBeat()); token.Wait() && token.Error() != nil {
		log.Fatal(viper.GetString("MQTT.ICTopic"), " 主题订阅失败: ", token.Error())
	}
}
