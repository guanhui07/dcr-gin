package ticket

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	utils "dcr-gin/app/utils"
	"github.com/gin-gonic/gin"
)

// Info 信息详情  get  http://localhost:9090/adminapi/ticket/info?id=23
func Info(c *gin.Context) {
	var getByIdDto requestDto.GetById
	if err := c.ShouldBindQuery(&getByIdDto); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	info, err := service.TicketInfo(c, getByIdDto)
	if err == nil {
		utils.Success(c, info, "获取信息成功")
		return
	}
	utils.Fail(c, err.Error())

}
