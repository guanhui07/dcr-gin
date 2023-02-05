package system

import (
	"context"
	"dcr-gin/app/global"
	"dcr-gin/app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// Logout @Summary
// @Tags 退出登录接口
// @version 1.0
// @description 用户登录
// @Produce application/json
// Success 200 {object} Res {"code":0,"data":gin.H{},"message":"退出成功"}
// @Router /login [post]
/**
post http://localhost:9090/adminapi/system/logout

*/
func Logout(c *gin.Context) {
	userId, _ := c.Get("userId")
	key := fmt.Sprintf("cache:logout:userId:%v", userId)
	timer := time.Duration(720) * time.Second
	// setNx 锁
	err := global.Redis.SetNX(context.Background(), key, userId, timer).Err()
	if err == nil {
		global.Redis.Set(context.Background(), key, userId, time.Duration(1)*time.Second)
		global.Redis.Del(context.Background(), key)
		utils.Success(c, gin.H{}, "退出成功")
		return
	}
	utils.Fail(c, "退出登陆失败")

}
