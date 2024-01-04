package system

import "ZZK_YUNYING_TASK/service"

type SystemApiGroup struct {
	SysUserApi
	SysVideoApi
}

var (
	userService     = service.ServiceGroupApp.SysUserService
	SysVideoService = service.ServiceGroupApp.SysVideoService
)
