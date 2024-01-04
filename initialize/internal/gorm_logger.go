package internal

import (
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// 将数据库的log 转换到zap 打印出来
func (w *writer) Printf(mes string, data ...interface{}) {
	w.Writer.Printf(mes, data...)
}
