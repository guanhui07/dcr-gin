package station

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	"dcr-gin/app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/dump"
)

// 查多行列表
// StationList 站点列表 get http://localhost:9090/adminapi/station/list?page=1&page_rows=10
func StationList(c *gin.Context) {
	var stationDto requestDto.StationList
	var err error
	dump.P(err)
	// 绑定 get
	if err := c.ShouldBindQuery(&stationDto); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	responseErr, listWithPage := service.StationList(c, stationDto)
	for k, item := range listWithPage.Data {
		device, _ := utils.HashidsEncode(fmt.Sprintf("%v", item.Id))
		listWithPage.Data[k].Device = device
	}
	if responseErr == nil {
		utils.Success(c, listWithPage, "站点列表获取成功")
		return
	}
	utils.Fail(c, responseErr.Error())

}
