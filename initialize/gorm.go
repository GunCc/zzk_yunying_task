package initialize

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/initialize/internal"
	"ZZK_YUNYING_TASK/model/system"
	"fmt"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	ms := global.TASK_CONFIG.Mysql

	fmt.Println("mysql配置", global.TASK_CONFIG.Mysql)

	if ms.DbName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       ms.Dsn(),
		DefaultStringSize:         191,   // string 类型默认的长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", ms.DbName)

	dsn := ms.MysqlEmptyDsn()

	err := InitDatabase(dsn, createSql)

	if err != nil {
		global.TASK_LOGGER.Error("数据库创建失败！", zap.Error(err))
		return nil
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.GormConfig.GetGormConfig()); err != nil {
		global.TASK_LOGGER.Error("数据库初始化失败！", zap.Error(err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}

// 初始化表
func RegisterTables(db *gorm.DB) {
	// 注册表
	err := db.AutoMigrate(
		&system.SysUser{},
		&system.SysVideo{},
	)
	if err != nil {
		global.TASK_LOGGER.Error("表初始化失败", zap.Error(err))
		os.Exit(0)
	}
	global.TASK_LOGGER.Info("表初始化成功")
}
