package global

import (
	"ZZK_YUNYING_TASK/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	TASK_VIPER  *viper.Viper
	TASK_CONFIG config.Config
	TASK_ZAP    *zap.Logger
)
