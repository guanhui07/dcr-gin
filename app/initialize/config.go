package initialize

import (
	"dcr-gin/app/global"
	"dcr-gin/app/utils"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	// 获取到main.go 目录
	workDir, _ := os.Getwd()
	//isDev := GetEnvInfo("IS_DEV")
	//fmt.Println(workDir, "目录")
	configFilePath := path.Join(workDir, "application.prod.yml")
	//fmt.Println(configFilePath, "文件")
	if utils.IsDev() {
		//dump.P(2222)
		configFilePath = path.Join(workDir, "application.dev.yml")
	}
	// 实例化 viper
	v := viper.New()
	//文件路径设置
	v.SetConfigFile(configFilePath)
	// 读yaml配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 解析yaml到 结构体
	err := v.Unmarshal(&global.ServerConfig)
	if err != nil {
		fmt.Println("读取配置失败")
	}
	//fmt.Println(&global.ServerConfig)
}
