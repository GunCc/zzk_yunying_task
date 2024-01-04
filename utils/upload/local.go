package upload

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/utils"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
)

var (
	LOCAL_PATH = "uploads/file"
)

type Local struct {
}

func (l *Local) UploadFile(file *multipart.FileHeader) (string, string, error) {

	// 读取文件后缀名
	ext := path.Ext(file.Filename)

	// 读取文件名后加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建路径
	mkdirErr := os.MkdirAll("uploads/file", os.ModePerm)
	if mkdirErr != nil {
		global.TASK_LOGGER.Error("创建文件失败：", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("创建文件失败：" + mkdirErr.Error())
	}
	// 拼接路径
	path := LOCAL_PATH + "/" + filename
	storePath := LOCAL_PATH + "/" + filename

	// 读取文件
	f, openErr := file.Open()
	if openErr != nil {
		global.TASK_LOGGER.Error("文件打开失败：", zap.Any("err", openErr.Error()))
		return "", "", errors.New("文件打开失败：" + openErr.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(path)
	if createErr != nil {
		global.TASK_LOGGER.Error("文件创建失败：", zap.Any("err", createErr.Error()))
		return "", "", errors.New("文件创建失败：" + createErr.Error())
	}
	defer out.Close()

	// 传输拷贝文件
	_, copyErr := io.Copy(out, f)
	if createErr != nil {
		global.TASK_LOGGER.Error("文件拷贝失败：", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("文件拷贝失败：" + copyErr.Error())
	}
	return storePath, filename, nil
}
