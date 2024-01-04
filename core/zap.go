package core

import (
	"ZZK_YUNYING_TASK/core/internal"
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/utils"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {

	if b, _ := utils.PathExists(global.TASK_CONFIG.Zap.Director); !b {
		fmt.Println("创建文件:" + global.TASK_CONFIG.Zap.Director)
		os.Mkdir(global.TASK_CONFIG.Zap.Director, os.ModePerm)
	}
	cores := internal.Zap_Config.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))
	return logger
}
