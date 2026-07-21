package model

import (
	"go.uber.org/zap"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/redis_factory"
	"time"
)

func CreateAIFactory(sqlType string) *AIModel {
	return &AIModel{GormM: GormM{DB: UseDbConn(sqlType)}}
}

func CreateAPIKEYFactory(sqlType string) *APIKEYModel {
	return &APIKEYModel{GormM: GormM{DB: UseDbConn(sqlType)}}
}

func CreateCropperFactory(sqlType string) *CropperModel {
	return &CropperModel{GormM: GormM{DB: UseDbConn(sqlType)}}
}

func CreateInvadeFactory(sqlType string) *InvadeModel {
	return &InvadeModel{GormM: GormM{DB: UseDbConn(sqlType)}}
}

func CreateFireFactory(sqlType string) *FireModel {
	return &FireModel{GormM: GormM{DB: UseDbConn(sqlType)}}
}

func CreateLogsFactory(sqlType string) *LogsModel {
	return &LogsModel{GormM: GormM{DB: UseDbConn(sqlType)}}
}

func CreatePestProposalFactory(sqlType string) *PestProposalModel {
	return &PestProposalModel{GormM: GormM{DB: UseDbConn(sqlType)}}
}

type AIModel struct {
	GormM
	Id        int64   `gorm:"primaryKey" json:"id"`
	Types     string  `gorm:"column:Types" json:"Types"`
	DevID     int32   `gorm:"column:DevID" json:"DevID"`
	Cropper   string  `gorm:"column:Cropper" json:"Cropper"`
	Longitude float64 `gorm:"column:Longitude" json:"Longitude"` //经度
	Dimension float64 `gorm:"column:Dimension" json:"Dimension"` //纬度
}

type APIKEYModel struct {
	GormM
	Id     int64  `gorm:"primaryKey" json:"id"`
	APIKEY string `gorm:"column:APIKEY" json:"APIKEY"`
	Status int    `gorm:"column:status" json:"status"`
}

type CropperModel struct {
	GormM
	Id         int64  `gorm:"primaryKey" json:"id"`
	ICID       int32  `gorm:"column:ICID" json:"-"`
	Cropper    string `gorm:"column:Cropper" json:"-"`
	Pest       string `gorm:"column:Pest" json:"Pest"`
	IdentifyID string `gorm:"column:IdentifyID" json:"IdentifyID"`
	Datetime   string `gorm:"column:Datetime" json:"Datetime"`
	ImagePath  string `gorm:"column:ImagePath" json:"-"`
}

type CropperModels struct {
	CropperModel []CropperModel `json:"DetectionInfo"`
}

type InvadeModel struct {
	GormM
	Id          int64  `gorm:"primaryKey" json:"id"`
	ICID        int32  `gorm:"column:ICID" json:"-"`
	InvadeClass string `gorm:"column:InvadeClass" json:"InvadeClass"`
	IdentifyID  string `gorm:"column:IdentifyID" json:"IdentifyID"`
	Datetime    string `gorm:"column:Datetime" json:"Datetime"`
	ImagePath   string `gorm:"column:ImagePath" json:"-"`
}

type InvadeModels struct {
	InvadeModel []InvadeModel `json:"DetectionInfo"`
}

type FireModel struct {
	GormM
	Id         int64  `gorm:"primaryKey" json:"id"`
	ICID       int32  `gorm:"column:ICID" json:"-"`
	IdentifyID string `gorm:"column:IdentifyID" json:"IdentifyID"`
	Datetime   string `gorm:"column:Datetime" json:"Datetime"`
	ImagePath  string `gorm:"column:ImagePath" json:"-"`
}

type FireModels struct {
	FireModel []FireModel `json:"DetectionInfo"`
}

type LogsModel struct {
	GormM
	Id          int64  `gorm:"primaryKey" json:"-"`
	ICID        int32  `gorm:"ICID" json:"ICID"`
	Datetime    string `gorm:"column:Datetime" json:"Datetime"`
	Level1Class int    `gorm:"Level1Class" json:"Level1Class"`
	Log         string `gorm:"column:Log" json:"Log"`
}

type LogsModels struct {
	LogsModel []LogsModel `json:"MonitorLogs"`
}

type PestProposalModel struct {
	GormM
	Id       int64  `gorm:"primaryKey" json:"-"`
	Pest     string `gorm:"column:Pest" json:"Pest"`
	Proposal string `gorm:"column:Proposal" json:"Proposal"`
}

// 根据生态设备ID查询种植农作物种类
func (a *AIModel) QueryCropsAE(AEID int32) (string, error) {
	sql := "SELECT `Cropper` FROM  `dev_manage`  WHERE Types='AE' AND DevID=? LIMIT 1"
	result := a.Raw(sql, AEID).First(a)
	if result.Error == nil {
		return a.Cropper, nil
	} else {
		return "", result.Error
	}
}

// 根据生态设备ID查询种植农作物种类
func (a *AIModel) QueryCropsAS(ASID int32) (string, error) {
	sql := "SELECT `Cropper` FROM  `dev_manage`  WHERE Types='AS' AND DevID=? LIMIT 1"
	result := a.Raw(sql, ASID).First(a)
	if result.Error == nil {
		return a.Cropper, nil
	} else {
		return "", result.Error
	}
}

