package BigData

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type UpDev struct {
	Types     string  `form:"Types" json:"Types"  binding:"required,min=0"`
	DevID     int32   `form:"DevID" json:"DevID"  binding:"required,min=0"`
	Longitude float64 `form:"Longitude" json:"Longitude"  binding:"required,min=0"`
	Dimension float64 `form:"Dimension" json:"Dimension"  binding:"required,min=0"`
	Cropper   string  `form:"Cropper" json:"Cropper"  binding:"required,min=0"`
}

func (g UpDev) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&g); err != nil {
		// 将表单参数验证器出现的错误直接交给错误翻译器统一处理即可
		response.ValidatorError(context, err)
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(g, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "UserShow表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.BigData{}).UpDev(extraAddBindDataContext)
	}
}
