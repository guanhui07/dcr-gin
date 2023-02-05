package middleware

import (
	"context"
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"dcr-gin/app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// AuthMiddleWare 中间件校验token登录   --后台登录
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		tokeString := c.GetHeader("Authorization")
		fmt.Println("当前Authorization", tokeString)
		if tokeString == "" {
			//utils.FailCode(c, "请先登录", 300)
			//c.Abort()
			//return
		}
		// 从token中解析出数据
		token, claims, err := utils.ParseToken(tokeString)
		if err != nil || !token.Valid {
			//utils.FailCode(c, "登陆状态失效", 300)
			//c.Abort()
			//return
		}
		key := fmt.Sprintf("cache:logout:userId:%v", claims.UserId)
		authorization, _ := global.Redis.StrLen(context.Background(), key).Result()
		if authorization > 0 {
			//utils.FailCode(c, "已退出登陆，请重新登陆", 300)
			//c.Abort()
			//return
		}
		// 可以进一步查询数据库,是否有当前的用户id
		if first := global.DB.Where("id=?", claims.UserId).First(&model.User{}); first.Error != nil {
			//utils.FailCode(c, "登陆状态无效", 300)
			//c.Abort()
			//return
		}
		// 从token中解析出来的数据挂载到上下文上,方便后面的控制器使用
		//c.Set("userId", claims.UserId)
		//c.Set("userName", claims.Username)
		c.Set("userId", 1)
		c.Set("userName", "admin")
		c.Next()
	}
}
