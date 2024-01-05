package system

import (
	"ZZK_YUNYING_TASK/api"

	"github.com/gin-gonic/gin"
)

type NotifyRouter struct {
}

func (NotifyRouter) InitNotifyRouter(router *gin.RouterGroup) {
	notifyRouter := router.Group("notify")
	notifyApi := api.ApiGroupApp.SystemApiGroup.SysNotifyApi
	{
		notifyRouter.GET("socket-connection", notifyApi.SocketConnection)
	}
}
