package router

import "ZZK_YUNYING_TASK/router/system"

// 路由组
type RouterGroup struct {
	System system.SystemRouterGroup
}

var RouterGroupApp = new(RouterGroup)
