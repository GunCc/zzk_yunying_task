package main

import (
	"ZZK_YUNYING_TASK/core"
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/initialize"

	"go.uber.org/zap"
)

func main() {
	// 配置Viper服务读取配置文件，因为我们后续会将服务部署上去，所以使用viper好管理一些配置属性
	global.TASK_VIPER = core.Viper()
	// 日志系统
	global.TASK_LOGGER = core.Zap()

	// 替换全局日志
	zap.ReplaceGlobals(global.TASK_LOGGER)

	// 连接数据库
	global.TASK_DB = initialize.Gorm()

	// 注意!!如果数据库打开了,在程序关闭前要记得close
	if global.TASK_DB != nil {
		// 如果数据库实例存在初始化表
		initialize.RegisterTables(global.TASK_DB)
		db, _ := global.TASK_DB.DB()
		defer db.Close()
	}

	// 启动服务
	core.RunServer()
}
