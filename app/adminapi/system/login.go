package system

import (
	"context"
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"dcr-gin/app/requestDto"
	"dcr-gin/app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// Login @Summary 用户登录接口
// @Tags 用户登录
// @title 用户名和密码登录系统
// @version 1.0
// @description 用户登录
// @Produce application/json
// Param loginDto body dto.LoginDto true "用户登录参数"
// Success 200 {object} Res {"code":0,"data":null,"message":"操作成功"}
// @Router /login [post]
/**
post http://localhost:9090/adminapi/system/login
{
"username":"fsd",
"password":"sdf",
}
登录获取token
*/
func Login(c *gin.Context) {
	var loginDto requestDto.LoginDto
	// 解析前端传递过来的数据并且验证是否正确
	err := c.ShouldBindJSON(&loginDto)
	if err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}
	// 查询数据库登录操作
	user := model.User{}
	dbUserModel := global.DB.Model(user)
	// 查询一行
	dbUserModel.
		Where("username=?", loginDto.UserName).
		First(&user)
	//dump.P(user)
	if user.Id == 0 {
		utils.Fail(c, "账号不存在")
		return
	}
	if user.Status != 0 {
		utils.Fail(c, "账号已禁用")
		return
	}
	// 对账号和密码校验
	isOk, _ := utils.CheckPassword(user.Password, loginDto.Password, user.Salt)
	if !isOk {
		utils.Fail(c, "账号或密码错误")
		return
	}
	// token参数
	token := utils.TokenUser{
		Id:       user.Id,
		Username: user.UserName,
	}
	user.LastIp = c.ClientIP()
	user.UpdateTime = time.Now()
	// 更新一行
	dbUserModel.Save(&user)
	if token, err := utils.GenerateToken(token); err == nil {
		key := fmt.Sprintf("cache:logout:userId:%v", user.Id)
		global.Redis.Del(context.Background(), key)
		utils.Success(c, gin.H{
			"token":    token,
			"id":       user.Id,
			"username": user.UserName,
		}, "登陆成功")
		return
	}

	utils.Fail(c, "账号或密码错误")
	return

}
