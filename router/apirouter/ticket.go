package apirouter

import (
	"dcr-gin/app/api"
	"dcr-gin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitTicketRouter(Router *gin.RouterGroup) {
	//设置中间件 公用path
	TicketRouter := Router.Group("ticket", middleware.ApiMiddleWare())
	// /api/ticket/photoUuid
	TicketRouter.POST("photoUuid", api.GetTicketPhotoUuid)
	// /api/ticket/uploadPhoto
	TicketRouter.POST("uploadPhoto", api.UploadTicketPhoto)
	// /api/ticket/upload
	TicketRouter.POST("upload", api.UploadTicket)
}
