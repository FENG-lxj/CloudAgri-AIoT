package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/middleware/authorization"
	"goskeleton/app/http/middleware/cors"
	validatorFactory "goskeleton/app/http/validator/core/factory"
	"goskeleton/app/service/users/curd"
	"goskeleton/app/utils/gin_release"
	"net/http"
)

// 该路由主要设置 后台管理系统等后端应用路由

func InitWebRouter() *gin.Engine {
	var router *gin.Engine
	// 非调试模式（生产模式） 日志写到日志文件
	if variable.ConfigYml.GetBool("AppDebug") == false {
		router = gin_release.ReleaseRouter()
	} else {
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		pprof.Register(router)
	}
	// 设置可信任的代理服务器列表
	if variable.ConfigYml.GetInt("HttpServer.TrustProxies.IsOpen") == 1 {
		if err := router.SetTrustedProxies(variable.ConfigYml.GetStringSlice("HttpServer.TrustProxies.ProxyServerList")); err != nil {
			variable.ZapLog.Error(consts.GinSetTrustProxyError, zap.Error(err))
		}
	} else {
		_ = router.SetTrustedProxies(nil)
	}
	//根据配置进行设置跨域
	if variable.ConfigYml.GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}
	// 根据配置初始化超级管理员
	curd.CreateUserCurdFactory().SuperUserInit()
	// 创建路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "这里不是你该来的地方了！！！")
	})

	//  创建一个后端接口路由组
	backend := router.Group("/IoTA/")
	{
		//  【不需要token】中间件验证的路由  发送手机验证码、登录
		noAuth := backend.Group("users/")
		{
			// 发送手机验证码
			noAuth.POST("captcha", validatorFactory.Create(consts.ValidatorPrefix+"PhoneCaptcha"))
			// 登录
			noAuth.POST("login", validatorFactory.Create(consts.ValidatorPrefix+"UsersLogin"))
			noAuth.POST("AiImageInfo", validatorFactory.Create(consts.ValidatorPrefix+"AiImageInfo"))
		}
		// 刷新token
		refreshToken := backend.Group("users/")
		{
			// 刷新token，当过期的token在允许失效的延长时间范围内，用旧token换取新token
			refreshToken.Use(authorization.RefreshTokenConditionCheck()).POST("refreshtoken", validatorFactory.Create(consts.ValidatorPrefix+"RefreshToken"))
		}
		//// 【需要token】中间件验证的路由
		//backend.Use(authorization.CheckTokenAuth())
		{
			// 超级管理员路由
			superAdmin := backend.Group("SuperAdmin/")
			// 超级用户鉴权
			superAdmin.Use(authorization.SuperuserAuthentication())
			{
				superAdmin.POST("SuperAdminAddUser", validatorFactory.Create(consts.ValidatorPrefix+"SuperAdminAddUser")) // 添加用户绑定项目
				superAdmin.GET("GetUserList", validatorFactory.Create(consts.ValidatorPrefix+"GetUserList"))
				superAdmin.POST("UpChatGPTKey", validatorFactory.Create(consts.ValidatorPrefix+"UpChatGPTKey"))
				superAdmin.POST("RemoveUser", validatorFactory.Create(consts.ValidatorPrefix+"RemoveUser"))
				superAdmin.POST("UpDev", validatorFactory.Create(consts.ValidatorPrefix+"UpDev"))
				superAdmin.POST("RemoveDev", validatorFactory.Create(consts.ValidatorPrefix+"RemoveDev"))
			}

			// AI组路由
			AI := backend.Group("ai/")
			{
				AI.GET("ProposalAI_AE", validatorFactory.Create(consts.ValidatorPrefix+"ProposalAI_AE"))
				AI.GET("ProposalAI_AS", validatorFactory.Create(consts.ValidatorPrefix+"ProposalAI_AS"))
				AI.GET("AgricultureExpertAI", validatorFactory.Create(consts.ValidatorPrefix+"AgricultureExpertAI")) //农业专家
				AI.GET("DetectionInfo", validatorFactory.Create(consts.ValidatorPrefix+"DetectionInfo"))             //检测信息
				AI.GET("DetectionImage", validatorFactory.Create(consts.ValidatorPrefix+"DetectionImage"))           //检测图片
				AI.GET("GetMonitorLogs", validatorFactory.Create(consts.ValidatorPrefix+"GetMonitorLogs"))           //AI监控日志
				AI.GET("PestProposal", validatorFactory.Create(consts.ValidatorPrefix+"PestProposal"))               //病虫害防治建议
			}
			BigData := backend.Group("bigdata/")
			{
				BigData.GET("GetAEdata", validatorFactory.Create(consts.ValidatorPrefix+"GetAEdata"))
				BigData.GET("GetASdata", validatorFactory.Create(consts.ValidatorPrefix+"GetASdata"))
				BigData.GET("GetDevInfo", validatorFactory.Create(consts.ValidatorPrefix+"GetDevInfo"))
				BigData.GET("GetMemCPU", validatorFactory.Create(consts.ValidatorPrefix+"GetMemCPU"))
			}
			//文件上传公共路由
			uploadFiles := backend.Group("upload/")
			{
				uploadFiles.POST("files", validatorFactory.Create(consts.ValidatorPrefix+"UploadFiles"))
				//uploadFiles.POST("AiImageInfo", validatorFactory.Create(consts.ValidatorPrefix+"AiImageInfo"))
			}
		}
	}
	return router
}
