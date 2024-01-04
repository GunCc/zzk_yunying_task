package api

import "ZZK_YUNYING_TASK/api/system"

type ApiGroup struct {
	SystemApiGroup system.SystemApiGroup
}

var ApiGroupApp = new(ApiGroup)
