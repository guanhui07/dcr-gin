package station

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	"dcr-gin/app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/dump"
	"strconv"
)

// 查多行列表
// StationList 站点列表 get http://localhost:9090/adminapi/station/list?page=1&page_rows=10
func StationList(c *gin.Context) {
	var stationDto requestDto.StationListReq
	var err error
	dump.P(err)
	// 绑定 get
	err = c.ShouldBindQuery(&stationDto)
	if err != nil {
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

// 查一行
// StationList 站点列表 get http://localhost:9090/adminapi/station/detail?id=1
func StationDetail(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	ret, err := service.FindStationById(c, id)
	//for k, item := range listWithPage.Data {
	//	device, _ := utils.HashidsEncode(fmt.Sprintf("%v", item.Id))
	//	listWithPage.Data[k].Device = device
	//}
	if err == nil {
		utils.Success(c, ret, "站点列表获取成功")
		return
	}
	utils.Fail(c, err.Error())

}

// AddStation 注册站点 post json
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
func AddStation(c *gin.Context) {
	var stationDto requestDto.AddStationReq
	// 解析到结构体
	err := c.ShouldBindJSON(&stationDto)
	if err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	err = service.CreateStation(c, stationDto)
	if err == nil {
		utils.Success(c, "", "注册站点成功")
		return
	}
	utils.Fail(c, err.Error())
}

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
	var stationDto requestDto.EditStationReq
	//if err := c.ShouldBindJSON(&stationDto); err != nil {
	//	message := utils.ShowErrorMessage(err)
	//	utils.Fail(c, message)
	//	return
	//}
	//	// 绑定到结构体
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
	idStr := fmt.Sprintf("%v", id)
	encodeStationId, _ := utils.HashidsEncode(idStr)
	if encodeStationId != "" {
		utils.Success(c, encodeStationId, "")
		return
	}
	utils.Fail(c, "获取失败")
}

// ChangeStationStatus  查站点一行 修改状态
// post  http://localhost:9090/adminapi/station/changeStatus
// {"id":2,"status":1}
// curl -X POST http://localhost:9090/adminapi/station/changeStatus -d '{"id":1,"status":1}'
func ChangeStationStatus(c *gin.Context) {
	var stationDtoReq requestDto.StationStatusReq
	//return errors.New("此站点不存在")
	//var err error
	//dump.P(err)

	// 绑定到结构体
	err := c.ShouldBindJSON(&stationDtoReq)
	if err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}

	err = service.ChangeStationStatus(c, stationDtoReq)
	if err == nil {
		utils.Success(c, "", "修改站点状态成功")
		return
	}
	utils.Fail(c, err.Error())
}
