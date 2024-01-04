package system

import "ZZK_YUNYING_TASK/service"

type SystemApiGroup struct {
	SysUserApi
}

var (
	userService = service.ServiceGroupApp.SysUserService
)
