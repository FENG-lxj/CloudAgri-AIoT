package main

import (
	"IOTA_ZZ_server/config"
	"IOTA_ZZ_server/pkg/InfluxDB"
	"IOTA_ZZ_server/pkg/MyMQTT"
	"IOTA_ZZ_server/pkg/MySQL"
	"IOTA_ZZ_server/pkg/logo"
)

func main() {
	logo.Banner()       // 项目初始化    *
	config.Init()       // 初始化配置文件   *
	InfluxDB.InitDB()   // 初始化InfluxDB时序数据库
	MySQL.Init()        // 初始化MySQL数据库
	MyMQTT.Controller() // 启动MQTT控制器
	select {}           //阻塞
}
