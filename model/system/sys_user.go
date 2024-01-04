package system

import (
	"ZZK_YUNYING_TASK/global"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.TASK_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`
	NickName string    `json:"nickname" gorm:"index;comment:昵称"`
	Password string    `json:"password" gorm:"password;comment:密码"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
