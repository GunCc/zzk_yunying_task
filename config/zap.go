package config

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

type Zap struct {
	Level    string `yaml:"level" json:"level" `      // 级别
	Director string `yaml:"director" json:"director"` // 存放的目录名
	MaxAge   int    `yaml:"maxAge" json:"maxAge" `    // 存活时间 以天为单位
	Format   string `yaml:"format" json:"format" `    // 打印格式 有json和console
}

// 根据字符串转换为 zapcore.Level
func (z *Zap) TransportLevel() zapcore.Level {
	// 转换为小写
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
