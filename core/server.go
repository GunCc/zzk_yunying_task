package core

import "ZZK_YUNYING_TASK/initialize"

// 启动服务
func RunServer() {
	// 初始化redis服务，涉及到用户权限方面使用jwt，所以开启redis服务
	initialize.Redis()
}
