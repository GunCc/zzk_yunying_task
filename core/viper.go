package core

import (
	"ZZK_YUNYING_TASK/core/internal"
	"ZZK_YUNYING_TASK/global"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 优先级：命令行 > 默认值
func Viper(path ...string) *viper.Viper {
	var config string

	// 获取环境变量
	if len(path) == 0 {
		switch gin.Mode() {
		case gin.DebugMode:
			config = internal.ConfigDefaultFile
		case gin.TestMode:
			config = internal.ConfigTestFile
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
	}

	fmt.Printf("您当前使用的配置文件是%s\n", config)

	viper := viper.New()

	// 设置文件
	viper.SetConfigFile(config)

	// 设置文件类型
	viper.SetConfigType("yaml")

	// 读取文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("文件错误：%s\n", err))
	}

	// 查看文件
	viper.WatchConfig()

	// 监听文件修改
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("文件修改为：%s \n", in.Name)
		// 序列化配置文件
		if err = viper.Unmarshal(&global.TASK_CONFIG); err != nil {
			fmt.Println("错误：", err)
		}
	})

	// 将配置文件数据存储到 TASK_CONFIG 中
	if err = viper.Unmarshal(&global.TASK_CONFIG); err != nil {
		fmt.Println("错误：", err)

	}
	return viper
}
