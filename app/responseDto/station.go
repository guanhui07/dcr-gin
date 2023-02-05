package responseDto

import (
	"dcr-gin/app/model"
	"time"
)

type ResponseStationResp struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// ResponseStation 定义返回的数据模型
type ResponseStation struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	IpAddress   string     `json:"ip_address"`
	TigerShaped string     `json:"tiger_shaped"`
	Heartbeat   string     `json:"heartbeat"`
	AddTime     *time.Time `json:"add_time"`
	UpdateTime  *time.Time `json:"update_time"`
}

// ToStationModelToRes 将数据模型转换为返回值的
func ToStationModelToRes(station model.Station) ResponseStation {
	return ResponseStation{
		Id:        station.Id,
		Name:      station.Name,
		IpAddress: station.IpAddress,
	}
}

// ToStationModelListToRes 列表的转换
func ToStationModelListToRes(stationSlice []model.Station) []ResponseStation {
	result := make([]ResponseStation, 0)
	for _, item := range stationSlice {
		result = append(result, ToStationModelToRes(item))
	}
	return result
}
