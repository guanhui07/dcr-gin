package station

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	utils "dcr-gin/app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// EditStation 编辑站点 post json  查一行 编辑站点
/**
post  http://localhost:9090/adminapi/station/edit
{
"name":"fsd",
"ip_address":"sdf",
"tiger_shaped":"sdf",
"lng":"sdf",
"lat":"sdf",
"id":"1"
}
*/
func EditStation(c *gin.Context) {
	var stationDto requestDto.EditStation
	//if err := c.ShouldBindJSON(&stationDto); err != nil {
	//	message := utils.ShowErrorMessage(err)
	//	utils.Fail(c, message)
	//	return
	//}
	err := c.ShouldBindJSON(&stationDto)
	if err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	err = service.EditStation(c, stationDto)
	if err == nil {
		utils.Success(c, "", "站点编辑成功")
		return
	}
	utils.Fail(c, err.Error())
}

// EncodeStationId 获取站点Id
// http://localhost:9090/adminapi/station/getLastId
// get /adminapi/station/getLastId
func EncodeStationId(c *gin.Context) {
	id, _ := service.FindStationLastId(c)
	encodeStationId, _ := utils.HashidsEncode(fmt.Sprintf("%v", id))
	if encodeStationId != "" {
		utils.Success(c, encodeStationId, "")
		return
	}
	utils.Fail(c, "获取失败")
}
