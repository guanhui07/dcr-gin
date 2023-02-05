package user

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	utils "dcr-gin/app/utils"
	"github.com/gin-gonic/gin"
)

// AddUser 新增用户
/**
post http://localhost:9090/adminapi/user/add
{
"username":"fsd",
"password":"sdf",
}
*/
func AddUser(c *gin.Context) {
	// 1.获取前端传递过来的数据
	var userResDto requestDto.User
	if err := c.ShouldBindJSON(&userResDto); err != nil {
		// 2.校验数据是否合法
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	err := service.AddUser(c, userResDto)
	if err == nil {
		utils.Success(c, "", "添加成功")
		return
	}
	utils.Fail(c, err.Error())
}
