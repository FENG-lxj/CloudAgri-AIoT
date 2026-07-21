package AI

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type GetMonitorLogs struct {
	TimePeriod int32 `form:"TimePeriod" json:"TimePeriod"  binding:"required,min=1"` // 时间段，多少小时以内的数据
}

func (a GetMonitorLogs) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&a); err != nil {
		// 将表单参数验证器出现的错误直接交给错误翻译器统一处理即可
		response.ValidatorError(context, err)
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(a, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "UserShow表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.AI{}).GetMonitorLogs(extraAddBindDataContext)
	}
}
