package test

import (
	"dcr-gin/app/global"
	"dcr-gin/app/requestDto"
	"dcr-gin/app/service"
	"dcr-gin/app/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/errorx"
	"go.uber.org/zap"
)

func Test1(c *gin.Context) {

	//c.Set("userName","admin")
	admin_name := c.Query("id")
	admin_name1 := c.DefaultQuery("page", "0")

	//$_GET
	admin_name4, exists := c.Get("userName")
	// $_POST
	admin_name3 := c.PostForm("test1")
	c.Request.FormValue("test2")
	// $_GET
	admin_name1 = c.Query("userName1")
	// 获取uri path /abc/:path
	admin_name1 = c.Param("path")
	admin_name2 := c.PostFormArray("userName2")
	dump.P(admin_name1, admin_name2, admin_name3, admin_name4)
	if !exists {
		utils.Fail(c, "userName未传")
		return
	}
	if admin_name != "admin" {
		utils.Fail(c, "非管理员账户,无法修改密码")
		return
	}
	var userRes requestDto.ChangeAdminPw
	if err := c.ShouldBindJSON(&userRes); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(c, message)
		return
	}

	result := service.ChangeAdminPw(c, userRes)

	if result == nil {
		utils.Success(c, "", "修改密码成功")
		return
	}
	utils.Fail(c, result.Error())

	apiBodyList := make([]map[string]interface{}, 0)
	itemMap1 := map[string]interface{}{
		"k": "v1",
	}
	apiBodyList = append(apiBodyList, itemMap1)

}

func Test2(c *gin.Context) error {
	url := "abc"
	reqMap := "a1"
	loggerStr := fmt.Sprintf("url:%+v,reqMap:%s", url, reqMap)
	// zap打印日志
	global.Logger.Info("post请求参数解析错误", zap.String("http", loggerStr))
	// 返回错误
	return errors.New("test")
}

func Test3(c *gin.Context) error {
	url := "abc"
	reqMap := "a1"
	loggerStr := fmt.Sprintf("url:%+v,reqMap:%s", url, reqMap)
	// zap打印日志
	global.Logger.Info("post请求参数解析错误", zap.String("http", loggerStr))
	// 返回错误
	return errorx.New("test")
}
