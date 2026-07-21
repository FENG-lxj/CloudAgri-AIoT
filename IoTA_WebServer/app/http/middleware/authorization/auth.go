package authorization

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/service/users/ipLocation"
	userstoken "goskeleton/app/service/users/token"
	"goskeleton/app/utils/response"
	"strings"
)

type HeaderParams struct {
	Authorization string `header:"Authorization" binding:"required,min=20"`
}

// CheckTokenAuth 检查token完整性、有效性中间件
func CheckTokenAuth() gin.HandlerFunc {
	return func(context *gin.Context) {

		headerParams := HeaderParams{}

		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			response.TokenErrorParam(context, consts.JwtTokenMustValid+err.Error())
			return
		}
		token := strings.Split(headerParams.Authorization, " ")
		if len(token) == 2 && len(token[1]) >= 20 {
			tokenIsEffective := userstoken.CreateUserFactory().IsEffective(token[1])
			if tokenIsEffective {
				if customToken, err := userstoken.CreateUserFactory().ParseToken(token[1]); err == nil {
					//key := variable.ConfigYml.GetString("Token.BindContextKeyName")
					// token验证通过，同时绑定在请求上下文
					context.Set("phone", customToken.Phone)
				}
				context.Next()
			} else {
				response.ErrorTokenAuthFail(context)
			}
		} else {
			response.ErrorTokenBaseInfo(context)
		}
	}
}

// SuperuserAuthentication 超级用户鉴权中间件
func SuperuserAuthentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		phone := context.GetString("phone")
		SuperuserPhone := variable.ConfigYml.GetString("User.Superuser")
		if phone != SuperuserPhone {
			response.Fail(context, consts.CurdSuperuserFailCode, consts.CurdSuperuserFailMsg, "")
		}
		context.Next()
	}
}

// CheckTokenAuthWithRefresh 检查token完整性、有效性并且自动刷新中间件
func CheckTokenAuthWithRefresh() gin.HandlerFunc {
	return func(context *gin.Context) {

		headerParams := HeaderParams{}

		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			response.TokenErrorParam(context, consts.JwtTokenMustValid+err.Error())
			return
		}
		token := strings.Split(headerParams.Authorization, " ")
		if len(token) == 2 && len(token[1]) >= 20 {
			tokenIsEffective := userstoken.CreateUserFactory().IsEffective(token[1])
			// 判断token是否有效
			if tokenIsEffective {
				if customToken, err := userstoken.CreateUserFactory().ParseToken(token[1]); err == nil {
					key := variable.ConfigYml.GetString("Token.BindContextKeyName")
					// token验证通过，同时绑定在请求上下文
					context.Set(key, customToken)
					// 在自动刷新token的中间件中，将请求的认证键、值，原路返回，与后续刷新逻辑格式保持一致
					context.Header("Refresh-Token", "")
					context.Header("Access-Control-Expose-Headers", "Refresh-Token")
				}
				context.Next()
			} else {
				// 判断token是否满足刷新条件
				if userstoken.CreateUserFactory().TokenIsMeetRefreshCondition(token[1]) {
					userIP := context.ClientIP()
					userAddr := ipLocation.GetIPLocation(userIP)
					// 刷新token
					if newToken, ok := userstoken.CreateUserFactory().RefreshToken(token[1], userIP, userAddr); ok {
						if customToken, err := userstoken.CreateUserFactory().ParseToken(newToken); err == nil {
							key := variable.ConfigYml.GetString("Token.BindContextKeyName")
							// token刷新成功，同时绑定在请求上下文
							context.Set(key, customToken)
						}
						// 新token放入header返回
						context.Header("Refresh-Token", newToken)
						context.Header("Access-Control-Expose-Headers", "Refresh-Token")
						context.Next()
					} else {
						response.ErrorTokenRefreshFail(context)
					}
				} else {
					response.ErrorTokenRefreshFail(context)
				}
			}
		} else {
			response.ErrorTokenBaseInfo(context)
		}
	}
}

// RefreshTokenConditionCheck 刷新token条件检查中间件，针对已经过期的token，要求是token格式以及携带的信息满足配置参数即可
func RefreshTokenConditionCheck() gin.HandlerFunc {
	return func(context *gin.Context) {

		headerParams := HeaderParams{}
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			response.TokenErrorParam(context, consts.JwtTokenMustValid+err.Error())
			return
		}
		token := strings.Split(headerParams.Authorization, " ")
		if len(token) == 2 && len(token[1]) >= 20 {
			// 判断token是否满足刷新条件
			if userstoken.CreateUserFactory().TokenIsMeetRefreshCondition(token[1]) {
				context.Next()
			} else {
				response.ErrorTokenRefreshFail(context)
			}
		} else {
			response.ErrorTokenBaseInfo(context)
		}
	}
}