func (api *APIKEYModel) UpChatGPTKey(APIKEY string) bool {
	sql := "UPDATE apikey SET status=0 WHERE status=1"
	result := api.Exec(sql)
	if result.Error != nil {
		variable.ZapLog.Error("UpChatGPTKey 更新旧APIKEY状态出错：", zap.Error(result.Error))
		return false
	}
	sql2 := "INSERT INTO apikey(APIKEY,status) SELECT ?,? FROM DUAL WHERE NOT EXISTS (SELECT 1 FROM apikey WHERE APIKEY=?)"
	result2 := api.Exec(sql2, APIKEY, 1, APIKEY)
	if result2.RowsAffected > 0 {
		// 从连接池获取一个连接
		redisClient := redis_factory.GetOneRedisClient()
		// 清除Redis缓存
		_, err := redisClient.Execute("del", "APIKEY")
		if err != nil {
			variable.ZapLog.Error("UpChatGPTKey 清除旧的APIKEY缓存出错：", zap.Error(err))
			redisClient.ReleaseOneRedisClient()
			return false
		}
		// 关闭连接
		redisClient.ReleaseOneRedisClient()
		return true
	} else {
		return false
	}
}

func (api *APIKEYModel) GetChatGPTAPIKey() string {
	sql := "SELECT APIKEY FROM apikey WHERE status=1 LIMIT 1"
	result := api.Raw(sql).First(api)
	if result.Error == nil {
		return api.APIKEY
	} else {
		variable.ZapLog.Error("GetChatGPTAPIKey 从数据库读取APIKEY出错：", zap.Error(result.Error))
		return ""
	}
}

func (cm *CropperModel) InsertCropperPestInfo(ICID int32, Cropper string, Pest string, IdentifyID string, Timestamp int64, finnalSavePath string) bool {
	sql := "INSERT INTO cropper_image(ICID, Cropper, Pest, IdentifyID, Datetime, ImagePath) SELECT ?,?,?,?,?,? FROM DUAL WHERE NOT EXISTS (SELECT 1 FROM cropper_image WHERE IdentifyID=?)"
	result := cm.Exec(sql, ICID, Cropper, Pest, IdentifyID, time.Unix(Timestamp, 0).Format(variable.DateFormat), finnalSavePath, IdentifyID)
	if result.RowsAffected > 0 {
		return true
	} else {
		variable.ZapLog.Error("InsertCropperPestInfo 存入病虫害信息出错：", zap.Error(result.Error))
		return false
	}
}

func (im *InvadeModel) InsertInvadeInfo(ICID int32, InvadeClass string, IdentifyID string, Timestamp int64, finnalSavePath string) bool {
	sql := "INSERT INTO invade_image(ICID, InvadeClass, IdentifyID, Datetime, ImagePath) SELECT ?,?,?,?,? FROM DUAL WHERE NOT EXISTS (SELECT 1 FROM invade_image WHERE IdentifyID=?)"
	result := im.Exec(sql, ICID, InvadeClass, IdentifyID, time.Unix(Timestamp, 0).Format(variable.DateFormat), finnalSavePath, IdentifyID)
	if result.RowsAffected > 0 {
		return true
	} else {
		variable.ZapLog.Error("InsertInvadeInfo 存入农场入侵信息出错：", zap.Error(result.Error))
		return false
	}
}

func (fm *FireModel) InsertFireInfo(ICID int32, IdentifyID string, Timestamp int64, finnalSavePath string) bool {
	sql := "INSERT INTO fire_image(ICID, IdentifyID, Datetime, ImagePath) SELECT ?,?,?,? FROM DUAL WHERE NOT EXISTS (SELECT 1 FROM fire_image WHERE IdentifyID=?)"
	result := fm.Exec(sql, ICID, IdentifyID, time.Unix(Timestamp, 0).Format(variable.DateFormat), finnalSavePath, IdentifyID)
	if result.RowsAffected > 0 {
		return true
	} else {
		variable.ZapLog.Error("InsertFireInfo 存入火灾检测信息出错：", zap.Error(result.Error))
		return false
	}
}

func (cm *CropperModel) CropperDetectionInfo(ICID, StartNum, Num int32) (CropperModels, error) {
	var CropperModels CropperModels
	sql := "SELECT id,Pest,IdentifyID,Datetime FROM cropper_image WHERE ICID=? ORDER BY Datetime DESC LIMIT ? OFFSET ?"
	rows, err := cm.Raw(sql, ICID, Num, StartNum).Rows() // 执行查询语句，返回查询结果的行数据集
	if err != nil {
		variable.ZapLog.Error("CropperDetectionInfo 查询检测信息列表出错:", zap.Error(err))
		return CropperModels, err
	}
	defer rows.Close()

	for rows.Next() {
		var c CropperModel
		err := rows.Scan(&c.Id, &c.Pest, &c.IdentifyID, &c.Datetime)
		if err != nil {
			variable.ZapLog.Error("CropperDetectionInfo 读取数据出错:", zap.Error(err))
			return CropperModels, err
		}
		CropperModels.CropperModel = append(CropperModels.CropperModel, c)
	}

	return CropperModels, nil
}

