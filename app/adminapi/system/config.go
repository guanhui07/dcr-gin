package system

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	utils "dcr-gin/app/utils"
	"github.com/gin-gonic/gin"
)

// EditConfig 设置 post json
/**
post http://localhost:9090/adminapi/system/edit/config
{
"upload_auto_retry":2,
"upload_fail_retry_time":2,
"heartbeat_time":2
}
*/
func EditConfig(c *gin.Context) {
	var configDto requestDto.CreateConfig
	err := c.ShouldBindJSON(&configDto)
	if err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	err = service.AddConfig(c, configDto)
	if err == nil {
		utils.Success(c, "", "编辑配置成功")
		return
	}
	utils.Fail(c, err.Error())
}

// ConfigInfo a
/**
get http://localhost:9090/adminapi/system/config/info
curl -X GET http://localhost:9090/adminapi/system/config/info
{
"upload_auto_retry":2,
"upload_fail_retry_time":2,
"heartbeat_time":2
}
*/
func ConfigInfo(c *gin.Context) {
	var config requestDto.CreateConfig
	err, responseConfig := service.ConfigInfo(c, config)
	if err == nil {
		utils.Success(c, responseConfig, "获取配置成功")
		return
	}
	utils.Fail(c, err.Error())

}
