package system

import (
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"dcr-gin/app/service"
	"dcr-gin/app/utils"
	"fmt"
	"github.com/gookit/goutil/dump"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2" //carbon
	//"github.com/onsi/gomega/matchers"
	"time"
)

// TodayStatisticCount 今日信息上报概览
type TodayStatisticCount struct {
	UploadStatus int64  `json:"upload_status"`
	Total        int64  `json:"total"`
	Ratio        string `json:"ratio"`
}

// TodayStationCount 今日站点统计
type TodayStationCount struct {
	Name  string `json:"name"`
	Total int64  `json:"total"`
}

// ReturnValue 首页返回值
type ReturnValue struct {
	TodayStatusStatistics []TodayStatisticCount `json:"today_status_statistics"`
	TodayTicketCount      []TodayStationCount   `json:"today_ticket_count"`
	RightPic              interface{}           `json:"right_pic"`
}

// TicketSuccessStatus 成功状态信息
type TicketSuccessStatus struct {
	Month int   `json:"month"`
	Total int64 `json:"total"`
}

// TicketFailStatus 失败状态信息
type TicketFailStatus struct {
	Month int   `json:"month"`
	Total int64 `json:"total"`
}

// TicketTotal 月总计信息数组
type TicketTotal struct {
	Month int   `json:"month"`
	Total int64 `json:"total"`
}

type TicketMonthStatistic struct {
	SuccessCount int64  `json:"success_count"` // 成功数
	FailCount    int64  `json:"fail_count"`    // 失败数
	SucceeRate   string `json:"succee_rate"`   // 成功率
	TotalCount   int64  `json:"total_count"`   //  总数
}

// StatisticsTicket a
// get http://localhost:9090/adminapi/ticket/home/statistic
func StatisticsTicket(ctx *gin.Context) {
	TicketModel := model.Ticket{}

	var TodayCountDataSlice []TodayStatisticCount    // 总计
	var StationCountDataSlice []TodayStationCount    // 站点信息统计
	var SuccessStatusDataSlice []TicketSuccessStatus //月信息成功统计
	var FailStatusDataSlice []TicketFailStatus       //月信息失败统
	var MonthTicketTotalDataSlice []TicketTotal      // 月总信息数
	SuccessStatusMap := make(map[int]int64)          // 月成功数
	FailStatusMap := make(map[int]int64)             // 月成失败数
	MonthStatisticData := TicketMonthStatistic{}
	MonthStatisticListMap := make(map[int]TicketMonthStatistic)

	startTime := time.Now().Format("2006-01-02 15:04:05")
	year := time.Now().Year()
	startYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local).
		Format("2006-01-02 15:04:05")
	_ = carbon.Now().String()
	_ = carbon.Now().AddDays(6).String()
	startTime = carbon.Now().SubMonths(60).String()
	dump.P(startTime)
	var value ReturnValue
	var TotalWeightCount int64

	dbTicketModel := global.DB.Model(TicketModel)
	TotalWeightCount = getTotalWeightCount(dbTicketModel, startTime, TotalWeightCount)

	_, TodayCountDataSlice = getTodayCountDataSlice(TicketModel, startTime, TodayCountDataSlice)

	_, StationCountDataSlice = getStationCountDataSlice(TicketModel, startTime, StationCountDataSlice)
	// 每月上传失败
	_, FailStatusDataSlice = getFailStatusDataSlice(dbTicketModel, startYear, FailStatusDataSlice)
	// 每月上传成功
	_, SuccessStatusDataSlice = getSuccessStatusDataSlice(dbTicketModel, startYear, SuccessStatusDataSlice)
	// 每月总数
	_, MonthTicketTotalDataSlice = getMonthTicketTotalDataSlice(dbTicketModel, startYear, MonthTicketTotalDataSlice)

	TodayCountDataSlice = dealTotalWeightCount(TodayCountDataSlice, TotalWeightCount)

	// 成功
	if len(SuccessStatusDataSlice) > 0 {
		for _, selenet := range SuccessStatusDataSlice {
			SuccessStatusMap[selenet.Month] = selenet.Total
		}
	}
	// 失败
	if len(FailStatusDataSlice) > 0 {
		for _, fItem := range FailStatusDataSlice {
			FailStatusMap[fItem.Month] = fItem.Total
		}
	}

	if len(MonthTicketTotalDataSlice) > 0 {
		var successRate string = "0"
		for _, eItem := range MonthTicketTotalDataSlice {
			successValue := SuccessStatusMap[eItem.Month]
			totalValue := eItem.Total
			if totalValue > 0 {
				value := float64(successValue) / float64(totalValue)
				successRate = fmt.Sprintf("%.2f", value)
			}
			MonthStatisticData.SucceeRate = successRate
			MonthStatisticData.FailCount = FailStatusMap[eItem.Month]
			MonthStatisticData.SuccessCount = successValue
			MonthStatisticData.TotalCount = eItem.Total
			MonthStatisticListMap[eItem.Month] = MonthStatisticData
		}
	}

	value.TodayStatusStatistics = TodayCountDataSlice
	value.TodayTicketCount = StationCountDataSlice
	value.RightPic = MonthStatisticListMap

	utils.Success(ctx, value, "获取配置成功")
}

