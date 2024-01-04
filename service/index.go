package service

import "ZZK_YUNYING_TASK/service/system"

type ServiceGroup struct {
	system.SystemServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
