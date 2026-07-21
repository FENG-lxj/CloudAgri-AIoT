package model

import (
	"go.uber.org/zap"
	"goskeleton/app/global/variable"
	"time"
)

func CreateDevManageFactory(sqlType string) *DevManageModel {
	return &DevManageModel{GormM: GormM{DB: UseDbConn(sqlType)}}
}

type DevManageModel struct {
	GormM
	Id             int64   `gorm:"primaryKey" json:"-"`
	Types          string  `gorm:"column:Types" json:"-"`
	DevID          int32   `gorm:"column:DevID" json:"DevID"`
	Cropper        string  `gorm:"column:Cropper" json:"Cropper"`
	Longitude      float64 `gorm:"column:Longitude" json:"Longitude"` //经度
	Dimension      float64 `gorm:"column:Dimension" json:"Dimension"` //纬度
	LatestNewsTime string  `gorm:"column:LatestNewsTime" json:"-"`    //最新信息时间
	Status         int     `json:"Status"`                            //设备当前状态
}

type DevManageModels struct {
	DevManageModel []DevManageModel `json:"DevInfo"`
}

func (dm *DevManageModel) GetDevInfo(Types string) (DevManageModels, error) {
	var DevManageModels DevManageModels
	sql := "SELECT DevID,Cropper,Longitude,Dimension,LatestNewsTime FROM dev_manage WHERE Types=?"
	rows, err := dm.Raw(sql, Types).Rows() // 执行查询语句，返回查询结果的行数据集
	if err != nil {
		variable.ZapLog.Error("GetDevInfo 查询检测信息列表出错:", zap.Error(err))
		return DevManageModels, err
	}
	defer rows.Close()

	for rows.Next() {
		var d DevManageModel
		err := rows.Scan(&d.DevID, &d.Cropper, &d.Longitude, &d.Dimension, &d.LatestNewsTime)
		if err != nil {
			variable.ZapLog.Error("GetDevInfo 读取数据出错:", zap.Error(err))
			return DevManageModels, err
		}
		// 将日期时间字符串解析为时间，并指定时区
		loc, _ := time.LoadLocation("Local")
		t, err := time.ParseInLocation("2006-1-2 15:04:05", d.LatestNewsTime, loc)
		if err != nil {
			variable.ZapLog.Error("GetDevInfo 日期时间格式转换出错: ", zap.Error(err))
			return DevManageModels, err
		}
		// 判断当前时间差是否小于 * 分钟
		if time.Now().In(loc).Sub(t) < time.Duration(variable.ConfigYml.GetInt("Dev.DeviceTimeoutPeriod"))*time.Minute {
			d.Status = 1
		} else {
			d.Status = 0
		}

		DevManageModels.DevManageModel = append(DevManageModels.DevManageModel, d)
	}

	return DevManageModels, nil
}

func (dm *DevManageModel) UpDev(Types string, DevID int32, Longitude, Dimension float64, Cropper string) error {
	sql := "UPDATE dev_manage SET Longitude=?,Dimension=?,Cropper=? WHERE Types=? AND DevID=?"
	result := dm.Exec(sql, Longitude, Dimension, Cropper, Types, DevID)
	if result.Error != nil {
		variable.ZapLog.Error("UpDev 更新设备信息出错：", zap.Error(result.Error))
		return result.Error
	}
	return nil
}

func (dm *DevManageModel) RemoveDev(Types string, DevID int32) error {
	sql := "DELETE FROM dev_manage WHERE Types=? AND DevID=?"
	result := dm.Exec(sql, Types, DevID)
	if result.Error != nil {
		variable.ZapLog.Error("RemoveDev 删除设备出错：", zap.Error(result.Error))
		return result.Error
	}
	return nil
}
