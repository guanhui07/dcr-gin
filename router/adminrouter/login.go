package adminrouter

import (
	"dcr-gin/app/adminapi/system"
	"github.com/gin-gonic/gin"
)

func InitLoginRouter(Router *gin.RouterGroup) {
	// 公用path
	systemRouter := Router.Group("/system")
	// /adminapi/system/login
	systemRouter.POST("/login", system.Login)
}
