package main

import (
	"dcr-gin/app/global"
	"dcr-gin/app/initialize"
	"dcr-gin/app/job"
	"dcr-gin/docs"
	"fmt"
	"github.com/gookit/goutil/dump"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"

	// 初始化数据库连接及日志文件
	_ "dcr-gin/app/common"
	// 数据模型中init方法的执行
	_ "dcr-gin/app/model"
	// 文档
	_ "dcr-gin/docs"
)

// @title 权限系统API文档
// @version 1.0
// @description 使用gin+mysql实现权限系统的api接口
// @host 127.0.0.1:9090/api/v1
// @BasePath
func main() {
	// 1.初始化配置
	initialize.InitConfig()
	// 初始化自定义校验器
	initialize.InitValidate()

	//2.初始化路由
	router := initialize.Routers()
	//注册redis
	initialize.Redis()
	// 获取端口号
	PORT := strconv.Itoa(global.ServerConfig.Port)
	fmt.Println(PORT + "当前端口")
	docs.SwaggerInfo.BasePath = "/api/v1"

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", PORT))
	// swagger访问地址:localhost:5555/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	dump.P("http://localhost:5555/swagger/index.html")
	global.Logger.Sugar().Infof("服务已经启动:localhost:%s", PORT)

	//定时轮询上传信息数据
	//TicketUpload()
	ch := make(chan int, 1)
	go job.Run(ch)

	// 启动服务
	err := router.Run(fmt.Sprintf(":%s", PORT))
	if err != nil {
		global.Logger.Sugar().Panic("服务启动失败:%s", err.Error())
	}
}
