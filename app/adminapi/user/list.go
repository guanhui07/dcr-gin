package user

import (
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	utils "dcr-gin/app/utils"
	"github.com/gin-gonic/gin"
)

// UserList 用户列表 get http://localhost:9090/adminapi/user/list?page=1&page_rows=10
func UserList(c *gin.Context) {
	var userResDto requestDto.UserListReq
	err := c.ShouldBindQuery(&userResDto)
	if err != nil {
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
	var userResDto requestDto.UserReq
	err := c.ShouldBindJSON(&userResDto)
	if err != nil {
		// 2.校验数据是否合法
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	err = service.AddUser(c, userResDto)
	if err == nil {
		utils.Success(c, "", "添加成功")
		return
	}
	utils.Fail(c, err.Error())
}

// ChangeStatus 修改状态
/**
{
"status":"sdf",
"id":"sdf",
}
*/
func ChangeStatus(c *gin.Context) {
	var userRes requestDto.ChangeUserStatusReq
	err := c.ShouldBindJSON(&userRes)
	if err != nil {
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

// ChangePwReq 更新用户密码
/**
{
"password":"sdf",
"id":12,
}
*/

func ChangePw(c *gin.Context) {
	var userResDto requestDto.ChangePwReq
	err := c.ShouldBindJSON(&userResDto)
	if err != nil {
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
	// 从协程上下文取 中间件set
	adminName, _ := c.Get("userName")
	if adminName != "admin" {
		utils.Fail(c, "非管理员账户,无法修改密码")
	}
	var userResDto requestDto.ChangeAdminPwReq
	err := c.ShouldBindJSON(&userResDto)
	if err != nil {
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
	var userResDto requestDto.EditUserReq
	err := c.ShouldBindJSON(&userResDto)
	if err != nil {
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
