package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/service/users/curd"
	"goskeleton/app/service/users/ipLocation"
	"goskeleton/app/service/users/phoneCaptcha"
	userstoken "goskeleton/app/service/users/token"
	"goskeleton/app/utils/response"
	"time"
)

type Users struct {
}

// 发送手机验证码
func (u *Users) Captcha(context *gin.Context) {
	phone := context.GetString(consts.ValidatorPrefix + "phone")
	if curd.CreateUserCurdFactory().IsUser(phone) {
		//生成验证码
		_, CreationTime := phoneCaptcha.GetPhoneCode(phone)
		// 获取当前时间戳
		currentTime := time.Now().Unix()
		// 计算CreationTime到当前时间的时间差
		duration := currentTime - CreationTime
		if duration >= 60 {
			code := phoneCaptcha.Create(phone)
			if code != "" {
				data := gin.H{
					"phone": phone,
					"time":  time.Now().Format(variable.DateFormat),
					"code":  code,
				}
				response.Success(context, consts.CurdStatusOkMsg, data)
				return
			}
			response.Fail(context, consts.CurdRegisterFailCode, consts.CurdRegisterFailMsg, "")
			return
		}
		response.Fail(context, consts.CurdRegisterRepeatCode, consts.CurdRegisterRepeatMsg, "")
		return
	}
	response.Fail(context, consts.CurdLoginNotUserCode, consts.CurdLoginNotUserMsg, "")
	return
}

// 用户登录
func (u *Users) Login(context *gin.Context) {
	phone := context.GetString(consts.ValidatorPrefix + "phone")
	inputCode := context.GetString(consts.ValidatorPrefix + "phoneCaptcha")
	code, _ := phoneCaptcha.GetPhoneCode(phone)
	if inputCode != code {
		if code == "" {
			response.Fail(context, consts.CurdVerifyCaptchaVoidCode, consts.CurdVerifyCaptchaVoidMsg, "")
			return
		}
		response.Fail(context, consts.CurdVerifyCaptchaFailCode, consts.CurdVerifyCaptchaFailMsg, "")
		return
	}
	userIP := context.ClientIP()
	userAddr := ipLocation.GetIPLocation(userIP)
	userModelFact := curd.CreateUserCurdFactory()
	userModel := userModelFact.Login(phone)
	isSuper := 0

	if userModel != nil {
		userTokenFactory := userstoken.CreateUserFactory()
		if userToken, err := userTokenFactory.GenerateToken(phone, variable.ConfigYml.GetInt64("Token.JwtTokenCreatedExpireAt")); err == nil {
			if userTokenFactory.RecordLoginToken(phone, time.Now().Format(variable.DateFormat), userToken, userIP) {
				if phone == variable.ConfigYml.GetString("User.Superuser") {
					isSuper = 1
				}
				data := gin.H{
					"phone":       phone,
					"isSuper":     isSuper,
					"headImg":     userModel.HeadImg,
					"token":       userToken,
					"createdTime": userModel.CreatedTime,
				}
				response.Success(context, consts.CurdStatusOkMsg, data)
				go userModel.UpdateUserloginInfo(userIP, userAddr, userModel.Phone)
				return
			}
		}
	}
	response.Fail(context, consts.CurdLoginFailCode, consts.CurdLoginFailMsg, "")
}

func (u *Users) RemoveUser(context *gin.Context) {
	RemovePhone := context.GetString(consts.ValidatorPrefix + "RemovePhone")
	if curd.CreateUserCurdFactory().RemoveUser(RemovePhone) {
		response.Success(context, consts.CurdStatusOkMsg, "")
		return
	}
	response.Fail(context, consts.CurdDeleteFailCode, consts.CurdDeleteFailMsg, "")
}

// 刷新用户token
func (u *Users) RefreshToken(context *gin.Context) {
	oldToken := context.GetString(consts.ValidatorPrefix + "token")
	userIP := context.ClientIP()
	userAddr := ipLocation.GetIPLocation(userIP)
	if newToken, ok := userstoken.CreateUserFactory().RefreshToken(oldToken, userIP, userAddr); ok {
		res := gin.H{
			"token": newToken,
		}
		response.Success(context, consts.CurdStatusOkMsg, res)
	} else {
		response.Fail(context, consts.CurdRefreshTokenFailCode, consts.CurdRefreshTokenFailMsg, "")
	}
}

// 添加新管理用户
func (u *Users) AddUser(context *gin.Context) {
	phone := context.GetString(consts.ValidatorPrefix + "phone")
	//fmt.Print(phone)
	userIP := context.ClientIP()
	userAddr := ipLocation.GetIPLocation(userIP)
	if curd.CreateUserCurdFactory().AddUser(phone, userIP, userAddr, time.Now().Format(variable.DateFormat), 0) {
		response.Success(context, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(context, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, "")
	}
}

func (u *Users) GetUserList(context *gin.Context) {
	UsersList := curd.CreateUsersCurdFactory().GetUserList()
	if UsersList != nil {
		response.Success(context, consts.CurdStatusOkMsg, UsersList)
		return
	}
	response.Fail(context, consts.CurdFetchFailCode, consts.CurdFetchFailMsg, "")
}
