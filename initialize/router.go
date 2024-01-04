package initialize

import "github.com/gin-gonic/gin"

func Routers() *gin.Engine {
	Router := gin.Default()

	PublicGroup := Router.Group("")

	// 测试路由
	{
		PublicGroup.GET("/test", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	// 公共路由
	{

	}

	// 私有路由
	// PrivateGroup := Router.Group("")
	{
	}

	return Router
}
