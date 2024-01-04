package main

import (
	"ZZK_YUNYING_TASK/core"
	"ZZK_YUNYING_TASK/global"
	"fmt"
)

func main() {
	// 配置Viper服务读取配置文件，因为我们后续会将服务部署上去，所以使用viper好管理一些配置属性
	global.TASK_VIPER = core.Viper()

	fmt.Println("运行成功")
}
