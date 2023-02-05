package station

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	utils "dcr-gin/app/utils"
	"github.com/gin-gonic/gin"
)

// RegisterStation 注册站点 post json
/**
post http://localhost:9090/adminapi/station/create
{
"name":"fsd",
"ip_address":"sdf",
"tiger_shaped":"sdf",
"lng":"sdf",
"lat":"sdf",
}
*/
func RegisterStation(c *gin.Context) {
	var stationDto requestDto.AddStation
	if err := c.ShouldBindJSON(&stationDto); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	err := service.CreateStation(c, stationDto)
	if err == nil {
		utils.Success(c, "", "注册站点成功")
		return
	}
	utils.Fail(c, err.Error())
}
