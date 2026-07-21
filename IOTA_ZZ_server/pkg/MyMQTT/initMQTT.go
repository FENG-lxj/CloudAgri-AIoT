package MyMQTT

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"log"
	"time"
)

func Init() *mqtt.ClientOptions {
	// 设置MQTT客户端连接参数
	opts := mqtt.NewClientOptions()
	opts.AddBroker(viper.GetString("MQTT.Broker"))
	opts.SetClientID(viper.GetString("MQTT.ClientID"))
	opts.SetUsername(viper.GetString("MQTT.Username"))
	opts.SetPassword(viper.GetString("MQTT.Password"))
	// 设置掉线重连的处理函数
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Println("MQTT连接断开，正在尝试重连...")
		for {
			// 连接MQTT服务器
			if token := client.Connect(); token.Wait() && token.Error() == nil {
				log.Println("MQTT重新连接成功")
				break
			}
			time.Sleep(5 * time.Second)
		}
	})
	return opts
}
