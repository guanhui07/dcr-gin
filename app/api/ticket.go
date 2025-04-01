package api

import (
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	"dcr-gin/app/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// GetTicketPhotoUuid
//@function: GetTicketPhotoUuid
//@description: 获取信息图片名字
//@param: c *gin.Context
//@return:
// post http://localhost:9090/api/ticket/photoUuid
//{"station_id":"1","ticket_id":"1"}
/**
curl --location --request POST 'http://localhost:9090/api/ticket/photoUuid' \
--header 'Content-Type: application/json' \
--data-raw '{"station_id":"1","ticket_id":"1"}'
*/
func GetTicketPhotoUuid(c *gin.Context) {
	var stationTicketDto requestDto.StationTicketReq
	//从协程上下文获取参数 中间件set
	stationId, _ := c.Get("stationId")
	tigerShaped, _ := c.Get("tigerShaped")
	sign, _ := c.Get("sign")
	device, _ := c.Get("device")
	stationIdSrt := fmt.Sprintf("%v", stationId) // 强制转字符串
	stationTicketDto.StationId = stationIdSrt
	err := c.ShouldBindBodyWith(&stationTicketDto, binding.JSON) // 绑定json 到dto
	if err != nil {                                              // 校验数据是否合法
		msg := utils.ShowErrorMessage(err)
		utils.Fail(c, msg)
		return
	}
	//构建签名切片
	signParam := make(map[string]string)
	signParam["ticket_id"] = stationTicketDto.TicketId
	signParam["device"] = fmt.Sprintf("%v", device)
	//验证签名
	bool := utils.CheckMd5Sign(signParam, fmt.Sprintf("%v", tigerShaped), fmt.Sprintf("%v", sign))
	if !bool {
		//utils.Fail(c, "签名错误")
		//return
	}
	//获取图片名字
	ret, err := service.GetTicketPhotoUuid(c, stationTicketDto)
	if err == nil {
		utils.Success(c, ret, "")
		return
	}
	utils.Fail(c, "缓存文件名错误")

}

// UploadTicketPhoto
//@function: UploadTicketPhoto
//@description: 上传信息图片
//@param: c *gin.Context
//@return:
/**
post http://localhost:9090/api/ticket/uploadPhoto
{
"device":"fsd",
"ticket_id":"12",
}

curl --location --request POST 'http://localhost:9090/api/ticket/uploadPhoto' \
--form 'ticket_id="12"' \
--form 'device="sdfdf"'
*/
func UploadTicketPhoto(c *gin.Context) {
	var ticketUploadReqDto requestDto.TicketUploadReq
	// 从post 表单获取字段
	ticketUploadReqDto.TicketId = c.Request.FormValue("ticket_id")
	ticketUploadReqDto.Device = c.Request.FormValue("device")
	if ticketUploadReqDto.TicketId == "" {
		utils.Fail(c, "信息Id不能为空")
		return
	}
	if ticketUploadReqDto.Device == "" {
		utils.Fail(c, "站点标识不能为空")
		return
	}
	//获取握手符号 从请求获取上下文  一般从中间件设置
	tigerShaped, _ := c.Get("tigerShaped")
	sign, _ := c.Get("sign")
	//构建签名切片
	signParam := make(map[string]string)
	signParam["ticket_id"] = ticketUploadReqDto.TicketId
	// 强转字符串
	signParam["device"] = fmt.Sprintf("%v", ticketUploadReqDto.Device)
	//验证签名
	bool := utils.CheckMd5Sign(signParam, fmt.Sprintf("%v", tigerShaped), fmt.Sprintf("%v", sign))
	if !bool {
		//utils.Fail(c, "签名错误")
		//return
	}
	columnName, err := service.UploadTicketPhoto(c, ticketUploadReqDto)
	if err == nil {
		utils.Success(c, columnName, "图片上传成功")
		return
	}
	utils.Fail(c, err.Error())
	return
}

// UploadTicket
//@function: UploadTicket
//@description: 上传信息数据
//@param: c *gin.Context
//@return:
/**
post http://localhost:9090/api/ticket/upload
{
"device":"aa",
"ticket_id":"12",
}
*/
func UploadTicket(c *gin.Context) {
	var ticketApiReqDto requestDto.TicketApiReq
	err := c.ShouldBindBodyWith(&ticketApiReqDto, binding.JSON)
	if err != nil {
		// 校验数据是否合法
		msg := utils.ShowErrorMessage(err)
		utils.Fail(c, msg)
		return
	}
	//构建签名切片
	signParam := make(map[string]string)
	signParam["device"] = ticketApiReqDto.Device
	signParam["ticket_id"] = ticketApiReqDto.TicketId
	signParam["ticket_sn"] = ticketApiReqDto.TicketSn
	signParam["number_plate"] = ticketApiReqDto.LorryNumber
	signParam["rough_weight"] = ticketApiReqDto.RoughWeight
	signParam["tare_weight"] = ticketApiReqDto.TareWeight
	signParam["net_weight"] = ticketApiReqDto.NetWeight
	signParam["time"] = ticketApiReqDto.Time
	//获取握手符号
	tigerShaped, _ := c.Get("tigerShaped")
	sign, _ := c.Get("sign")
	//验证签名 强转字符串
	bool := utils.CheckMd5Sign(signParam, fmt.Sprintf("%v", tigerShaped), fmt.Sprintf("%v", sign))
	if !bool {
		//utils.Fail(c, "签名错误")
		//return
	}
	stationId, _ := c.Get("stationId")
	stationIdSrt := fmt.Sprintf("%v", stationId)
	// 强转int64
	stationIdInt, _ := strconv.ParseInt(stationIdSrt, 10, 64)
	ticket := model.Ticket{}
	err = global.DB.Where("station_id = ?", stationId).
		Where("client_ticket_id = ?", ticketApiReqDto.TicketId).
		First(&ticket).
		Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 已经找到
		utils.Success(c, gin.H{}, "信息上传成功")
		return
	}
	// 找不到则增加
	ticket.StationId = stationIdInt
	ticket.TicketSn = ticketApiReqDto.TicketSn
	ticket.ClientTicketId = ticketApiReqDto.TicketId
	ticket.LorryNumber = ticketApiReqDto.LorryNumber
	ticket.RoughWeight = ticketApiReqDto.RoughWeight
	ticket.TareWeight = ticketApiReqDto.TareWeight
	ticket.NetWeight = ticketApiReqDto.NetWeight
	ticket.RoughTime = time.Now() //ticketApiReqDto.Time
	ticket.AddTime = time.Now()
	ticket.UpdateTime = time.Now()
	ticket.UploadStatus = 0
	// 增加记录
	tx := global.DB.Create(&ticket)
	if tx.Error == nil {
		utils.Success(c, gin.H{}, "信息上传成功")
		return
	}
	utils.Fail(c, "上传失败！")
}
