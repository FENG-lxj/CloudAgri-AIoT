package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/service/server_Info"
	"goskeleton/app/service/users/curd"
	"goskeleton/app/utils/response"
	"strconv"
)

type BigData struct {
}

// 获取AI实时生态数据农业生产建议
func (b *BigData) GetAEdata(context *gin.Context) {
	AEID := int32(context.GetFloat64(consts.ValidatorPrefix + "AEID"))
	StartTime := context.GetString(consts.ValidatorPrefix + "StartTime")
	AEDatas := curd.CreateAEDataCurdFactory().GetAEDatas(AEID, StartTime)
	if AEDatas != nil {
		response.Success(context, consts.CurdStatusOkMsg, AEDatas)
		return
	}
	response.Fail(context, consts.CurdFetchFailCode, consts.CurdFetchFailMsg, "")
}

// 获取AI实时生态数据农业生产建议
func (b *BigData) GetASdata(context *gin.Context) {
	ASID := int32(context.GetFloat64(consts.ValidatorPrefix + "ASID"))
	StartTime := context.GetString(consts.ValidatorPrefix + "StartTime")
	ASDatas := curd.CreateASDataCurdFactory().GetASDatas(ASID, StartTime)
	if ASDatas != nil {
		response.Success(context, consts.CurdStatusOkMsg, ASDatas)
		return
	}
	response.Fail(context, consts.CurdFetchFailCode, consts.CurdFetchFailMsg, "")
}

// 获取设备信息
func (b *BigData) GetDevInfo(context *gin.Context) {
	Types := context.GetString(consts.ValidatorPrefix + "Types")
	DevManageModels, err := curd.CreateDevManageCurdFactory().GetDevInfo(Types)
	if err == nil {
		response.Success(context, consts.CurdStatusOkMsg, DevManageModels)
		return
	}
	response.Fail(context, consts.CurdFetchFailCode, consts.CurdFetchFailMsg, "")
}

// 修改设备信息
func (b *BigData) UpDev(context *gin.Context) {
	Types := context.GetString(consts.ValidatorPrefix + "Types")
	DevID := int32(context.GetFloat64(consts.ValidatorPrefix + "DevID"))
	Longitude := context.GetFloat64(consts.ValidatorPrefix + "Longitude")
	Dimension := context.GetFloat64(consts.ValidatorPrefix + "Dimension")
	Cropper := context.GetString(consts.ValidatorPrefix + "Cropper")
	if curd.CreateDevManageCurdFactory().UpDev(Types, DevID, Longitude, Dimension, Cropper) == nil {
		response.Success(context, consts.CurdStatusOkMsg, "")
		return
	}
	response.Fail(context, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, "")
}

// 删除设备
func (b *BigData) RemoveDev(context *gin.Context) {
	Types := context.GetString(consts.ValidatorPrefix + "Types")
	DevID := int32(context.GetFloat64(consts.ValidatorPrefix + "DevID"))
	if curd.CreateDevManageCurdFactory().RemoveDev(Types, DevID) == nil {
		response.Success(context, consts.CurdStatusOkMsg, "")
		return
	}
	response.Fail(context, consts.CurdDeleteFailCode, consts.CurdDeleteFailMsg, "")
}

// 获取服务器内存占比和CPU占比
func (b *BigData) GetMemCPU(context *gin.Context) {
	cpuPer, err := server_Info.GetCpuPercent()
	if err != nil {
		response.Fail(context, consts.CurdFetchFailCode, consts.CurdFetchFailMsg, "")
	} else {
		cpuPer, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", cpuPer), 64)
	}
	memPer, err := server_Info.GetMemPercent()
	if err != nil {
		response.Fail(context, consts.CurdFetchFailCode, consts.CurdFetchFailMsg, "")
	} else {
		memPer, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", memPer), 64)
	}

	data := gin.H{
		"CPUPer": cpuPer,
		"MemPer": memPer,
	}
	response.Success(context, consts.CurdStatusOkMsg, data)
}
