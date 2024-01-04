package system

import (
	"ZZK_YUNYING_TASK/api"

	"github.com/gin-gonic/gin"
)

// 用户相关路由
type UserRouter struct {
}

func (UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	userApi := api.ApiGroupApp.SystemApiGroup.SysUserApi
	{
		// 登陆注册路由
		userRouter.POST("/register", userApi.Register)
		userRouter.POST("/login", userApi.Login)
	}
}
