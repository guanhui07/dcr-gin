package service

import (
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"dcr-gin/app/requestDto"
	"dcr-gin/app/responseDto"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

// TicketList 列表查询返回
func TicketList(c *gin.Context, params requestDto.TicketListReq) (pageResult responseDto.ResponsePageResult, err error) {
	var ticketModelList []model.Ticket
	page := params.Page
	pageRows := params.PageRows

	db := global.DB.Model(&ticketModelList).Preload("ResponseStation")
	fmt.Println(params.StationId)
	// 拼接where sql
	if params.StationId > 0 {
		db.Where("station_id=?", params.StationId)
	}
	if params.TicketSn != "" {
		db.Where("ticket_sn LIKE ?", "%"+params.TicketSn+"%")
	}
	if params.LorryNumber != "" {
		db.Where("lorry_number LIKE ?", "%"+params.LorryNumber+"%")
	}
	if params.UploadStatus > -1 {
		// 强转int64
		db.Where("upload_status = ?", int64(params.UploadStatus))
	}
	if params.NetWeightS > 0 {
		db.Where("net_weight >= ?", params.NetWeightS)
	}
	if params.NetWeightE > 0 {
		db.Where("net_weight <= ?", params.NetWeightE)
	}
	if params.RoughTimeS != "" {
		db.Where("rough_time >= ?", params.RoughTimeS)
	}
	if params.RoughTimeE != "" {
		db.Where("rough_time <= ?", params.RoughTimeE)
	}
	if params.RoughWeightS > 0 {
		db.Where("rough_weight >= ?", params.RoughWeightS)
	}
	if params.RoughWeightE > 0 {
		db.Where("rough_weight <= ?", params.RoughWeightE)
	}
	var count int64
	//查询总条目
	db.Find(&ticketModelList).Count(&count)
	// 查询列表
	if err := db.Limit(pageRows).
		Offset((page - 1) * pageRows).
		Order("id desc").
		Find(&ticketModelList).Error; err != nil {
		return pageResult, errors.New("获取信息列表失败")
	}
	localUrl := global.ServerConfig.Local.Url
	// 遍历
	for k, item := range ticketModelList {
		if item.PhotoGoods != "" {
			ticketModelList[k].PhotoGoods = localUrl + item.PhotoGoods
		}
		if item.PhotoFront != "" {
			ticketModelList[k].PhotoFront = localUrl + item.PhotoFront
		}
		if item.PhotoLorryNumber != "" {
			ticketModelList[k].PhotoLorryNumber = localUrl + item.PhotoLorryNumber
		}
		if item.PhotoBehind != "" {
			ticketModelList[k].PhotoBehind = localUrl + item.PhotoBehind
		}
	}
	// 返回列表Dto
	return responseDto.ResponsePageResult{
		Page:     page,
		PageRows: pageRows,
		Total:    count,
		List:     ticketModelList,
	}, err

}
