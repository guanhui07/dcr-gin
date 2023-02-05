package initialize

import (
	middleware2 "dcr-gin/app/middleware"
	"dcr-gin/router/adminrouter"
	"dcr-gin/router/apirouter"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 注册中间件
	Router.Use(
		middleware2.CorsMiddleWare(),    // 跨域的
		middleware2.LoggerMiddleWare(),  // 日志
		middleware2.RecoverMiddleWare(), // 异常的
	)
	// 配置全局公用路径 后台
	AdminApiGroup := Router.Group("adminapi")
	// 注册管理端路由
	adminrouter.InitLoginRouter(AdminApiGroup)
	adminrouter.InitUserRouter(AdminApiGroup)
	adminrouter.InitConfigRouter(AdminApiGroup)
	adminrouter.InitStationRouter(AdminApiGroup)
	adminrouter.InitTicketRouter(AdminApiGroup)

	// 注册api路由 - 站点用
	ApiGroup := Router.Group("api")
	apirouter.InitTicketRouter(ApiGroup)
	return Router
}
