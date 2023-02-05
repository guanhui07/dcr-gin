package utils

import (
	"github.com/chenhg5/collection"
	"github.com/spf13/viper"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func IsDev() bool {
	isDev := GetEnvInfo("IS_DEV")
	return true
	return isDev
}

func Collect(in interface{}) collection.Collection {
	//判断in 类型是否是 slice map [] struct  里面已经判断了
	return collection.Collect(in)
}
