package mapper

import (
	"IOTA_ZZ_server/model"
	"IOTA_ZZ_server/pkg/MySQL"
	"log"
)

func SaveICLog(datas model.ICLog) {
	sql := "INSERT INTO monitor_logs (ICID, Datetime, Level1Class, Log) VALUES (?, ?, ?, ?)"

	result := MySQL.GetSQL().Exec(sql, datas.ICID, datas.Datetime, datas.Level1Class, datas.Log)
	if result.Error != nil {
		log.Println("SaveICLog 日志信息存入数据库出错: ", result.Error)
	}
}
