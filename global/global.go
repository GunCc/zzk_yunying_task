package global

import (
	"ZZK_YUNYING_TASK/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	TASK_VIPER  *viper.Viper
	TASK_CONFIG config.Config
	TASK_LOGGER *zap.Logger
	TASK_REDIS  *redis.Client

	TASK_DB *gorm.DB
)
