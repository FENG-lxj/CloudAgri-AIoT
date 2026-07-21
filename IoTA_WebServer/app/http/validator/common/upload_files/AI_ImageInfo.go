package upload_files

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/files"
	"goskeleton/app/utils/response"
	"strconv"
	"strings"
)

type AiImageInfo struct {
	ICID        int32  `form:"ICID" json:"ICID"  binding:"required,min=1"`               //摄像头设备ID
	Level1Class int32  `form:"Level1Class" json:"Level1Class"  binding:"required,min=1"` //一级分类：农作物病虫害(1)、农场入侵(2)、火灾(3)
	Level2Class string `form:"Level2Class" json:"Level2Class"  binding:"min=0"`          //二级分类：农作物：水稻、小麦...，入侵：动物入侵、人物入侵
	Level3Class string `form:"Level3Class" json:"Level3Class"  binding:"min=0"`          //三级分类：稻枯叶病...
	IdentifyID  string `form:"IdentifyID" json:"IdentifyID"  binding:"required,min=1"`   //识别ID
	Timestamp   int64  `form:"Timestamp" json:"Timestamp"  binding:"required,min=1"`     //时间戳(秒)
}

// 验证器语法
func (a AiImageInfo) CheckParams(context *gin.Context) {
	//基本的验证规则没有通过
	if err := context.ShouldBind(&a); err != nil {
		// 将表单参数验证器出现的错误直接交给错误翻译器统一处理即可
		response.ValidatorError(context, err)
		return
	}
	tmpFile, err := context.FormFile(variable.ConfigYml.GetString("AiImageSetting.UploadFileField")) //  file 是一个文件结构体（文件对象）
	var isPass bool
	//获取文件发生错误，可能上传了空文件等
	if err != nil {
		response.Fail(context, consts.FilesUploadFailCode, consts.FilesUploadFailMsg, err.Error())
		return
	}
	//超过系统设定的最大值：5M，tmpFile.Size 的单位是 bytes 和我们定义的文件单位M 比较，就需要将我们的单位*1024*1024(即2的20次方)，一步到位就是 << 20
	sizeLimit := variable.ConfigYml.GetInt64("AiImageSetting.Size")
	if tmpFile.Size > sizeLimit<<20 {
		response.Fail(context, consts.FilesUploadMoreThanMaxSizeCode, consts.FilesUploadMoreThanMaxSizeMsg+strconv.FormatInt(sizeLimit, 10)+"M", "")
		return
	}
	//不允许的文件mime类型
	if fp, err := tmpFile.Open(); err == nil {
		mimeType := files.GetFilesMimeByFp(fp)

		for _, value := range variable.ConfigYml.GetStringSlice("AiImageSetting.AllowMimeType") {
			if strings.ReplaceAll(value, " ", "") == strings.ReplaceAll(mimeType, " ", "") {
				isPass = true
				break
			}
		}
		_ = fp.Close()
	} else {
		response.ErrorSystem(context, consts.ServerOccurredErrorMsg, "")
		return
	}
	//凡是存在相等的类型，通过验证，调用控制器
	if !isPass {
		response.Fail(context, consts.FilesUploadMimeTypeFailCode, consts.FilesUploadMimeTypeFailMsg, "")
	} else {
		//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
		extraAddBindDataContext := data_transfer.DataAddContext(a, consts.ValidatorPrefix, context)
		if extraAddBindDataContext == nil {
			response.ErrorSystem(context, "UserShow表单验证器json化失败", "")
		} else {
			// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
			(&web.Upload{}).AiImageInfo(extraAddBindDataContext)
		}
	}
}
