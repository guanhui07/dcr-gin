package service

import (
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"dcr-gin/app/requestDto"
	"dcr-gin/app/responseDto"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

func findStation(c *gin.Context, stationParams model.Ticket) (userResp responseDto.ResponseUser, err error) {
	result := global.DB.
		Model(&model.User{}).
		First(&stationParams)
	// 影响行数
	if result.RowsAffected > 0 {
		return userResp, nil
	}
	return userResp, errors.New("用户不存在！")
}

// 位置json
type localJson struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

func CreateStation(c *gin.Context, stationParams requestDto.AddStationReq) error {
	stationModel := model.Station{}
	db := global.DB.Model(stationModel)
	err := db.Where("name=?", stationParams.Name).
		First(&stationModel).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		//找到记录
		return errors.New("此站点已经存在,无法创建")
	}

	location := localJson{}
	location.Lat = stationParams.Lat
	location.Lng = stationParams.Lng
	//序列化 为json str
	locationJson, _ := json.Marshal(location)

	insertData := model.Station{
		Name:        stationParams.Name,
		Location:    locationJson,
		IpAddress:   stationParams.IpAddress,
		TigerShaped: stationParams.TigerShaped,
		AddTime:     time.Now(),
		UpdateTime:  time.Now(),
	}
	// 找不到记录则 新增
	tx := db.Create(&insertData)
	return tx.Error
}

// EditStation 编辑站点
func EditStation(c *gin.Context, stationParams requestDto.EditStationReq) error {
	//,"lng":162.34,"lat":25.35,"ip_address":"192.168.1.222","tiger_shaped":"ElsHZsFObISAMrN7uaiHXwVOAmEjRQrE"
	stationModel := model.Station{}
	// 查找一行
	db := global.DB.Model(stationModel)
	err := db.
		Where("id=?", stationParams.Id).
		First(&stationModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//站点不存在
		return errors.New("此站点不存在,无法修改")
	}
	//{"id":1,"name":"张v415999","lng":162.34,"lat":25.35,"ip_address":"192.168.1.222","tiger_shaped":"ElsHZsFObISAMrN7uaiHXwVOAmEjRQrE"}

	if stationParams.Name != "" {
		stationModel.Name = stationParams.Name
	}
	if stationParams.Lat != 0 && stationParams.Lng != 0 {
		//传了坐标 序列化为json str
		location := localJson{}
		location.Lat = stationParams.Lat
		location.Lng = stationParams.Lng
		locationJson, _ := json.Marshal(location)
		stationModel.Location = locationJson
	}
	// ip
	if stationParams.IpAddress != "" {
		stationModel.IpAddress = stationParams.IpAddress
	}
	// 握手符号
	if stationParams.TigerShaped != "" {
		if len(stationParams.TigerShaped) != 32 {
			//return errors.New("握手符号长度必须32位")
		}
		stationModel.TigerShaped = stationParams.TigerShaped
	}
	stationModel.UpdateTime = time.Now()
	// 更新db
	ret := db.Select("name",
		"location",
		"ip_address", "tiger_shaped",
		"update_time").
		Save(&stationModel)
	return ret.Error
}

type Location struct {
	Lng float32 `json:"lng"`
	Lat float32 `json:"lat"`
}

type station struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	Location    datatypes.JSON `json:"location"`
	IpAddress   string         `json:"ip_address"`
	TigerShaped string         `json:"tiger_shaped"`
	Heartbeat   int64          `json:"heartbeat"`
	Status      int64          `json:"status"`
	AddTime     time.Time      `json:"add_time"`
	Device      string         `json:"device"`
}

type ResponseStation struct {
	CurrentPage int       `json:"current_page"`
	PageRows    int       `json:"page_rows"`
	TotalCount  int64     `json:"total_count"`
	Data        []station `json:"data"`
}

func StationList(c *gin.Context, params requestDto.StationListReq) (err error, configInfo ResponseStation) {
	var stationListCount []model.Station
	var stationListData []model.Station

	var responseData ResponseStation
	page := params.Page
	pageRows := params.PageRows

	var count int64
	db := global.DB.Model(stationListData)
	db.Find(&stationListCount).Count(&count)

	// 查多行 返回切片列表
	err = db.Limit(pageRows).
		Offset((page - 1) * pageRows).
		Find(&stationListData).Error
	if err != nil {
		return errors.New("获取用户列表失败"), responseData
	}

	var Data []station

	copier.Copy(&Data, stationListData)

	responseData = ResponseStation{
		CurrentPage: page,
		PageRows:    pageRows,
		TotalCount:  count,
		Data:        Data,
	}
	return err, responseData
}

// FindStationLastId 获取新站点id
func FindStationLastId(c *gin.Context) (id int64, err error) {
	var stationModel model.Station
	// 最后一行记录 order by id desc
	db := global.DB.Model(stationModel)
	result := db.Model(&model.Station{}).
		Order("id desc").
		First(&stationModel)
	if result.RowsAffected > 0 {
		return stationModel.Id + 1, nil
	}
	return 1, nil

}

// FindStationLastId 获取新站点id
func FindStationById(c *gin.Context, id int64) (station *gorm.DB, err error) {
	var stationModel model.Station
	// 最后一行记录 order by id desc
	db := global.DB.Model(stationModel)
	result := db.Model(&model.Station{}).
		Order("id desc").
		Where("id=?", id).
		First(&stationModel)
	//if result.RowsAffected > 0 {
	//	return stationModel.Id + 1, nil
	//}
	return result, nil

}

// ChangeStationStatus 修改站点状态
func ChangeStationStatus(c *gin.Context, param requestDto.StationStatusReq) error {
	stationModel := model.Station{}
	// 查找一行
	db := global.DB.Model(stationModel)
	err := db.
		Where("id=?", param.Id).
		First(&stationModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("此站点不存在")
	}

	if stationModel.Id == 0 {
		return errors.New("此站点不存在")
	}

	// 0为正常 非0位停用时间戳
	stationModel.Status = param.Status
	if param.Status != 0 {
		//获取unix时间戳
		stationModel.Status = time.Now().Unix()
	}
	stationModel.UpdateTime = time.Now()
	//更新 站点状态
	info := db.Model(&stationModel).
		Select("status", "update_time").
		Updates(stationModel)
	return info.Error
}
