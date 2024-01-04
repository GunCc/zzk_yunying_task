package internal

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type gorm_config struct {
}

var GormConfig = new(gorm_config)

// 对gorm进行配置
func (g *gorm_config) GetGormConfig() *gorm.Config {

	// 设置创建表时,不受外键约束

	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}

	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	// 设置日志模式
	config.Logger = _default.LogMode(logger.Error)

	return config
}
