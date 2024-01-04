package upload

import (
	"mime/multipart"
)

// OOS接口包含两个要实现的方法 一个是上传文件，一个是删除文件
type OOS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
}

// 根据配置文件实例化一个OOS接口
func NewOOS() OOS {
	return &Local{}
}
