package main

import (
	"ZZK_YUNYING_TASK/core"
	"ZZK_YUNYING_TASK/global"
)

func main() {
	// 配置Viper服务读取配置文件，因为我们后续会将服务部署上去，所以使用viper好管理一些配置属性
	global.TASK_VIPER = core.Viper()
	// 日志系统
	global.TASK_LOGGER = core.Zap()

	// 启动服务
	core.RunServer()
}
