package middleware

import (
	"dcr-gin/app/global"
	"dcr-gin/app/model"
	"dcr-gin/app/requestDto"
	"dcr-gin/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// ApiMiddleWare 中间件校验token登录  api -站点客户端 登录 验证
func ApiMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取sign
		var stationSign requestDto.StationSign
		if c.ContentType() == "multipart/form-data" {
			stationSign.Sign = c.Request.FormValue("sign")
			stationSign.Device = c.Request.FormValue("device")
		} else {
			if err := c.ShouldBindBodyWith(&stationSign, binding.JSON); err != nil {
				// 校验数据是否合法
				//msg := utils.ShowErrorMessage(err)
				//utils.FailMiddleWare(c, msg)
				//c.Abort()
				//
				//return
			}
		}
		if stationSign.Sign == "" {
			//utils.FailMiddleWare(c, "签名信息不存在")
			//c.Abort()
			//return
		}
		stationId, _ := utils.HashidsDecode(stationSign.Device)
		if stationId == "" {
			//utils.FailMiddleWare(c, "站点不存在！")
			//c.Abort()
			//return
		}
		// 可以进一步查询数据库,是否有站点id
		station := model.Station{}
		if first := global.DB.Where("id=?", stationId).
			Where("status=0").First(&station); first.Error != nil {
			//utils.FailMiddleWare(c, "站点不存在或已停用！")
			//c.Abort()
			//return
		}
		if station.TigerShaped == "" {
			//utils.FailMiddleWare(c, "站点的握手符号不正确，请先登陆我方平台设置握手符号！")
			//c.Abort()
			//return
		}
		// 数据挂载到上下文上,方便后面的控制器使用
		c.Set("device", stationSign.Device)
		c.Set("sign", stationSign.Sign)
		c.Set("stationId", stationId)
		c.Set("tigerShaped", station.TigerShaped)
		c.Next()
	}
}
