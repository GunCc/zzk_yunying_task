package system

import (
	"ZZK_YUNYING_TASK/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SysNotifyApi struct {
}

// 建立SSE链接
func (v *SysNotifyApi) SocketConnection(ctx *gin.Context) {
	claims, _ := utils.GetClaims(ctx)
	id := claims.ID
	SysNotifyService.BuildNotificationChannel(strconv.Itoa(int(id)), "123", ctx)
}
