package job

import (
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"dcr-gin/app/utils"
	"time"
)

type Rsp struct {
	Code int    `json:"code"`
	Rsp  string `json:"rsp"`
}

// Run 上传信息数据
func Run(stopChannel <-chan int) error {
	configModel := model.Config{}
	global.DB.First(&configModel)
	retryTime := configModel.UploadFailRetryTime

	if configModel.UploadAutoRetry == 0 {
		tick := time.NewTicker(time.Duration(retryTime) * time.Second)
		//fmt.Println(time.Duration(retryTime)*time.Second)
		for {
			select {
			case <-stopChannel:
				return nil
			case <-tick.C:
				//global.DB.First(&configModel)
				//retryTime = configModel.UploadFailRetryTime
				//tick = time.NewTicker(time.Duration(retryTime) * time.Second)
				ticketUpload()
			}
		}
	}
	return nil

}

type SendTicket struct {
	Id               int64     `json:"id"`
	StationId        int64     `json:"station_id"`
	StationName      string    `json:"station_name"`
	TicketSn         string    `json:"ticket_sn"`
	LorryNumber      string    `json:"lorry_number"`
	RoughWeight      string    `json:"rough_weight"`
	TareWeight       string    `json:"tare_weight"`
	NetWeight        string    `json:"net_weight"`
	RoughTime        string    `json:"rough_time"`
	PhotoFront       string    `json:"photo_front"`
	PhotoBehind      string    `json:"photo_behind"`
	PhotoLorryNumber string    `json:"photo_lorry_number"`
	PhotoGoods       string    `json:"photo_goods"`
	AddTime          string    `json:"add_time"`
	ClientTicketId   string    `json:"client_ticket_id"`
	UploadStatus     int64     `json:"upload_status"`
	UploadRetry      int64     `json:"upload_retry"`
	UploadTime       time.Time `json:"upload_time"`
}

type Res struct {
	Code int    `json:"code"`
	Rsp  string `json:"rsp"`
}

func ticketUpload() {
	url := global.ServerConfig.UploadTicket.Url
	ticketList := model.Ticket{}
	var WeightTicketDataSlice []SendTicket

	// join 查询
	global.DB.Model(ticketList).
		Where("pf_ticket_info.upload_status in ?", []int{0, 1}).
		Select("pf_ticket_info.id", "pf_ticket_info.station_id",
			"pf_station_info.name as station_name", "pf_ticket_info.ticket_sn",
			"pf_ticket_info.lorry_number", "pf_ticket_info.rough_weight",
			"pf_ticket_info.tare_weight", "pf_ticket_info.net_weight",
			"date_format(pf_ticket_info.rough_time,'%Y-%m-%d %H:%i:%s') as rough_time",
			"pf_ticket_info.photo_front", "pf_ticket_info.photo_behind",
			"pf_ticket_info.photo_lorry_number", "pf_ticket_info.photo_goods",
			"date_format(pf_ticket_info.add_time,'%Y-%m-%d %H:%i:%s') as add_time",
			"pf_ticket_info.client_ticket_id", "pf_ticket_info.upload_status",
			"pf_ticket_info.upload_retry").
		Joins("left join pf_station_info on pf_station_info.id = pf_ticket_info.station_id").
		Find(&WeightTicketDataSlice)

	for _, item := range WeightTicketDataSlice {
		// curl post上报数据
		res := utils.CurlPost(url, item)
		//v:= reflect.TypeOf(res["code"])
		//global.Logger.Info(fmt.Sprintf("code类型：%+v",v))
		//global.Logger.Info(fmt.Sprintf("返回的res数据是:%+v",res))
		//global.Logger.Info(fmt.Sprintf("返回的res数据是:%d",res["code"]))
		//global.Logger.Info(fmt.Sprintf("返回的res数据是:%+v",res["success"]))

		var SuccessCode int64 = 200
		if res == nil {
			continue
		}

		// 返回 res body体 code
		if res["code"] != float64(SuccessCode) {
			item.UploadRetry = item.UploadRetry + 1
			item.UploadStatus = 1
			global.DB.Model(ticketList).
				Where("id=?", item.Id).
				Select("upload_status", "upload_retry").
				Updates(&item)
			continue
		}
		item.UploadTime = time.Now()
		item.UploadStatus = 2
		// 更新数据
		global.DB.Model(ticketList).
			Where("id=?", item.Id).
			Select("upload_status", "upload_time").
			Updates(&item)
	}
}
