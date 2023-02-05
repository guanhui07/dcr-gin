package adminrouter

import (
	"dcr-gin/app/adminapi/system"
	ticket "dcr-gin/app/adminapi/ticket"
	"dcr-gin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitTicketRouter(Router *gin.RouterGroup) {
	//设置中间件 公用path
	TicketRouter := Router.Group("ticket", middleware.AuthMiddleWare())
	// /adminapi/ticket/info
	TicketRouter.GET("info", ticket.Info) // 注册
	// /adminapi/ticket/list
	TicketRouter.GET("list", ticket.List) // 注册
	// /adminapi/ticket/home/statistic
	TicketRouter.GET("home/statistic", system.StatisticsTicket) // 首页左三数据
	// /adminapi/ticket/home/list
	TicketRouter.GET("home/list", system.StatisticsList) // 首页右下表图
}
