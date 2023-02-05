package common

import (
	"dcr-gin/app/global"
	"dcr-gin/app/initialize"
	"dcr-gin/app/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"
)

// 初始化db
func initDB() {
	//2. 初始化配置文件
	initialize.InitConfig()
	// 从配置文件中获取参数
	host := global.ServerConfig.DataSource.Host
	port := strconv.Itoa(global.ServerConfig.DataSource.Port)
	database := global.ServerConfig.DataSource.Database
	username := global.ServerConfig.DataSource.Username
	password := global.ServerConfig.DataSource.Password
	charset := global.ServerConfig.DataSource.Charset
	loc := global.ServerConfig.DataSource.Loc
	//fmt.Println(host, port, database, username, password, "===aaa===>")
	// 字符串拼接
	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)
	//fmt.Println("数据库连接:", sqlStr)
	// 配置日志输出
	newLogger := logger.Default.LogMode(logger.Info)
	if !utils.IsDev() {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // 缓存日志时间
				LogLevel:                  logger.Silent, // 日志级别
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,         // Disable color
			},
		)
	}

	db, err := gorm.Open(mysql.Open(sqlStr), &gorm.Config{
		Logger: newLogger,
		//打印sql
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("打开数据库失败", err)
		panic("打开数据库失败" + err.Error())
	}
	global.DB = db
}

func init() {
	fmt.Println("开始连接数据库")
	initDB()
}
