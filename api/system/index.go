package system

import "ZZK_YUNYING_TASK/service"

type SystemApiGroup struct {
	SysUserApi
	SysVideoApi
	SysNotifyApi
}

var (
	userService      = service.ServiceGroupApp.SysUserService
	SysVideoService  = service.ServiceGroupApp.SysVideoService
	SysNotifyService = service.ServiceGroupApp.SysNotifyService
)
