package service

import (
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"dcr-gin/app/requestDto"
	"dcr-gin/app/responseDto"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func TicketInfo(c *gin.Context, getByIdDto requestDto.GetByIdReq) (ticketInfoRes responseDto.TicketInfoRes, err error) {
	ticketModel := model.Ticket{}
	//查询一行记录
	info := global.DB.Where("id=?", getByIdDto.Id).First(&ticketModel)
	localUrl := global.ServerConfig.Local.Url
	if ticketModel.PhotoGoods != "" {
		ticketModel.PhotoGoods = localUrl + ticketModel.PhotoGoods
	}
	if ticketModel.PhotoFront != "" {
		ticketModel.PhotoFront = localUrl + ticketModel.PhotoFront
	}
	if ticketModel.PhotoLorryNumber != "" {
		ticketModel.PhotoLorryNumber = localUrl + ticketModel.PhotoLorryNumber
	}
	if ticketModel.PhotoBehind != "" {
		ticketModel.PhotoBehind = localUrl + ticketModel.PhotoBehind
	}
	copier.Copy(&ticketInfoRes, ticketModel)
	if info.RowsAffected > 0 {
		return ticketInfoRes, nil
	}
	return ticketInfoRes, errors.New("信息不存在")

}
