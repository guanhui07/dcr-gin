package initialize

import (
	"context"
	"dcr-gin/app/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.ServerConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("redis 连接失败, err:", zap.Error(err))
		return
	}
	global.Logger.Info("redis 连接成功:", zap.String("pong", pong))
	global.Redis = client
}
