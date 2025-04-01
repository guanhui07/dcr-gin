package service

import (
	"context"
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"dcr-gin/app/requestDto"
	"dcr-gin/app/responseDto"
	"dcr-gin/app/utils/upload"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/wxnacy/wgo/arrays"
	"go.uber.org/zap"
	"strings"
	"time"
)

// FindTicket
// @function: FindTicket
// @description: 查找信息
// @param: c *gin.Context, ticketParams model.Ticket
// @return: userResp responseDto.ResponseUser, err error
func FindTicket(c *gin.Context, ticketParams model.Ticket) (userResp responseDto.ResponseUser, err error) {
	result := global.DB.Model(&model.User{}).First(&ticketParams)
	// 记录存在
	if result.RowsAffected > 0 {
		return userResp, nil
	}
	return userResp, errors.New("用户不存在！")

}

// UpdateTicket
// @function: UpdateTicket
// @description: 修改信息图片路径
// @param: u *model.SysUser, newPassword string
// @return: userInter *model.SysUser,err error
func UpdateTicket(Id string, columnName string, val string) (ticketResp *responseDto.Ticket, err error) {
	ticket := model.Ticket{}
	err = global.DB.
		Where("client_ticket_id = ?", Id).
		First(&ticket).Error
	if err != nil {
		//记录不存在
		return nil, err
	}

	switch columnName {
	case "photo_front":
		ticket.PhotoFront = val
	case "photo_behind":
		ticket.PhotoBehind = val
	case "photo_lorry_number":
		ticket.PhotoLorryNumber = val
	case "photo_goods":
		ticket.PhotoGoods = val
	default:
		return nil, errors.New("上传的照片不正确，没有该字段！")
	}

	ticket.UpdateTime = time.Now()
	// 更新图片url
	err = global.DB.Save(&ticket).Error

	return nil, err
}

// GetTicketPhotoUuid
// @function: GetTicketPhotoUuid
// @description: 获取图片名称
// @param: c *gin.Context, stationTicketId requestDto.StationTicketId
// @return: map[string]string, error
func GetTicketPhotoUuid(c *gin.Context, stationTicket requestDto.StationTicketReq) (map[string]string, error) {
	//优先从缓存中获取
	cacheKey := fmt.Sprintf("cache:ticketPhoto:stationId:%s:ticketId:%s",
		stationTicket.StationId, stationTicket.TicketId)
	//获取缓存中的信息图片名字
	picJson, err := global.Redis.Get(context.Background(), cacheKey).Result()
	map1 := make(map[string]string)
	// json_decode
	_ = json.Unmarshal([]byte(picJson), &map1)
	if err == nil {
		return map1, nil
	}
	global.Logger.Info("缓存中不存在信息图片名字,将构建缓存"+cacheKey, zap.Error(err))

	//生成信息图片名字
	photoFront := uuid.Must(uuid.NewV4()).String()
	photoBehind := uuid.Must(uuid.NewV4()).String()
	photoLorryNumber := uuid.Must(uuid.NewV4()).String()
	photoGoods := uuid.Must(uuid.NewV4()).String()
	//构建返回数据
	retMap := map[string]string{
		"photo_front":        photoFront,
		"photo_behind":       photoBehind,
		"photo_lorry_number": photoLorryNumber,
		"photo_goods":        photoGoods,
	}

	timer := time.Duration(global.ServerConfig.Local.ExpiresTime) * time.Second
	// json_encode
	str, err := json.Marshal(&retMap)
	//存储缓存中后返回
	err = global.Redis.Set(context.Background(), cacheKey, str, timer).Err()
	return retMap, err
}

// UploadTicketPhoto
// @function: UploadTicketPhoto
// @description: 上传图片（需要信息数据上传成功后再上传图片）
// @param: u *model.SysUser, newPassword string
// @return: userInter *model.SysUser,err error
// /api/ticket/uploadPhoto?stationId=12
func UploadTicketPhoto(c *gin.Context, params requestDto.TicketUploadReq) (string, error) {
	ticket := model.Ticket{}
	//验证信息是否存在
	firstTicket := global.DB.
		Where("client_ticket_id=?", params.TicketId).
		First(&ticket)
	if firstTicket.RowsAffected <= 0 {
		global.Logger.Info("上传图片时 信息不存在 客户端信息id：" + params.TicketId)
		return "", errors.New("信息不存在！")
	}
	//验证站点是否存在
	stationId, _ := c.Get("stationId")
	// 转字符串
	stationIdSrt := fmt.Sprintf("%v", stationId)
	station := model.Station{}
	firstStation := global.DB.
		Where("id=?", stationIdSrt).
		First(&station)
	if firstStation.RowsAffected <= 0 {
		global.Logger.Error("站点不存在! 站点id:" + stationIdSrt)
		return "", errors.New("站点不存在！")
	}
	//用站点id和信息id从缓存中获取图片名字
	cacheKey := fmt.Sprintf("cache:ticketPhoto:stationId:%s:ticketId:%s", stationIdSrt, params.TicketId)
	count, err := global.Redis.Exists(context.Background(), cacheKey).Result()
	if err != nil {
		global.Logger.Info("上传的信息图片缓存不存在！"+cacheKey, zap.Error(err))
		return "", errors.New("上传的信息图片缓存不存在！")
	}
	if count <= 0 {
		global.Logger.Info("上传的信息图片缓存不存在！图片长度小于0", zap.Error(err))
		return "", errors.New("上传的信息图片缓存不存在！")
	}
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.Logger.Info("上传文件失败", zap.Error(err))
		return "", errors.New("上传文件失败！")
	}
	s := strings.Split(header.Filename, ".")
	//上传名字名
	fileName := s[0]
	//获取缓存中的
	pic, err := global.Redis.Get(context.Background(), cacheKey).Result()

	map2 := make(map[string]any)
	// json_decode 到map里
	_ = json.Unmarshal([]byte(pic), &map2)
	var columnName string
	for k, v := range map2 {
		if v == fileName {
			columnName = fmt.Sprintf("%v", k)
		}
	}

	if columnName == "" {
		global.Logger.Info("图片名称不存在 "+fmt.Sprintf("图片:%+v", map2), zap.Error(err))
		return "", errors.New("图片名称不存在！")
	}
	arr := []string{"photo_front", "photo_behind", "photo_lorry_number", "photo_goods"}
	//判断对应键值是否存在指定列名
	index := arrays.ContainsString(arr, columnName)
	fmt.Println(index)
	if index == -1 {
		global.Logger.Info("上传的文件名不正确 " + fmt.Sprintf("图片:%s", columnName))
		return "", errors.New("上传的文件名不正确！")
	}

	// 转为字符串
	pathStr := fmt.Sprintf("%v", ticket.Id)
	//实现图片上传
	filePath, cacheKey, err := upload.UploadFile(header, pathStr, columnName)
	if err != nil {
		global.Logger.Info("上传文件失败 "+fmt.Sprintf("图片路径:%s", filePath), zap.Error(err))
		return "", errors.New("上传文件失败！")
	}

	_, err = UpdateTicket(params.TicketId, columnName, filePath)
	if err == nil {
		global.Logger.Info("上传图片成功后，修改信息失败 "+
			fmt.Sprintf("客户端信息id:%s 图片名称：%s 图片路径%+v", params.TicketId, columnName, filePath),
			zap.Error(err))
		return columnName, err
	}
	return columnName, err
}
