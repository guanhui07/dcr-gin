package user

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	utils "dcr-gin/app/utils"
	"github.com/gin-gonic/gin"
)

// UpdateUser 更新用户详情
// {"id":"23","username":"aa","status":1}
/**
post http://localhost:9090/adminapi/user/update
{
"username":"fsd",
"status":"sdf",
"id":"sdf",
}
*/
func UpdateUser(c *gin.Context) {
	var userResDto requestDto.EditUser
	if err := c.ShouldBindJSON(&userResDto); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	result := service.EditUser(c, userResDto)
	if result == nil {
		utils.Success(c, "", "修改用户信息成功")
		return
	}
	utils.Fail(c, result.Error())
}
