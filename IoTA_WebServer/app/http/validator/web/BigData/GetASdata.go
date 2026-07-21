package BigData

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type GetASdata struct {
	ASID      int32  `form:"ASID" json:"ASID"  binding:"required,min=0"` // 必填、对于商品id,表示它的数字>=1
	StartTime string `form:"StartTime" json:"StartTime"  binding:"required,min=0"`
}

func (g GetASdata) CheckParams(context *gin.Context) {
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
		(&web.BigData{}).GetASdata(extraAddBindDataContext)
	}
}
