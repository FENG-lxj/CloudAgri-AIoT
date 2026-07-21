package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/model"
	"goskeleton/app/service/upload_file"
	"goskeleton/app/utils/response"
)

type Upload struct {
}

// 开始上传
func (u *Upload) StartUpload(context *gin.Context) {
	savePath := variable.BasePath + variable.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath")
	if r, finnalSavePath := upload_file.Upload(context, savePath); r == true {
		response.Success(context, consts.CurdStatusOkMsg, finnalSavePath)
	} else {
		response.Fail(context, consts.FilesUploadFailCode, consts.FilesUploadFailMsg, "")
	}
}

func (u *Upload) AiImageInfo(context *gin.Context) {
	ICID := int32(context.GetFloat64(consts.ValidatorPrefix + "ICID"))
	Level1Class := int32(context.GetFloat64(consts.ValidatorPrefix + "Level1Class"))
	Level2Class := context.GetString(consts.ValidatorPrefix + "Level2Class")
	Level3Class := context.GetString(consts.ValidatorPrefix + "Level3Class")
	IdentifyID := context.GetString(consts.ValidatorPrefix + "IdentifyID")
	Timestamp := int64(context.GetFloat64(consts.ValidatorPrefix + "Timestamp"))
	savePath := variable.BasePath + variable.ConfigYml.GetString("AiImageSetting.UploadFileSavePath")

	switch Level1Class {
	case 1:
		if r, finnalSavePath := upload_file.UploadAiImage(context, savePath, "Cropper", IdentifyID); r == true {
			if model.CreateCropperFactory("").InsertCropperPestInfo(ICID, Level2Class, Level3Class, IdentifyID, Timestamp, finnalSavePath) {
				response.Success(context, consts.CurdStatusOkMsg, "")
			} else {
				response.Fail(context, consts.CurdAddAiImageInfoFailCode, consts.CurdAddAiImageInfoFailMsg, "")
			}
		} else {
			response.Fail(context, consts.FilesUploadFailCode, consts.FilesUploadFailMsg, "")
		}
	case 2:
		if r, finnalSavePath := upload_file.UploadAiImage(context, savePath, "Invade", IdentifyID); r == true {
			if model.CreateInvadeFactory("").InsertInvadeInfo(ICID, Level2Class, IdentifyID, Timestamp, finnalSavePath) {
				response.Success(context, consts.CurdStatusOkMsg, "")
			} else {
				response.Fail(context, consts.CurdAddAiImageInfoFailCode, consts.CurdAddAiImageInfoFailMsg, "")
			}
		} else {
			response.Fail(context, consts.FilesUploadFailCode, consts.FilesUploadFailMsg, "")
		}
	case 3:
		if r, finnalSavePath := upload_file.UploadAiImage(context, savePath, "Fire", IdentifyID); r == true {
			if model.CreateFireFactory("").InsertFireInfo(ICID, IdentifyID, Timestamp, finnalSavePath) {
				response.Success(context, consts.CurdStatusOkMsg, "")
			} else {
				response.Fail(context, consts.CurdAddAiImageInfoFailCode, consts.CurdAddAiImageInfoFailMsg, "")
			}
		} else {
			response.Fail(context, consts.FilesUploadFailCode, consts.FilesUploadFailMsg, "")
		}
	default:
		response.Fail(context, consts.CurdAddAiImageInfoFailCode, consts.CurdAddAiImageInfoFailMsg, "")
	}

}
