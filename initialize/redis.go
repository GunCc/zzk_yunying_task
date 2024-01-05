package initialize

import (
	"ZZK_YUNYING_TASK/global"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.TASK_CONFIG.Redis

	fmt.Println("redis配置", global.TASK_CONFIG.Redis)
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.TASK_LOGGER.Error("redis链接失败, 错误:", zap.Error(err))
	} else {
		global.TASK_LOGGER.Info("redis链接返回:", zap.String("pong", pong))
		global.TASK_REDIS = client
	}
}
