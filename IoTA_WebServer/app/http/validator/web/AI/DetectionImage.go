package AI

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type DetectionImage struct {
	Id          int32 `form:"Id" json:"Id"  binding:"required,min=1"`                         //数据库表唯一id
	Level1Class int32 `form:"Level1Class" json:"Level1Class"  binding:"required,min=1,max=3"` //一级分类：农作物病虫害(1)、农场入侵(2)、火灾(3)
}

func (a DetectionImage) CheckParams(context *gin.Context) {
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
		(&web.AI{}).DetectionImage(extraAddBindDataContext)
	}
}
