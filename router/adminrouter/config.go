package adminrouter

import (
	system "dcr-gin/app/adminapi/system"
	"dcr-gin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitConfigRouter(Router *gin.RouterGroup) {
	//设置中间件 公用path
	systemRouter := Router.Group("/system", middleware.AuthMiddleWare())
	// /adminapi/system/edit/config
	systemRouter.POST("edit/config", system.EditConfig)
	// /adminapi/system/config/info
	systemRouter.GET("config/info", system.ConfigInfo)
	// /adminapi/system/logout
	systemRouter.POST("logout", system.Logout)
}
