package core

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/initialize"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

// 启动服务
func RunServer() {
	// 初始化redis服务，涉及到用户权限方面使用jwt，所以开启redis服务
	initialize.Redis()

	// 获取端口
	port := fmt.Sprintf(":%d", global.TASK_CONFIG.Server.Port)

	// 初始化路由
	Router := initialize.Routers()

	s := initServer(port, Router)
	time.Sleep(10 * time.Microsecond)
	global.TASK_LOGGER.Info("启动成功，端口号：", zap.String("端口号：", port))
	global.TASK_LOGGER.Error(s.ListenAndServe().Error())
}
