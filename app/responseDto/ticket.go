package responseDto

import (
	"dcr-gin/app/model"
	"time"
)

// Ticket 定义返回的数据模型
type Ticket struct {
	Id         int64      `json:"id"`
	TicketSn   string     `json:"ticket_sn"`
	StationId  int64      `json:"station_id"`
	AddTime    *time.Time `json:"add_time"`
	UpdateTime *time.Time `json:"update_time"`
}

// ToTicketModelToRes 将数据模型转换为返回值的
func ToTicketModelToRes(ticket model.Ticket) Ticket {
	return Ticket{
		Id:        ticket.Id,
		TicketSn:  ticket.TicketSn,
		StationId: ticket.StationId,
	}
}

// ToTicketModelListToRes 列表的转换
func ToTicketModelListToRes(ticketSlice []model.Ticket) []Ticket {
	result := make([]Ticket, 0)
	for _, item := range ticketSlice {
		result = append(result, ToTicketModelToRes(item))
	}
	return result
}

type FileUploadAndDownload struct {
	Name string `json:"name" gorm:"comment:文件名"` // 文件名
	Url  string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag  string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key  string `json:"key" gorm:"comment:编号"`   // 编号
}
type TicketListRes struct {
	Id               int64           `json:"id"`
	TicketSn         string          `json:"ticket_sn"`
	StationId        int64           `json:"station_id"`
	ClientTicketId   string          `json:"client_ticket_id"`
	LorryNumber      string          `json:"lorry_number"`
	RoughWeight      string          `json:"rough_weight"`
	TareWeight       string          `json:"tare_weight"`
	NetWeight        string          `json:"net_weight"`
	RoughTime        time.Time       `json:"rough_time"`
	UploadStatus     int64           `json:"upload_status"`
	UploadRetry      int64           `json:"upload_retry"`
	PhotoFront       string          `json:"photo_front"`
	PhotoBehind      string          `json:"photo_behind"`
	PhotoLorryNumber string          `json:"photo_lorry_number"`
	PhotoGoods       string          `json:"photo_goods"`
	UploadTime       time.Time       `json:"upload_time,omitempty"`
	AddTime          *time.Time      `json:"add_time"`
	UpdateTime       *time.Time      `json:"update_time"`
	Station          ResponseStation `json:"station"`
}

type TicketInfoRes struct {
	Id               int64      `json:"id"`
	StationId        int64      `json:"station_id"`
	TicketSn         string     `json:"ticket_sn"`
	LorryNumber      string     `json:"lorry_number"`
	RoughWeight      string     `json:"rough_weight"`
	TareWeight       string     `json:"tare_weight"`
	NetWeight        string     `json:"net_weight"`
	RoughTime        time.Time  `json:"rough_time"`
	UploadStatus     int64      `json:"upload_status"`
	UploadRetry      int64      `json:"upload_retry"`
	PhotoFront       string     `json:"photo_front"`
	PhotoBehind      string     `json:"photo_behind"`
	PhotoLorryNumber string     `json:"photo_lorry_number"`
	PhotoGoods       string     `json:"photo_goods"`
	AddTime          *time.Time `json:"add_time"`
	UpdateTime       *time.Time `json:"update_time"`
}