func getTotalWeightCount(dbTicketModel *gorm.DB, startTime string, TotalWeightCount int64) int64 {
	dbTicketModel.
		Where("add_time>?", startTime).
		Count(&TotalWeightCount)
	return TotalWeightCount
}

func dealTotalWeightCount(TodayCountDataSlice []TodayStatisticCount, TotalWeightCount int64) []TodayStatisticCount {
	if len(TodayCountDataSlice) > 0 {
		for key, element := range TodayCountDataSlice {
			if TotalWeightCount > 0 {
				value := float64(element.Total) / float64(TotalWeightCount)
				successRate := fmt.Sprintf("%.2f", value)
				TodayCountDataSlice[key].Ratio = successRate
			}
		}
	}
	return TodayCountDataSlice
}

func getSuccessStatusDataSlice(dbTicketModel *gorm.DB, startYear string, SuccessStatusDataSlice []TicketSuccessStatus) (*gorm.DB, []TicketSuccessStatus) {
	SuccessCount := dbTicketModel.
		Where("add_time>?", startYear).
		Where("upload_status=?", 2).
		Select("date_format(add_time,'%m') as month ", "count(id) as total").
		Group("month").Find(&SuccessStatusDataSlice)
	return SuccessCount, SuccessStatusDataSlice
}

func getFailStatusDataSlice(dbTicketModel *gorm.DB, startYear string, FailStatusDataSlice []TicketFailStatus) (*gorm.DB, []TicketFailStatus) {
	FailCount := dbTicketModel.
		Where("add_time>?", startYear).
		Where("upload_status=?", 1).
		Select("date_format(add_time,'%m') as month ", "count(id) as total").
		Group("month").Find(&FailStatusDataSlice)
	return FailCount, FailStatusDataSlice
}

func getMonthTicketTotalDataSlice(dbTicketModel *gorm.DB, startYear string, MonthTicketTotalDataSlice []TicketTotal) (*gorm.DB, []TicketTotal) {
	TotalCount := dbTicketModel.
		Where("add_time>?", startYear).
		Select("date_format(add_time,'%m') as month ", "count(id) as total").
		Group("month").
		Find(&MonthTicketTotalDataSlice)
	return TotalCount, MonthTicketTotalDataSlice
}

func getStationCountDataSlice(TicketModel model.Ticket, startTime string, StationCountDataSlice []TodayStationCount) (*gorm.DB, []TodayStationCount) {
	return global.DB.Model(TicketModel).
		Where("pf_ticket_info.add_time>?", startTime).
		Select("count(pf_ticket_info.id) as total", "pf_station_info.name").
		Joins("left join pf_station_info on pf_station_info.id = pf_ticket_info.station_id").
		Group("pf_station_info.name").
		Find(&StationCountDataSlice), StationCountDataSlice
}

func getTodayCountDataSlice(TicketModel model.Ticket, startTime string, TodayCountDataSlice []TodayStatisticCount) (*gorm.DB, []TodayStatisticCount) {
	totalCountvalue := global.DB.Model(TicketModel).
		Where("add_time>?", startTime).
		Select("count(id) as total", "upload_status").
		Group("upload_status").
		Find(&TodayCountDataSlice) //分类今日统计
	return totalCountvalue, TodayCountDataSlice
}

// StatisticsList 首页右下表图
func StatisticsList(c *gin.Context) {
	responseErr, list := service.RightTicketList(c)
	if responseErr == nil {
		utils.Success(c, list, "获取重试信息成功")
		return
	}
	utils.Fail(c, responseErr.Error())

}
