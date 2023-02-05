package middleware

import (
	"dcr-gin/app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RecoverMiddleWare() gin.HandlerFunc {
	//闭包
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.Fail(c, fmt.Sprint(err))
				c.Abort()
				return
			}
		}()
	}
}
