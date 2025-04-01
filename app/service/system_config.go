package service

import (
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"dcr-gin/app/requestDto"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func AddConfig(c *gin.Context, configDto requestDto.CreateConfigReq) error {
	configModel := model.Config{}
	var info *gorm.DB
	var userId int64
	var retryTime int32
	var heartBeatTime int32
	userId = 0

	if configDto.UploadAutoRetry != 0 {
		userId = c.GetInt64("userId") // 从Gin 的上下文取
	}

	//已经找到
	db := global.DB.Model(configModel)
	err := db.First(&configModel).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		retryTime = configModel.UploadFailRetryTime

		if configDto.UploadFailRetryTime != 0 {
			retryTime = configDto.UploadFailRetryTime
		}
		heartBeatTime = configModel.HeartbeatTime

		if configDto.HeartbeatTime != 0 {
			heartBeatTime = configDto.HeartbeatTime
		}

		configModel.UploadAutoRetry = userId
		configModel.UploadFailRetryTime = retryTime
		configModel.HeartbeatTime = heartBeatTime
		configModel.UpdateTime = time.Now().Format("2006-01-02 15:04:05")

		//更新
		info = db.
			Select("upload_auto_retry",
				"upload_fail_retry_time",
				"heartbeat_time", "update_time").
			Save(&configModel)
		return info.Error
	}
	// 找不到则 新增
	addConfig := model.Config{
		UploadAutoRetry:     userId,
		UploadFailRetryTime: configDto.UploadFailRetryTime,
		HeartbeatTime:       configDto.HeartbeatTime,
		AddTime:             time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime:          time.Now().Format("2006-01-02 15:04:05"),
	}
	info = db.Create(&addConfig)
	return info.Error
}

// ConfigInfo 获取详情
func ConfigInfo(c *gin.Context, config requestDto.CreateConfigReq) (err error, configInfo model.Config) {
	configModel := model.Config{}

	db := global.DB.Model(configModel)
	err = db.First(&configModel).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, configModel
	}
	return errors.New("暂无配置信息"), configModel
}
