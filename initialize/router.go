package initialize

import (
	"ZZK_YUNYING_TASK/router"

	"github.com/gin-gonic/gin"

	_ "ZZK_YUNYING_TASK/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	PublicGroup := Router.Group("")

	// 系统路由
	systemRouter := router.RouterGroupApp.System

	// 生成swagger api文档
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
