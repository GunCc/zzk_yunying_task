package utils

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/model/system/request"

	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	// 从请求头中获取 x-token
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.TASK_LOGGER.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}
