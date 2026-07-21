package config

import (
	"github.com/spf13/viper"
	"log"
)

func Init() {
	// 设置配置文件的名字
	viper.SetConfigName("config")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 添加配置文件的路径，指定 config 目录下寻找
	viper.AddConfigPath("./config")
	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("配置文件初始化出错:", err)
	}
}
