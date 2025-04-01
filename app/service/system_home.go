package service

import (
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"time"
)

type RetryTicketList struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Date      string    `json:"date"`
	AddTime   time.Time `json:"-"`
	StationId int64     `json:"station_id"`
	TicketSn  string    `json:"ticket_sn"`
}

func RightTicketList(c *gin.Context) (err error, list []RetryTicketList) {
	ticketListData := model.Ticket{}
	var RetryTicketListValue []RetryTicketList
	err = global.DB.Model(ticketListData).Where("pf_ticket_info.upload_status=?", 1).
		Select("pf_ticket_info.id",
			"date_format(pf_ticket_info.add_time,'%Y-%m-%d %H:%i:%s') as date",
			"pf_ticket_info.station_id",
			"pf_ticket_info.ticket_sn",
			"pf_station_info.name").
		Joins("left join pf_station_info on pf_station_info.id = pf_ticket_info.station_id").
		Find(&RetryTicketListValue).Error
	if err != nil {
		return errors.New("获取重试信息失败"), []RetryTicketList{}
	}

	var retryList []RetryTicketList
	copier.Copy(&retryList, RetryTicketListValue)
	return err, retryList
}
