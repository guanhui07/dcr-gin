package user

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	utils "dcr-gin/app/utils"
	"github.com/gin-gonic/gin"
)

// UserList 用户列表 get http://localhost:9090/adminapi/user/list?page=1&page_rows=10
func UserList(c *gin.Context) {
	var userResDto requestDto.UserList
	if err := c.ShouldBindQuery(&userResDto); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	responseErr, list := service.UserList(c, userResDto)

	if responseErr == nil {
		utils.Success(c, list, "获取配置成功")
		return
	}
	utils.Fail(c, responseErr.Error())
}
