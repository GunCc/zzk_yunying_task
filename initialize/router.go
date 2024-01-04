package initialize

import (
	"ZZK_YUNYING_TASK/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	PublicGroup := Router.Group("")

	// 系统路由
	systemRouter := router.RouterGroupApp.System

	// 测试路由
	{
		PublicGroup.GET("/test", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	// 公共路由
	{
		systemRouter.InitUserRouter(PublicGroup)
	}

	// 私有路由
	// PrivateGroup := Router.Group("")
	{
	}

	return Router
}
