package station

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	"dcr-gin/app/utils"
	"github.com/gin-gonic/gin"
)

// ChangeStationStatus  查站点一行 修改状态
// post  http://localhost:9090/adminapi/station/changeStatus
// {"id":2,"status":1}
// curl -X POST http://localhost:9090/adminapi/station/changeStatus -d '{"id":1,"status":1}'
func ChangeStationStatus(c *gin.Context) {
	var stationDto requestDto.StationStatus
	//return errors.New("此站点不存在")
	//var err error
	//dump.P(err)
	err := c.ShouldBindJSON(&stationDto)
	if err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}

	err = service.ChangeStationStatus(c, stationDto)
	if err == nil {
		utils.Success(c, "", "修改站点状态成功")
		return
	}
	utils.Fail(c, err.Error())
}
