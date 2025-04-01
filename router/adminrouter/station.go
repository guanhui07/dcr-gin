package adminrouter

import (
	station "dcr-gin/app/adminapi/station"
	"dcr-gin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitStationRouter(Router *gin.RouterGroup) {
	//设置中间件 公用path
	StationRouter := Router.Group("/station", middleware.AuthMiddleWare())
	// /adminapi/station/create
	StationRouter.POST("create", station.AddStation) // 注册
	// /adminapi/station/edit
	StationRouter.POST("edit", station.EditStation) // 编辑站点
	// /adminapi/station/list
	StationRouter.GET("list", station.StationList) // 列表
	// /adminapi/station/changeStatus
	StationRouter.POST("changeStatus", station.ChangeStationStatus) // 修改状态
	// /adminapi/station/getLastId
	StationRouter.GET("getLastId", station.EncodeStationId) // 获取新的站点标识
}
