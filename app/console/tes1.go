package main

import (
	"dcr-gin/app/global"
	"dcr-gin/app/initialize"
	"dcr-gin/app/model"
	"errors"
	"github.com/gookit/goutil/dump"
	"gorm.io/gorm"
	// 初始化数据库连接及日志文件
	_ "dcr-gin/app/common"
)

// go run app/console/test1.go
func init() {
	initialize.InitConfig()
	// 初始化自定义校验器
	initialize.InitValidate()
	//注册redis
	initialize.Redis()
}

func main() {
	//dump.P("aaa")
	stationModel := model.Station{}
	// 查找一行
	db := global.DB.Model(stationModel)
	if errors.Is(db.Where("id=?", 1).First(&stationModel).Error, gorm.ErrRecordNotFound) {
		//记录不存在
		dump.P("此站点不存在")
		return
	}
	dump.P("此站点存在")
}
