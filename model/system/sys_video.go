package system

import "ZZK_YUNYING_TASK/global"

type SysVideo struct {
	global.TASK_MODEL
	Url       string `json:"url" gorm:"comment:路径"`
	Name      string `json:"name" gorm:"comment:名称"`
	Tag       string `json:"tag" gorm:"comment:标签"`
	Key       string `json:"key" gorm:"comment:编号"`
	UserId    uint   `json:"user_id" gorm:"comment:上传用户的ID"`
	StartTime string `json:"start_time" gorm:"-"`
	EndTime   string `json:"end_time" gorm:"-"`
}
