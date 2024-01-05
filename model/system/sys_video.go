package system

import "ZZK_YUNYING_TASK/global"

type SysVideo struct {
	global.TASK_MODEL
	Url            string `json:"url" gorm:"comment:路径"`
	Name           string `json:"name" gorm:"comment:名称"`
	Tag            string `json:"tag" gorm:"comment:标签"`
	InputFileName  string `json:"input_file_name" gorm:"comment:当前文件名"`
	OutputFileName string `json:"output_file_name" gorm:"comment:输出文件名"`
	UserId         uint   `json:"user_id" gorm:"comment:上传用户的ID"`
	Status         int    `json:"status" gorm:"comment:视频转码状态 0 为进行中 1 为完成 2为失败；default:0"`
}
