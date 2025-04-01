package ticket

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	utils "dcr-gin/app/utils"
	"github.com/gin-gonic/gin"
)

// Info 信息详情  get  http://localhost:9090/adminapi/ticket/info?id=23
func Info(c *gin.Context) {
	var getByIdDto requestDto.GetByIdReq
	err := c.ShouldBindQuery(&getByIdDto)
	if err != nil {
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

// List 信息列表 get http://localhost:9090/adminapi/user/list?page=1&page_rows=10&station_id=23
func List(c *gin.Context) {
	var ticketListReqDto requestDto.TicketListReq
	err := c.ShouldBindQuery(&ticketListReqDto)
	if err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	list, err := service.TicketList(c, ticketListReqDto)
	if err == nil {
		utils.Success(c, list, "获取信息列表成功")
		return
	}
	utils.Fail(c, err.Error())

}
