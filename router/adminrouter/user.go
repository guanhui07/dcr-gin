package adminrouter

import (
	user "dcr-gin/app/adminapi/user"
	"dcr-gin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	//设置中间件 公用path
	UserRouter := Router.Group("/user", middleware.AuthMiddleWare())
	// /adminapi/user/add
	UserRouter.POST("add", user.AddUser)
	// /adminapi/user/update
	UserRouter.POST("update", user.UpdateUser)
	// /adminapi/user/change/status
	UserRouter.POST("change/status", user.ChangeStatus)
	// /adminapi/user/change/passwd
	UserRouter.POST("change/passwd", user.ChangePw)
	// /adminapi/user/change/admin/passwd
	UserRouter.POST("change/admin/passwd", user.ChangAdminPw)
	// /adminapi/user/list
	UserRouter.GET("list", user.UserList)
}
