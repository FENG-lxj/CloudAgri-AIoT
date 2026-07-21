package users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

// 常用的验证规则：
// required  必填；
// len=11 长度=11；
// min=3  如果是数字，验证的是数据范围，最小值为3，如果是文本，验证的是最小长度为3，
// max=6 如果是数字，验证的是数字最大值为6，如果是文本，验证的是最大长度为6
// mail 验证邮箱
// gt=3  对于文本就是长度>=3
// lt=6  对于文本就是长度<=6

type Captcha struct {
	BaseField
}

// 特别注意: 表单参数验证器结构体的函数，绝对不能绑定在指针上
// 这部分代码项目启动后会加载到容器，如果绑定在指针，一次请求之后，会造成容器中的代码段被污染

func (c Captcha) CheckParams(context *gin.Context) {
	//1.先按照验证器提供的基本语法，基本可以校验90%以上的不合格参数
	if err := context.ShouldBind(&c); err != nil {
		response.ValidatorError(context, err)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(c, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "UserRegister表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Users{}).Captcha(extraAddBindDataContext)
	}

}
