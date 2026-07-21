package mapper

import (
	"IOTA_ZZ_server/pkg/MySQL"
	"log"
	"time"
)

//func SaveAEDevGPS(AEID int32, Longitude, Dimension float64) {
//	sql := "UPDATE dev_manage SET Longitude=?, Dimension=?, LatestNewsTime=? WHERE Types=? AND DevID=?"
//	res := MySQL.GetSQL().Exec(sql, Longitude, Dimension, time.Now().Format("2006-01-02 15:04:05"), "AE", AEID)
//	if res.Error != nil {
//		log.Println("SaveAEDevGPS 更新 AE 经纬度信息出错:", res.Error)
//	}
//}

func UpASDevLatestNewsTime(ASID int32) {
	sql := "UPDATE dev_manage SET LatestNewsTime=? WHERE Types=? AND DevID=?"
	res := MySQL.GetSQL().Exec(sql, time.Now().Format("2006-01-02 15:04:05"), "AS", ASID)
	if res.Error != nil {
		log.Println("UpASDevLatestNewsTime 更新 AS LatestNewsTime 信息出错:", res.Error)
	}
}

func UPICHeartBeat(ICID int32) {
	sql := "UPDATE dev_manage SET LatestNewsTime=? WHERE Types=? AND DevID=?"
	res := MySQL.GetSQL().Exec(sql, time.Now().Format("2006-01-02 15:04:05"), "IC", ICID)
	if res.Error != nil {
		log.Println("UPICHeartBeat 更新 IC LatestNewsTime 信息出错:", res.Error)
	}
}

func SelectDevExist(Types string, DevID int32) bool {
	var count int
	sql := "SELECT COUNT(*) FROM dev_manage WHERE Types=? AND DevID=?"
	result := MySQL.GetSQL().Raw(sql, Types, DevID).Scan(&count)
	if result.Error != nil {
		log.Println("SelectDevExist 查询数据库设备信息是否已存在出错: ", result.Error)
	}
	if count > 0 {
		return true
	} else {
		return false
	}
}

func InsertDevInfo(Types string, DevID int32) {
	sql := "INSERT INTO dev_manage(Types, DevID, Cropper, Longitude, Dimension) SELECT ?,?,?,?,? FROM DUAL WHERE NOT EXISTS (SELECT 1 FROM dev_manage WHERE Types=? AND DevID=?)"
	result := MySQL.GetSQL().Exec(sql, Types, DevID, "待填写", 0, 0, Types, DevID)
	if result.Error != nil {
		log.Println("InsertDevInfo 新建设备信息出错：", result.Error)
	}
}
