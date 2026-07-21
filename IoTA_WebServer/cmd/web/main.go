package main

import (
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/logo"
	_ "goskeleton/bootstrap"
	"goskeleton/routers"
)

// 这里可以存放后端路由（例如后台管理系统）
func main() {
	logo.Banner()
	router := routers.InitWebRouter()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Web.Port"))
}
