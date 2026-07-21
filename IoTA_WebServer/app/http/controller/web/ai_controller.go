package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/model"
	"goskeleton/app/service/users/curd"
	"goskeleton/app/utils/response"
	"time"
)

type AI struct {
}

// 获取AI实时生态数据农业生产建议
func (a *AI) AEProposalAI(context *gin.Context) {
	AEID := int32(context.GetFloat64(consts.ValidatorPrefix + "AEID"))
	Proposal := curd.CreateAICurdFactory().GetAIProposalAE(AEID)
	if Proposal != "" {
		data := gin.H{
			"time":       time.Now().Format(variable.DateFormat),
			"AIProposal": Proposal,
		}
		response.Success(context, consts.CurdStatusOkMsg, data)
		return
	}
	response.Fail(context, consts.CurdAIFailCode, consts.CurdAIFailMsg, "")
}

// 获取AI实时农作物存储环境数据相关管理建议
func (a *AI) ASProposalAI(context *gin.Context) {
	ASID := int32(context.GetFloat64(consts.ValidatorPrefix + "ASID"))
	Proposal := curd.CreateAICurdFactory().GetAIProposalAS(ASID)
	if Proposal != "" {
		data := gin.H{
			"time":       time.Now().Format(variable.DateFormat),
			"AIProposal": Proposal,
		}
		response.Success(context, consts.CurdStatusOkMsg, data)
		return
	}
	response.Fail(context, consts.CurdAIFailCode, consts.CurdAIFailMsg, "")
}

// 向AI农业专家提问
func (a *AI) AgricultureExpertAI(context *gin.Context) {
	Issue := context.GetString(consts.ValidatorPrefix + "Issue")
	Reply := curd.CreateAICurdFactory().AgricultureExpertAI(Issue)
	if Reply != "" {
		if Reply == "429" {
			response.Fail(context, consts.CurdChatGPT429Code, consts.CurdChatGPT429Msg, "")
			return
		}
		data := gin.H{
			"time":    time.Now().Format(variable.DateFormat),
			"AIReply": Reply,
		}
		response.Success(context, consts.CurdStatusOkMsg, data)
		return
	}
	response.Fail(context, consts.CurdAIFailCode, consts.CurdAIFailMsg, "")
}

// 更换APIKEY
func (a *AI) UpChatGPTKey(context *gin.Context) {
	APIKEY := context.GetString(consts.ValidatorPrefix + "APIKEY")
	if curd.CreateAICurdFactory().UpChatGPTKey(APIKEY) {
		response.Success(context, consts.CurdStatusOkMsg, "")
		return
	}
	response.Fail(context, consts.CurdUpFailCode, consts.CurdUpFailMsg, "")
}

// 检测信息
func (a *AI) DetectionInfo(context *gin.Context) {
	ICID := int32(context.GetFloat64(consts.ValidatorPrefix + "ICID"))
	Level1Class := int32(context.GetFloat64(consts.ValidatorPrefix + "Level1Class"))
	StartNum := int32(context.GetFloat64(consts.ValidatorPrefix + "StartNum"))
	Num := int32(context.GetFloat64(consts.ValidatorPrefix + "Num"))

	switch Level1Class {
	case 1:
		if CropperModels, err := model.CreateCropperFactory("").CropperDetectionInfo(ICID, StartNum, Num); err == nil {
			response.Success(context, consts.CurdStatusOkMsg, CropperModels)
			return
		}
		response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	case 2:
		if InvadeModels, err := model.CreateInvadeFactory("").InvadeDetectionInfo(ICID, StartNum, Num); err == nil {
			response.Success(context, consts.CurdStatusOkMsg, InvadeModels)
			return
		}
		response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	case 3:
		if FireModels, err := model.CreateFireFactory("").FireDetectionInfo(ICID, StartNum, Num); err == nil {
			response.Success(context, consts.CurdStatusOkMsg, FireModels)
			return
		}
		response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	}

}

// 检测图片
func (a *AI) DetectionImage(context *gin.Context) {
	Id := int32(context.GetFloat64(consts.ValidatorPrefix + "Id"))
	Level1Class := int32(context.GetFloat64(consts.ValidatorPrefix + "Level1Class"))

	switch Level1Class {
	case 1:
		if CropperImage := model.CreateCropperFactory("").CropperDetectionImage(Id); CropperImage != "" {
			context.File(CropperImage)
			return
		}
		response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	case 2:
		if InvadeImage := model.CreateInvadeFactory("").InvadeDetectionImage(Id); InvadeImage != "" {
			context.File(InvadeImage)
			return
		}
		response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	case 3:
		if FireImage := model.CreateFireFactory("").FireDetectionImage(Id); FireImage != "" {
			context.File(FireImage)
			return
		}
		response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	}
}

func (a *AI) GetMonitorLogs(context *gin.Context) {
	TimePeriod := int32(context.GetFloat64(consts.ValidatorPrefix + "TimePeriod"))
	if LogsModels, err := model.CreateLogsFactory("").MonitorLogs(TimePeriod); err == nil {
		response.Success(context, consts.CurdStatusOkMsg, LogsModels)
		return
	}
	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
}

func (a *AI) PestProposal(context *gin.Context) {
	Pest := context.GetString(consts.ValidatorPrefix + "Pest")
	if Proposal := model.CreatePestProposalFactory("").PestProposal(Pest); Proposal != "" {
		data := gin.H{
			"Pest":     Pest,
			"Proposal": Proposal,
		}
		response.Success(context, consts.CurdStatusOkMsg, data)
		return
	}
	response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
}