func (im *InvadeModel) InvadeDetectionInfo(ICID, StartNum, Num int32) (InvadeModels, error) {
	var InvadeModels InvadeModels
	sql := "SELECT id,InvadeClass,IdentifyID,Datetime FROM invade_image WHERE ICID=? ORDER BY Datetime DESC LIMIT ? OFFSET ?"
	rows, err := im.Raw(sql, ICID, Num, StartNum).Rows() // 执行查询语句，返回查询结果的行数据集
	if err != nil {
		variable.ZapLog.Error("InvadeDetectionInfo 查询检测信息列表出错:", zap.Error(err))
		return InvadeModels, err
	}
	defer rows.Close()

	for rows.Next() {
		var i InvadeModel
		err := rows.Scan(&i.Id, &i.InvadeClass, &i.IdentifyID, &i.Datetime)
		if err != nil {
			variable.ZapLog.Error("InvadeDetectionInfo 读取数据出错:", zap.Error(err))
			return InvadeModels, err
		}
		InvadeModels.InvadeModel = append(InvadeModels.InvadeModel, i)
	}

	return InvadeModels, nil
}

func (fm *FireModel) FireDetectionInfo(ICID, StartNum, Num int32) (FireModels, error) {
	var FireModels FireModels
	sql := "SELECT id,IdentifyID,Datetime FROM fire_image WHERE ICID=? ORDER BY Datetime DESC LIMIT ? OFFSET ?"
	rows, err := fm.Raw(sql, ICID, Num, StartNum).Rows() // 执行查询语句，返回查询结果的行数据集
	if err != nil {
		variable.ZapLog.Error("FireDetectionInfo 查询检测信息列表出错:", zap.Error(err))
		return FireModels, err
	}
	defer rows.Close()

	for rows.Next() {
		var f FireModel
		err := rows.Scan(&f.Id, &f.IdentifyID, &f.Datetime)
		if err != nil {
			variable.ZapLog.Error("FireDetectionInfo 读取数据出错:", zap.Error(err))
			return FireModels, err
		}
		FireModels.FireModel = append(FireModels.FireModel, f)
	}

	return FireModels, nil
}

func (cm *CropperModel) CropperDetectionImage(Id int32) string {
	sql := "SELECT ImagePath FROM cropper_image WHERE id=?"
	result := cm.Raw(sql, Id).First(cm)
	if result.Error == nil {
		return cm.ImagePath
	} else {
		variable.ZapLog.Error("CropperDetectionImage 从数据库查询图片路径出错：", zap.Error(result.Error))
		return ""
	}
}

func (im *InvadeModel) InvadeDetectionImage(Id int32) string {
	sql := "SELECT ImagePath FROM invade_image WHERE id=?"
	result := im.Raw(sql, Id).First(im)
	if result.Error == nil {
		return im.ImagePath
	} else {
		variable.ZapLog.Error("InvadeDetectionImage 从数据库查询图片路径出错：", zap.Error(result.Error))
		return ""
	}
}

func (fm *FireModel) FireDetectionImage(Id int32) string {
	sql := "SELECT ImagePath FROM fire_image WHERE id=?"
	result := fm.Raw(sql, Id).First(fm)
	if result.Error == nil {
		return fm.ImagePath
	} else {
		variable.ZapLog.Error("FireDetectionImage 从数据库查询图片路径出错：", zap.Error(result.Error))
		return ""
	}
}

func (lm *LogsModel) MonitorLogs(TimePeriod int32) (LogsModels, error) {
	var LogsModels LogsModels
	sql := "SELECT ICID,Datetime,Level1Class,Log FROM monitor_logs WHERE Datetime >= DATE_SUB(NOW(), INTERVAL ? HOUR)"
	rows, err := lm.Raw(sql, TimePeriod).Rows() // 执行查询语句，返回查询结果的行数据集
	if err != nil {
		variable.ZapLog.Error("MonitorLogs 查询检测信息列表出错:", zap.Error(err))
		return LogsModels, err
	}
	defer rows.Close()

	for rows.Next() {
		var l LogsModel
		err := rows.Scan(&l.ICID, &l.Datetime, &l.Level1Class, &l.Log)
		if err != nil {
			variable.ZapLog.Error("MonitorLogs 读取数据出错:", zap.Error(err))
			return LogsModels, err
		}
		LogsModels.LogsModel = append(LogsModels.LogsModel, l)
	}

	return LogsModels, nil
}

func (ppm *PestProposalModel) PestProposal(Pest string) string {
	sql := "SELECT Proposal FROM pest_proposal WHERE Pest=?"
	result := ppm.Raw(sql, Pest).First(ppm)
	if result.Error == nil {
		return ppm.Proposal
	} else {
		variable.ZapLog.Error("PestProposal 从数据库查询病虫害建议出错：", zap.Error(result.Error))
		return ""
	}
}
