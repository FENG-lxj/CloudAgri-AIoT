package users

type BaseField struct {
	Phone string `form:"phone" json:"phone"  binding:"required,min=11,max=11"` // 必填、对于手机号,表示它的长度=11位
}

type BaseCaptcha struct {
	PhoneCaptcha string `form:"phoneCaptcha" json:"phoneCaptcha"  binding:"required,len=6"` // 必填、对于手机验证码,表示它的长度=6位
}
