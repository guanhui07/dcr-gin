package user

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	utils "dcr-gin/app/utils"
	"github.com/gin-gonic/gin"
)

// ChangeStatus 修改状态
/**
{
"status":"sdf",
"id":"sdf",
}
*/
func ChangeStatus(c *gin.Context) {
	var userRes requestDto.ChangeUserStatus
	if err := c.ShouldBindJSON(&userRes); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}

	result := service.ChangeStatus(c, userRes)

	if result == nil {
		utils.Success(c, "", "修改用户状态成功")
		return
	}
	utils.Fail(c, result.Error())
}

// ChangePw 更新用户密码
/**
{
"password":"sdf",
"id":12,
}
*/

func ChangePw(c *gin.Context) {
	var userResDto requestDto.ChangePw
	if err := c.ShouldBindJSON(&userResDto); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	result := service.ChangePw(c, userResDto)

	if result == nil {
		utils.Success(c, "", "修改密码成功")
		return
	}
	utils.Fail(c, result.Error())
}

// ChangAdminPw 修改admin密码
/*
post http://localhost:9090/adminapi/user/change/admin/passwd
{
"old_password":"fsd",
"repeat_password":"sdf",
"password":"sdf",
}
*/
func ChangAdminPw(c *gin.Context) {
	//c.Set("userName","admin")
	adminName, _ := c.Get("userName")
	if adminName != "admin" {
		utils.Fail(c, "非管理员账户,无法修改密码")
	}
	var userResDto requestDto.ChangeAdminPw
	if err := c.ShouldBindJSON(&userResDto); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}

	result := service.ChangeAdminPw(c, userResDto)

	if result == nil {
		utils.Success(c, "", "修改密码成功")
		return
	}
	utils.Fail(c, result.Error())

}
